# Reorder Point Analysis

## Metadata
- **Level:** L5
- **Use Case:** UC-5
- **Systems:** WMS, OMS, SRM
- **Expected tools:** 5–8
- **Realistic user:** Ops manager re-evaluating inventory policy after a stockout

## Question
"Our reorder point for PROD-001 is 20 units, but we still ran out. Is the threshold wrong, or is something else going on?"

## Required Facts

**Critical:**
- Current PROD-001 state: available 8, reorder point 20 — already below threshold
- Recent demand was abnormally high (ORD-007 alone requests 50 units — more than the reorder point itself)
- Supplier lead time: 3 days → during a 3-day lead time, demand of 50 units can exceed any reasonable buffer
- **Diagnosis:** reorder point of 20 is not wrong for average demand, but insufficient against a single urgent 50-unit order. Either the reorder point should rise OR demand of that size should trigger a different workflow (express procurement).

**Important:**
- Stock movement history shows consistent inbound/outbound flow
- Incoming PO-008 (150 units) exists — the system DID order more, but too late for ORD-007
- Suggestion: raise reorder point, or add a "critical order" alert when a single order exceeds X% of stock

**Nice to have:**
- Mention of SUP-001 being reliable (so the lead time can be trusted for new policy)
- Quantified suggestion (e.g., reorder point 40–50 given demand patterns)

## Forbidden Content
- Claim the reorder point is correct as-is without qualification
- Blame the supplier (they're reliable; the issue is policy)

## Flexibility Notes
- This is an analytical/policy question. Multiple reasonable conclusions exist.
- The judge focuses on: (a) correct factual picture and (b) coherent reasoning.
- No specific numeric recommendation is required — the logic matters.

## Ground Truth

| Input fact | Value |
|------------|-------|
| PROD-001 reorder point | 20 |
| PROD-001 available | 8 |
| Single-order demand (ORD-007) | 50 units |
| Supplier lead time | 3 days (SUP-001 and SUP-004) |
| Incoming PO-008 | 150 units, ETA Feb 18 |
| Conclusion direction | Reorder point is too low relative to realized large orders |
