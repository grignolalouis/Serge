# SERGE Evaluation Scenarios

15 test scenarios (5 per use case) for evaluating the SERGE supply chain agent.

Each scenario defines a natural language question, the expected MCP tool calls, the ground truth answer derived from seed data, and an evaluation checklist.

## Use Cases

| UC | Name | Systems | Complexity | Scenarios |
|----|------|---------|------------|-----------|
| 1 | Order Status Inquiry | OMS, TMS | Low | `uc1-order-status/` |
| 2 | Root Cause Analysis | OMS, TMS, WMS, SRM | Medium | `uc2-root-cause/` |
| 3 | Supplier Comparison | SRM | Low–Medium | `uc3-supplier-comparison/` |

## Evaluation Criteria (LLM-as-Judge)

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Factual correctness | 30% | All facts match ground truth |
| Completeness | 25% | All required data points covered |
| Reasoning quality | 20% | Causal chain is logical (UC-2 mainly) |
| Source attribution | 15% | Agent cites which system data comes from |
| Response structure | 10% | Output is organized and readable |

## Protocol

- 3 runs per scenario (45 total) to measure consistency
- Pass threshold: ≥ 3.5/5.0 weighted score
- Target: ≥ 4.0/5.0 average, ≥ 90% factual accuracy
