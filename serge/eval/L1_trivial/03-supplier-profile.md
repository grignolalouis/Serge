# Supplier Profile

## Metadata
- **Level:** L1
- **Use Case:** UC-3
- **Systems:** SRM
- **Expected tools:** 1
- **Realistic user:** Procurement reviewing a supplier

## Question
"Give me the profile for supplier SUP-001."

## Required Facts

**Critical:**
- Name: Somchai Mango Farm
- Region: Chanthaburi
- Lead time: 3 days
- Rating: 4.5

**Important:**
- Product categories: premium, standard
- Payment terms: net_30

**Nice to have:**
- Contact email

## Forbidden Content
- Wrong rating or lead time
- Fabricated PO history (unless another tool is called)

## Flexibility Notes
- Agent may present as prose or table
- Agent may or may not comment on the rating being high

## Ground Truth

| Key | Value |
|-----|-------|
| id | SUP-001 |
| name | Somchai Mango Farm |
| region | Chanthaburi |
| lead_time_days | 3 |
| rating | 4.5 |
| payment_terms | net_30 |
| product_categories | premium, standard |
