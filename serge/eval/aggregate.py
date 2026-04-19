#!/usr/bin/env python3
"""Aggregate SERGE evaluation results into a structured report.

Captures score breakdowns by level and use case, plus supply-chain-agent-specific
metrics: time savings vs human baseline, system coverage, tool efficiency, and
failure taxonomy.
"""

import json
import math
import os
import sys
from collections import Counter


def load_json(path):
    try:
        with open(path) as f:
            return json.load(f)
    except (FileNotFoundError, json.JSONDecodeError):
        return None


def std_dev(values):
    if len(values) < 2:
        return 0.0
    mean = sum(values) / len(values)
    return math.sqrt(sum((x - mean) ** 2 for x in values) / len(values))


def main():
    scenarios_dir = sys.argv[1]
    results_dir = sys.argv[2]
    runs_per_scenario = int(sys.argv[3])

    by_scenario = []
    all_scores = []
    all_passes = []
    hallucination_runs = 0
    critical_missing_runs = 0
    same_conclusion_runs = 0
    total_cost = 0.0
    total_duration_ms = 0
    total_tool_calls = 0
    judge_totals = {"factual_accuracy": 0, "completeness": 0, "no_hallucination": 0, "reasoning": 0}
    judge_count = 0
    failure_counter = Counter()
    total_human_minutes_saved = 0.0
    total_human_minutes = 0
    total_agent_seconds = 0.0

    for scenario_file in sorted(os.listdir(scenarios_dir)):
        if not scenario_file.endswith(".json"):
            continue
        scenario = load_json(os.path.join(scenarios_dir, scenario_file))
        if not scenario:
            continue

        sid = scenario["id"]
        expected_systems = set(scenario.get("expected_systems", []))
        expected_min_tools = scenario.get("expected_tool_count_min", 1)
        human_min = scenario.get("estimated_human_minutes", 0)

        result_dir = os.path.join(results_dir, sid)
        if not os.path.isdir(result_dir):
            continue

        runs = []
        scores = []
        same_conclusions = []

        for r in range(1, runs_per_scenario + 1):
            agent = load_json(os.path.join(result_dir, f"run_{r}.json"))
            judge = load_json(os.path.join(result_dir, f"run_{r}_judge.json"))
            if not agent or not judge:
                continue

            ws = judge.get("weighted_score", 0)
            passed = judge.get("pass", False)
            critical = judge.get("critical_facts_present", False)
            same_conclusion = judge.get("same_conclusion_as_ground_truth", False)
            hallucinations = judge.get("hallucinations", []) or []
            failure_cat = judge.get("failure_category", "none")

            dur_ms = agent.get("duration_ms", 0)
            cost = agent.get("cost_usd", 0)
            tools_used = agent.get("tools_used", [])
            tool_calls = agent.get("tool_calls_count", agent.get("num_turns", 0))
            systems_touched = set(agent.get("systems_touched", []))

            # Derived: time savings
            agent_sec = dur_ms / 1000.0
            human_sec = human_min * 60.0
            time_saved_sec = max(0, human_sec - agent_sec)
            time_saved_pct = (time_saved_sec / human_sec * 100) if human_sec > 0 else 0

            # Derived: tool efficiency
            tool_efficiency = (expected_min_tools / tool_calls) if tool_calls > 0 else 0

            # Derived: system coverage
            systems_hit_expected = len(expected_systems & systems_touched)
            systems_coverage = (systems_hit_expected / len(expected_systems)) if expected_systems else 1.0

            scores.append(ws)
            same_conclusions.append(same_conclusion)
            all_scores.append(ws)
            all_passes.append(1 if passed else 0)
            if hallucinations:
                hallucination_runs += 1
            if not critical:
                critical_missing_runs += 1
            if same_conclusion:
                same_conclusion_runs += 1
            failure_counter[failure_cat] += 1
            total_cost += cost
            total_duration_ms += dur_ms
            total_tool_calls += tool_calls
            total_human_minutes_saved += time_saved_sec / 60.0
            total_human_minutes += human_min
            total_agent_seconds += agent_sec

            judge_count += 1
            for k in judge_totals:
                judge_totals[k] += judge.get(k, 0)

            runs.append({
                "run": r,
                "weighted_score": round(ws, 2),
                "pass": passed,
                "critical_facts_present": critical,
                "same_conclusion_as_ground_truth": same_conclusion,
                "failure_category": failure_cat,
                "hallucination_count": len(hallucinations),
                "duration_ms": dur_ms,
                "tool_calls_count": tool_calls,
                "tools_used": tools_used,
                "systems_touched": sorted(systems_touched),
                "system_coverage": round(systems_coverage, 2),
                "tool_efficiency": round(tool_efficiency, 2),
                "time_saved_pct": round(time_saved_pct, 1),
                "cost_usd": round(cost, 4),
            })

        if not runs:
            continue

        avg = round(sum(scores) / len(scores), 2)
        pr = round(sum(1 for r in runs if r["pass"]) / len(runs), 2)
        sd = round(std_dev(scores), 2)
        functional_consistency = round(sum(same_conclusions) / len(same_conclusions), 2) if same_conclusions else 0

        by_scenario.append({
            "scenario_id": sid,
            "title": scenario["title"],
            "level": scenario.get("level", "?"),
            "use_case": scenario.get("use_case", "?"),
            "expected_systems": sorted(expected_systems),
            "expected_tool_count_min": expected_min_tools,
            "estimated_human_minutes": human_min,
            "avg_score": avg,
            "pass_rate": pr,
            "consistency_std": sd,
            "functional_consistency": functional_consistency,
            "runs": runs,
        })

    # By complexity level
    level_groups = {}
    for s in by_scenario:
        level_groups.setdefault(s["level"], []).append(s)

    by_level = []
    for level in ["L1", "L2", "L3", "L4", "L5"]:
        items = level_groups.get(level, [])
        if not items:
            continue
        scores = [r["weighted_score"] for s in items for r in s["runs"]]
        passes = [r["pass"] for s in items for r in s["runs"]]
        tools = [r["tool_calls_count"] for s in items for r in s["runs"]]
        dur = [r["duration_ms"] for s in items for r in s["runs"]]
        coverage = [r["system_coverage"] for s in items for r in s["runs"]]
        efficiency = [r["tool_efficiency"] for s in items for r in s["runs"]]
        systems_avg = [len(r["systems_touched"]) for s in items for r in s["runs"]]
        time_saved = [r["time_saved_pct"] for s in items for r in s["runs"]]
        by_level.append({
            "level": level,
            "scenario_count": len(items),
            "avg_score": round(sum(scores) / len(scores), 2) if scores else 0,
            "pass_rate": round(sum(passes) / len(passes), 2) if passes else 0,
            "avg_tool_calls": round(sum(tools) / len(tools), 1) if tools else 0,
            "avg_duration_ms": round(sum(dur) / len(dur)) if dur else 0,
            "avg_systems_touched": round(sum(systems_avg) / len(systems_avg), 1) if systems_avg else 0,
            "avg_system_coverage": round(sum(coverage) / len(coverage), 2) if coverage else 0,
            "avg_tool_efficiency": round(sum(efficiency) / len(efficiency), 2) if efficiency else 0,
            "avg_time_saved_pct": round(sum(time_saved) / len(time_saved), 1) if time_saved else 0,
        })

    # By use case
    uc_groups = {}
    for s in by_scenario:
        uc_groups.setdefault(s["use_case"], []).append(s)

    by_use_case = []
    for uc, items in sorted(uc_groups.items()):
        scores = [r["weighted_score"] for s in items for r in s["runs"]]
        passes = [r["pass"] for s in items for r in s["runs"]]
        by_use_case.append({
            "use_case": uc,
            "scenario_count": len(items),
            "avg_score": round(sum(scores) / len(scores), 2) if scores else 0,
            "pass_rate": round(sum(passes) / len(passes), 2) if passes else 0,
        })

    # Summary
    n = len(all_scores)
    summary = {
        "total_scenarios": len(by_scenario),
        "total_runs": n,
        "pass_rate": round(sum(all_passes) / n, 2) if n else 0,
        "avg_weighted_score": round(sum(all_scores) / n, 2) if n else 0,
        "hallucination_rate": round(hallucination_runs / n, 2) if n else 0,
        "critical_facts_missing_rate": round(critical_missing_runs / n, 2) if n else 0,
        "functional_consistency_rate": round(same_conclusion_runs / n, 2) if n else 0,
        "total_human_baseline_minutes": round(total_human_minutes * runs_per_scenario, 0),
        "total_agent_minutes": round(total_agent_seconds / 60, 1),
        "total_time_saved_minutes": round(total_human_minutes_saved, 1),
        "avg_time_saved_pct_per_scenario": round(
            (total_human_minutes_saved / (total_human_minutes * runs_per_scenario) * 100)
            if total_human_minutes > 0 else 0, 1),
    }

    if judge_count > 0:
        summary["avg_factual_accuracy"] = round(judge_totals["factual_accuracy"] / judge_count, 2)
        summary["avg_completeness"] = round(judge_totals["completeness"] / judge_count, 2)
        summary["avg_no_hallucination"] = round(judge_totals["no_hallucination"] / judge_count, 2)
        summary["avg_reasoning"] = round(judge_totals["reasoning"] / judge_count, 2)

    failure_taxonomy = [
        {"category": cat, "count": cnt, "rate": round(cnt / n, 2) if n else 0}
        for cat, cnt in failure_counter.most_common()
    ]

    report = {
        "summary": summary,
        "failure_taxonomy": failure_taxonomy,
        "by_scenario": by_scenario,
        "by_level": by_level,
        "by_use_case": by_use_case,
        "metrics": {
            "total_cost_usd": round(total_cost, 4),
            "total_duration_ms": total_duration_ms,
            "total_tool_calls": total_tool_calls,
        },
    }

    print(json.dumps(report, indent=2))


if __name__ == "__main__":
    main()
