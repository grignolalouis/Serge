# Operational Scan

## Metadata
- **Level:** L4
- **Use Case:** UC-4
- **Systems:** OMS, TMS, WMS, SRM
- **Expected tools:** 4–6
- **Realistic user:** Ops manager starting their day

## Question
"Give me a quick scan of everything that needs my attention right now."

## Required Facts

**Critical:**
At minimum, the agent must surface:
- **Overdue orders** (ORD-007, ORD-011)
- **Shipment exceptions** (SHP-009)
- **Low stock** (PROD-001, PROD-005, PROD-008)

**Important:**
- Each item has enough context to act on (IDs, severity, quick reason)
- Connection noted between ORD-007 being stuck and PROD-001 low stock

**Nice to have:**
- Prioritization (most urgent first)
- Overdue supplier orders mentioned if any

## Forbidden Content
- Invented issues
- Missing the known scenario problems

## Flexibility Notes
- Each system has a dedicated "list overdue/exception/low stock" tool — 4 parallel calls is the efficient path
- The agent may also pull details on specific items; bonus

## Ground Truth (scenario state)

| Category | Items |
|----------|-------|
| Overdue orders | ORD-007 (processing), ORD-011 (shipped/exception) |
| Exception shipments | SHP-009 (refrigeration failure) |
| Low stock | PROD-001 (8/20), PROD-005 (12/15), PROD-008 (6/15) |
| Key connection | ORD-007 is stuck because PROD-001 is short |
