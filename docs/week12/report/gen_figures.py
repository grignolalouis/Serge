#!/usr/bin/env python3
"""Generate data-driven figures for the final SERGE report.

Reads the two evaluation reports (baseline / run 2) and produces PDF figures
using matplotlib. All output goes into ./figures/.
"""

import json
import os
from pathlib import Path

import matplotlib.pyplot as plt
import matplotlib as mpl
import numpy as np

# ──────────────────────────────────────────────────────────────────────────
# Styling
# ──────────────────────────────────────────────────────────────────────────
mpl.rcParams.update({
    "font.family": "serif",
    "font.size": 10,
    "axes.titlesize": 11,
    "axes.labelsize": 10,
    "legend.fontsize": 9,
    "xtick.labelsize": 9,
    "ytick.labelsize": 9,
    "axes.grid": True,
    "grid.alpha": 0.25,
    "grid.linestyle": "-",
    "axes.axisbelow": True,
    "figure.dpi": 100,
    "pdf.fonttype": 42,
})

# Palette — restrained, professional
C_BASELINE = "#B8B8B8"       # grey — run 1
C_IMPROVED = "#2E5E8B"       # blue — run 2
C_ACCENT   = "#D17A3E"       # orange — emphasis
C_OK       = "#4A7F4A"       # green — pass / success
C_WARN     = "#C44E4E"       # red — failure

SCRIPT_DIR = Path(__file__).parent
FIG_DIR = SCRIPT_DIR / "figures"
FIG_DIR.mkdir(exist_ok=True)

# ──────────────────────────────────────────────────────────────────────────
# Load data
# ──────────────────────────────────────────────────────────────────────────
PROJECT_ROOT = SCRIPT_DIR.parent.parent.parent
EVAL_DIR = PROJECT_ROOT / "serge" / "eval"

with open(EVAL_DIR / "report_baseline_run1" / "evaluation_report.json") as f:
    R1 = json.load(f)
with open(EVAL_DIR / "report_run2" / "evaluation_report.json") as f:
    R2 = json.load(f)


def save(fig, name):
    path = FIG_DIR / f"{name}.pdf"
    fig.savefig(path, bbox_inches="tight", pad_inches=0.05)
    plt.close(fig)
    print(f"  wrote {path.relative_to(PROJECT_ROOT)}")


