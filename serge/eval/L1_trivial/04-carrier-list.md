# List Carriers

## Metadata
- **Level:** L1
- **Use Case:** UC-4
- **Systems:** TMS
- **Expected tools:** 1
- **Realistic user:** Logistics manager checking options

## Question
"What carriers do we work with?"

## Required Facts

**Critical:**
- Kerry Express (CAR-001), ground, 15 THB/kg, 2 days
- SCG Cold Chain (CAR-002), refrigerated, 25 THB/kg, 1 day
- Flash Express Premium (CAR-003), refrigerated, 35 THB/kg, 1 day

**Important:**
- The 3 carriers are listed (no missing ones)

## Forbidden Content
- Fabricated carriers
- Wrong costs or transit times

## Flexibility Notes
- Table format is natural but not required
- Any units for cost are acceptable (฿/kg, THB/kg, baht/kg)

## Ground Truth

| Carrier ID | Name | Type | Cost/kg | Transit days |
|-----------|------|------|---------|--------------|
| CAR-001 | Kerry Express | ground | 15 | 2 |
| CAR-002 | SCG Cold Chain | refrigerated | 25 | 1 |
| CAR-003 | Flash Express Premium | refrigerated | 35 | 1 |
