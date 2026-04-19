# Full Product Supply Chain History

## Metadata
- **Level:** L5
- **Use Case:** UC-4 / UC-2 (holistic trace)
- **Systems:** SRM, WMS, OMS
- **Expected tools:** 6–10
- **Realistic user:** Ops director preparing a review of a flagship product

## Question
"Give me the complete supply chain history for PROD-001 — from purchase orders to inventory to customer demand. What does it tell you?"

## Required Facts

**Critical:**
- **Product:** Nam Doc Mai Mango Box 5kg, premium, 450 THB, refrigerated, 7-day shelf life
- **Suppliers:** SUP-001 (280 THB) and SUP-004 (270 THB)
- **Recent POs:** PO-001, PO-004, PO-006, PO-008, PO-012 all tied to PROD-001 (from scenario data)
- **Current inventory:** 8 units available
- **Active demand:** ORD-007 (50 units), ORD-009 (5), ORD-023 (10) — combined exceeds stock
- **Finding:** stock is critically low; incoming PO-008 (150 units) arrives Feb 18

**Important:**
- Historical movements trace (inbound from POs, outbound to orders)
- Margin analysis (450 − 280 = 170 THB/unit ≈ 60% gross margin)
- Comment that the reorder point (20) may be too low given demand spikes

**Nice to have:**
- Visual representation (table, diagram, timeline)
- Proactive recommendation (increase reorder point, add SUP-004 as backup)

## Forbidden Content
- Wrong current stock value
- Missing one of the two suppliers
- Invented POs or orders

## Flexibility Notes
- This is a synthesis task — no single tool answers it. The agent must chain across all three systems.
- The judge does NOT require a specific tool order or count; the quality of the synthesis is what matters.
- Extra analysis (margin, trend, recommendations) is bonus but expected at this level.

## Ground Truth

| Dimension | Value |
|-----------|-------|
| Product | PROD-001 Nam Doc Mai Mango Box 5kg |
| Sell price | 450 THB |
| Suppliers | SUP-001 (280 THB), SUP-004 (270 THB) |
| Current available stock | 8 |
| Reorder point | 20 |
| Incoming resupply | PO-008 (SUP-001, 150 units, ETA 2026-02-18) |
| Active open demand | ~65 units across ORD-007, ORD-009, ORD-023 |
| Verdict | Severe stock shortage |
