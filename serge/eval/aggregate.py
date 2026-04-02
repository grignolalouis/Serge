#!/usr/bin/env python3
"""Aggregate SERGE evaluation results into a structured report."""

import json
import math
import os
import sys


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
    total_cost = 0.0
    total_duration = 0
    total_tool_calls = 0

    for scenario_file in sorted(os.listdir(scenarios_dir)):
        if not scenario_file.endswith(".json"):
            continue
        scenario = load_json(os.path.join(scenarios_dir, scenario_file))
        if not scenario:
            continue

        sid = scenario["id"]
        result_dir = os.path.join(results_dir, sid)
        if not os.path.isdir(result_dir):
            continue

        runs = []
        scores = []

        for r in range(1, runs_per_scenario + 1):
            agent = load_json(os.path.join(result_dir, f"run_{r}.json"))
            judge = load_json(os.path.join(result_dir, f"run_{r}_judge.json"))
            if not agent or not judge:
                continue

            ws = judge.get("weighted_score", 0)
            passed = judge.get("pass", False)
            dur = agent.get("duration_ms", 0)
            cost = agent.get("cost_usd", 0)
            turns = agent.get("num_turns", 0)

            scores.append(ws)
            all_scores.append(ws)
            all_passes.append(1 if passed else 0)
            total_cost += cost
            total_duration += dur
            total_tool_calls += turns

            runs.append({
                "run": r,
                "weighted_score": round(ws, 2),
                "pass": passed,
                "duration_ms": dur,
                "tool_calls_count": turns,
                "cost_usd": round(cost, 4),
            })

        if not runs:
            continue

        avg = round(sum(scores) / len(scores), 2)
        pr = round(sum(1 for s in scores if s >= 3.5) / len(scores), 2)
        sd = round(std_dev(scores), 2)

        by_scenario.append({
            "scenario_id": sid,
            "title": scenario["title"],
            "use_case": scenario["use_case"],
            "difficulty": scenario["difficulty"],
            "avg_score": avg,
            "pass_rate": pr,
            "consistency_std": sd,
            "runs": runs,
        })

    # By use case
    uc_groups = {}
    for s in by_scenario:
        uc = s["use_case"]
        uc_groups.setdefault(uc, []).append(s)

    by_use_case = []
    for uc, items in sorted(uc_groups.items()):
        uc_scores = [r["weighted_score"] for s in items for r in s["runs"]]
        uc_passes = [r["pass"] for s in items for r in s["runs"]]
        uc_tools = [r["tool_calls_count"] for s in items for r in s["runs"]]
        uc_dur = [r["duration_ms"] for s in items for r in s["runs"]]
        by_use_case.append({
            "use_case": uc,
            "avg_score": round(sum(uc_scores) / len(uc_scores), 2) if uc_scores else 0,
            "pass_rate": round(sum(uc_passes) / len(uc_passes), 2) if uc_passes else 0,
            "avg_tool_calls": round(sum(uc_tools) / len(uc_tools), 1) if uc_tools else 0,
            "avg_duration_ms": round(sum(uc_dur) / len(uc_dur)) if uc_dur else 0,
        })

    # By difficulty
    diff_groups = {}
    for s in by_scenario:
        d = s["difficulty"]
        diff_groups.setdefault(d, []).append(s)

    by_difficulty = []
    for d in ["easy", "medium", "hard"]:
        items = diff_groups.get(d, [])
        if not items:
            continue
        d_scores = [r["weighted_score"] for s in items for r in s["runs"]]
        d_passes = [r["pass"] for s in items for r in s["runs"]]
        by_difficulty.append({
            "difficulty": d,
            "avg_score": round(sum(d_scores) / len(d_scores), 2) if d_scores else 0,
            "pass_rate": round(sum(d_passes) / len(d_passes), 2) if d_passes else 0,
        })

    # Summary
    n = len(all_scores)
    summary = {
        "total_runs": n,
        "pass_rate": round(sum(all_passes) / n, 2) if n else 0,
        "avg_weighted_score": round(sum(all_scores) / n, 2) if n else 0,
        "avg_factual_accuracy": round(sum(s.get("factual_accuracy", 0) for s in []) / max(n, 1), 2),
        "avg_completeness": round(sum(s.get("completeness", 0) for s in []) / max(n, 1), 2),
        "avg_reasoning": round(sum(s.get("reasoning", 0) for s in []) / max(n, 1), 2),
        "avg_source_attribution": round(sum(s.get("source_attribution", 0) for s in []) / max(n, 1), 2),
        "avg_structure": round(sum(s.get("structure", 0) for s in []) / max(n, 1), 2),
    }

    # Compute per-criterion averages from judge files
    criterion_sums = {"factual_accuracy": 0, "completeness": 0, "reasoning": 0, "source_attribution": 0, "structure": 0}
    criterion_count = 0
    for s in by_scenario:
        sid = s["scenario_id"]
        for r in range(1, runs_per_scenario + 1):
            judge = load_json(os.path.join(results_dir, sid, f"run_{r}_judge.json"))
            if judge:
                criterion_count += 1
                for k in criterion_sums:
                    criterion_sums[k] += judge.get(k, 0)

    if criterion_count > 0:
        summary["avg_factual_accuracy"] = round(criterion_sums["factual_accuracy"] / criterion_count, 2)
        summary["avg_completeness"] = round(criterion_sums["completeness"] / criterion_count, 2)
        summary["avg_reasoning"] = round(criterion_sums["reasoning"] / criterion_count, 2)
        summary["avg_source_attribution"] = round(criterion_sums["source_attribution"] / criterion_count, 2)
        summary["avg_structure"] = round(criterion_sums["structure"] / criterion_count, 2)

    report = {
        "summary": summary,
        "by_scenario": by_scenario,
        "by_use_case": by_use_case,
        "by_difficulty": by_difficulty,
        "metrics": {
            "total_cost_usd": round(total_cost, 4),
            "total_duration_ms": total_duration,
            "total_tool_calls": total_tool_calls,
        },
    }

    print(json.dumps(report, indent=2))


if __name__ == "__main__":
    main()
