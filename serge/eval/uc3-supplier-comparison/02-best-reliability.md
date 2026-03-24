# Best Reliability

## Use Case
UC-3: Supplier Comparison | Difficulty: easy

## Question
"Which supplier has the best delivery reliability?"

## Expected Tool Calls
1. `srm_compare_suppliers()` — get performance stats

## Ground Truth
- **SUP-001 (Somchai Mango Farm)** is the most reliable:
  - 100% on-time rate (3/3 received POs on time)
  - PO-001: expected Jan 5, received Jan 5 (on time)
  - PO-004: expected Jan 15, received Jan 14 (1 day early)
  - PO-006: expected Jan 25, received Jan 25 (on time)
  - Average lead time: 3 days (matches declared lead time)
  - Rating: 4.5/5.0
- SUP-003 (Prachuap Harvest): 50% on-time — middling
- SUP-002 (Northern Fruits Co.): 0% on-time — worst

## Must NOT Contain
- Wrong supplier identified as most reliable
- Incorrect PO delivery dates

## Evaluation Checklist
- [ ] SUP-001 correctly identified as most reliable
- [ ] 100% on-time rate stated
- [ ] Evidence provided (specific PO dates)
- [ ] Other suppliers ranked for context
- [ ] No hallucinated data
