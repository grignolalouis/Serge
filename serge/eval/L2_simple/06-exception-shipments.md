# Shipment Exceptions

## Metadata
- **Level:** L2
- **Use Case:** UC-4
- **Systems:** TMS
- **Expected tools:** 1
- **Realistic user:** Logistics coordinator monitoring issues

## Question
"Any shipments with problems right now?"

## Required Facts

**Critical:**
- SHP-009 has exception status
- Reason: refrigeration unit malfunction on delivery vehicle

**Important:**
- Associated order: ORD-011
- Carrier: Kerry Express (CAR-001)

## Forbidden Content
- Invented exception reasons
- Successfully-delivered shipments listed as exceptions

## Flexibility Notes
- If only SHP-009 is in the scenario dataset, that's the answer
- Agent may volunteer that the carrier (ground) was inappropriate for refrigerated cargo — bonus

## Ground Truth

| Shipment | Order | Carrier | Status | Reason |
|----------|-------|---------|--------|--------|
| SHP-009 | ORD-011 | CAR-001 (Kerry Express) | exception | Refrigeration unit malfunction on delivery vehicle |