# ──────────────────────────────────────────────────────────────────────────
# Figure: score by complexity level (run 1 vs run 2)
# ──────────────────────────────────────────────────────────────────────────
def fig_score_by_level():
    levels = ["L1", "L2", "L3", "L4", "L5"]
    r1 = {l["level"]: l["avg_score"] for l in R1["by_level"]}
    r2 = {l["level"]: l["avg_score"] for l in R2["by_level"]}
    s1 = [r1.get(l, 0) for l in levels]
    s2 = [r2.get(l, 0) for l in levels]

    fig, ax = plt.subplots(figsize=(6.5, 3.5))
    x = np.arange(len(levels))
    w = 0.38
    b1 = ax.bar(x - w/2, s1, w, label="Run 1 (baseline)", color=C_BASELINE, edgecolor="black", linewidth=0.5)
    b2 = ax.bar(x + w/2, s2, w, label="Run 2 (refined)", color=C_IMPROVED, edgecolor="black", linewidth=0.5)

    for bars in (b1, b2):
        for b in bars:
            h = b.get_height()
            ax.text(b.get_x() + b.get_width()/2, h + 0.06, f"{h:.2f}",
                    ha="center", va="bottom", fontsize=8)

    ax.set_xticks(x)
    ax.set_xticklabels([f"{l}\n{n}" for l, n in zip(levels,
        ["trivial", "simple", "moderate", "complex", "expert"])])
    ax.set_ylabel("Weighted score (0–5)")
    ax.set_ylim(0, 5.6)
    ax.set_title("Agent weighted score by scenario complexity level")
    ax.legend(loc="lower right", framealpha=0.95)
    ax.set_axisbelow(True)
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Figure: pass rate by use case (run 1 vs run 2)
# ──────────────────────────────────────────────────────────────────────────
def fig_pass_rate_by_usecase():
    ucs = ["UC-1", "UC-2", "UC-3", "UC-4", "UC-5"]
    labels = {
        "UC-1": "Order\nStatus",
        "UC-2": "Root\nCause",
        "UC-3": "Supplier\nSourcing",
        "UC-4": "Operational\nMonitoring",
        "UC-5": "Planning",
    }
    r1 = {u["use_case"]: u["pass_rate"] * 100 for u in R1["by_use_case"]}
    r2 = {u["use_case"]: u["pass_rate"] * 100 for u in R2["by_use_case"]}
    s1 = [r1.get(u, 0) for u in ucs]
    s2 = [r2.get(u, 0) for u in ucs]

    fig, ax = plt.subplots(figsize=(6.5, 3.5))
    x = np.arange(len(ucs))
    w = 0.38
    b1 = ax.bar(x - w/2, s1, w, label="Run 1 (baseline)", color=C_BASELINE, edgecolor="black", linewidth=0.5)
    b2 = ax.bar(x + w/2, s2, w, label="Run 2 (refined)", color=C_IMPROVED, edgecolor="black", linewidth=0.5)

    for bars in (b1, b2):
        for b in bars:
            h = b.get_height()
            ax.text(b.get_x() + b.get_width()/2, h + 2, f"{h:.0f}%",
                    ha="center", va="bottom", fontsize=8)

    ax.set_xticks(x)
    ax.set_xticklabels([labels[u] for u in ucs])
    ax.set_ylabel("Pass rate (%)")
    ax.set_ylim(0, 115)
    ax.set_title("Pass rate by business use case")
    ax.axhline(80, color=C_ACCENT, linestyle="--", linewidth=0.8, alpha=0.6, label="Target threshold (80%)")
    ax.legend(loc="lower right", framealpha=0.95)
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Figure: failure taxonomy (side-by-side donuts, run 1 vs run 2)
# ──────────────────────────────────────────────────────────────────────────
def fig_failure_taxonomy():
    cats_order = ["none", "missed_fact", "wrong_value", "hallucination", "bad_reasoning", "refused"]
    pretty = {
        "none": "Pass",
        "missed_fact": "Missed fact",
        "wrong_value": "Wrong value",
        "hallucination": "Hallucination",
        "bad_reasoning": "Bad reasoning",
        "refused": "Refused",
    }
    colors = {
        "none": C_OK,
        "missed_fact": "#E5A02E",
        "wrong_value": "#D17A3E",
        "hallucination": C_WARN,
        "bad_reasoning": "#7F4A7F",
        "refused": "#555555",
    }

    def prep(report):
        d = {f["category"]: f["count"] for f in report["failure_taxonomy"]}
        counts = [d.get(c, 0) for c in cats_order]
        labels = [pretty[c] for c in cats_order]
        return counts, labels

    c1, labels = prep(R1)
    c2, _      = prep(R2)

    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(7.5, 3.6))

    def make_donut(ax, counts, title, total):
        non_zero = [(i, c) for i, c in enumerate(counts) if c > 0]
        idxs = [i for i, _ in non_zero]
        vals = [c for _, c in non_zero]
        lbls = [labels[i] for i in idxs]
        cols = [colors[cats_order[i]] for i in idxs]
        wedges, _ = ax.pie(vals, labels=None, colors=cols, startangle=90,
                            wedgeprops=dict(width=0.42, edgecolor="white", linewidth=1.2))
        ax.set_title(title, y=-0.06)
        ax.text(0, 0.03, f"{total}", ha="center", va="center", fontsize=14, fontweight="bold")
        ax.text(0, -0.15, "runs", ha="center", va="center", fontsize=8, color="#555")
        # legend with counts
        leg_labels = [f"{lbl} ({v})" for lbl, v in zip(lbls, vals)]
        ax.legend(wedges, leg_labels, loc="center left", bbox_to_anchor=(1.0, 0.5),
                  fontsize=8, frameon=False)

    make_donut(ax1, c1, "Run 1 (baseline)", R1["summary"]["total_runs"])
    make_donut(ax2, c2, "Run 2 (refined)", R2["summary"]["total_runs"])
    fig.suptitle("Failure taxonomy across evaluation runs", y=1.02)
    fig.tight_layout()
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Figure: time saved per level (human vs agent)
# ──────────────────────────────────────────────────────────────────────────
def fig_time_savings():
    """Measured agent resolution time per complexity level (seconds)."""
    levels = ["L1", "L2", "L3", "L4", "L5"]
    per_level = {l["level"]: l["avg_duration_ms"] / 1000.0 for l in R2["by_level"]}
    vals = [per_level.get(l, 0) for l in levels]

    fig, ax = plt.subplots(figsize=(6.5, 3.2))
    x = np.arange(len(levels))
    bars = ax.bar(x, vals, color=C_IMPROVED, edgecolor="black", linewidth=0.5, width=0.55)

    for b, v in zip(bars, vals):
        ax.text(b.get_x() + b.get_width()/2, v + max(vals) * 0.02,
                f"{v:.0f} s", ha="center", va="bottom", fontsize=9, fontweight="bold")

    ax.set_xticks(x)
    ax.set_xticklabels([f"{l}\n{n}" for l, n in zip(
        levels, ["trivial", "simple", "moderate", "complex", "expert"])])
    ax.set_ylabel("Average agent resolution time (seconds)")
    ax.set_title("Measured agent resolution time per scenario (Run 2)")
    ax.set_ylim(0, max(vals) * 1.20)
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Figure: headline metrics dashboard
# ──────────────────────────────────────────────────────────────────────────
def fig_metrics_dashboard():
    metrics = [
        ("Pass rate",               "pass_rate",                   "%",  True),
        ("Weighted score",          "avg_weighted_score",          "/5", False),
        ("Functional consistency",  "functional_consistency_rate", "%",  True),
        ("Hallucination rate",      "hallucination_rate",          "%",  True),
    ]

    fig, axes = plt.subplots(1, 4, figsize=(8.5, 3.2))
    for ax, (name, key, unit, pct) in zip(axes, metrics):
        v1 = R1["summary"].get(key, 0)
        v2 = R2["summary"].get(key, 0)
        if pct:
            v1, v2 = v1 * 100, v2 * 100
        vals = [v1, v2]
        cols = [C_BASELINE, C_IMPROVED]
        bars = ax.bar(["Run 1", "Run 2"], vals, color=cols, edgecolor="black", linewidth=0.6, width=0.55)
        for b, v in zip(bars, vals):
            ax.text(b.get_x() + b.get_width()/2, v + (max(vals)*0.04 if max(vals) else 0.05),
                    (f"{v:.0f}{unit}" if unit == "%" else f"{v:.2f}{unit}"),
                    ha="center", va="bottom", fontsize=9, fontweight="bold")
        ax.set_title(name, fontsize=10)
        top = max(vals) * 1.25 if max(vals) > 0 else 1
        ax.set_ylim(0, top)
        ax.set_axisbelow(True)
        ax.tick_params(axis="x", labelsize=9)
        ax.tick_params(axis="y", labelsize=8)
        # remove y-ticks — values are annotated
        ax.set_yticks([])
        ax.spines["left"].set_visible(False)
        ax.spines["right"].set_visible(False)
        ax.spines["top"].set_visible(False)

    fig.suptitle("Headline metrics — baseline vs refined run", y=1.02, fontsize=11)
    fig.tight_layout()
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Figure: system coverage vs level (shows cross-system reasoning)
# ──────────────────────────────────────────────────────────────────────────
def fig_systems_touched():
    levels = ["L1", "L2", "L3", "L4", "L5"]
    r2 = {l["level"]: l.get("avg_systems_touched", 0) for l in R2["by_level"]}
    expected_max = 4  # 4 bounded contexts in total

    vals = [r2.get(l, 0) for l in levels]
    fig, ax = plt.subplots(figsize=(6.2, 3.2))
    bars = ax.bar(levels, vals, color=C_IMPROVED, edgecolor="black", linewidth=0.5, width=0.55)
    for b, v in zip(bars, vals):
        ax.text(b.get_x() + b.get_width()/2, v + 0.08, f"{v:.1f}", ha="center", va="bottom", fontsize=9)

    ax.axhline(expected_max, color=C_ACCENT, linestyle="--", linewidth=0.8, alpha=0.6)
    ax.text(0.1, expected_max + 0.08, "All 4 systems", color=C_ACCENT, fontsize=8)
    ax.set_ylim(0, expected_max + 0.6)
    ax.set_ylabel("Average systems touched per run")
    ax.set_title("Cross-system coverage by complexity level (run 2)")
    return fig


# ──────────────────────────────────────────────────────────────────────────
# Main
# ──────────────────────────────────────────────────────────────────────────
if __name__ == "__main__":
    figures = [
        ("score_by_level",      fig_score_by_level),
        ("pass_rate_by_uc",     fig_pass_rate_by_usecase),
        ("failure_taxonomy",    fig_failure_taxonomy),
        ("time_savings",        fig_time_savings),
        ("metrics_dashboard",   fig_metrics_dashboard),
        ("systems_touched",     fig_systems_touched),
    ]

    print("Generating figures...")
    for name, builder in figures:
        fig = builder()
        save(fig, name)
    print(f"\nDone. {len(figures)} figures in {FIG_DIR}")
