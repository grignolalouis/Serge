# [Scenario Title]

## Metadata
- **Level:** L[1-5]
- **Use Case:** UC-[1-5]
- **Systems:** [OMS, TMS, WMS, SRM]
- **Expected tools:** [count range, e.g. 3–5]
- **Realistic user:** [who would ask this and why]

## Question
"[Natural language question in the user's own words]"

## Required Facts

The answer must convey these facts. **Exact wording is flexible** — the judge checks for the factual content, not the phrasing.

**Critical (must be present to pass):**
- [Fact A with specific value]
- [Fact B with specific value]

**Important (expected, but can be implicit):**
- [Fact C]

**Nice to have (bonus, not required):**
- [Fact D]

## Forbidden Content

The answer must NOT contain:
- IDs or values not in the seed data
- [Specific context-appropriate forbidden claims]

## Flexibility Notes

The judge **should accept**:
- Different tool selection paths (e.g., using `list_all` + filter vs direct `get_by_id`)
- Any sensible ordering of tool calls
- Alternative phrasings and formats (tables, lists, prose)
- Extra context or analysis beyond what's required

The judge **should not penalize**:
- Minor unit/format variations (e.g., "฿450" vs "450 THB")
- Summarization vs full detail (both valid)
- Extra verification tool calls

## Ground Truth (for judge reference)

| Key | Value |
|-----|-------|
| [field] | [exact value from seed data] |
