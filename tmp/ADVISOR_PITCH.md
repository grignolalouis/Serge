# Advisor Presentation: AI Agent for Supply Chain Teams

## Proposed Title
**"Agentic AI Assistant for Supply Chain Operations: A Case Study in Multi-Source Data Access and Decision Support"**

---

## The Problem (1 slide)

### Supply Chain Teams Face Information Overload
- Data scattered across **5-10+ systems** (ERP, WMS, TMS, CRM, Excel)
- Workers spend **2-3 hours/day** searching for information
- Cross-functional queries require IT support or manual aggregation
- Knowledge loss when experienced workers leave

### Example Pain Point
> A demand planner needs to understand why a stockout occurred. They must check:
> - Sales data (CRM) → Inventory records (WMS) → Supplier deliveries (TMS) → Production schedules (ERP)
>
> This takes hours and often yields incomplete answers.

---

## The Solution (1 slide)

### An AI Agent That Understands Supply Chain Context

**Natural Language Interface:**
```
User: "Why did we have a stockout of SKU-1234 in warehouse A last week?"

Agent: Based on my analysis:
- Demand exceeded forecast by 23% (CRM data)
- Reorder was triggered on Day 3 but supplier XYZ had a 5-day delay (TMS)
- Safety stock was set too low for this SKU's variability (WMS)

Recommendation: Increase safety stock to 150 units and add backup supplier.
```

### Key Capabilities
1. **Query** multiple supply chain data sources via natural language
2. **Synthesize** information from structured and unstructured data
3. **Reason** about cause-effect relationships
4. **Recommend** actions based on analysis

---

## Why Now? (Academic Relevance)

### Emerging Research Area
- First academic papers on "Agentic LLMs in Supply Chain" published 2024-2025
- JD.com deployed real system with **22% accuracy improvement**
- Gartner predicts 50% of SC solutions will have agentic AI by 2030

### Gap in Literature
- Most papers focus on **single-agent** or **specific functions**
- Limited research on **human-agent collaboration** in SC context
- Opportunity to contribute to an emerging field

---

## Proposed Methodology

### Phase 1: Literature Review
- Map existing AI/agent applications in supply chain
- Identify frameworks and architectures

### Phase 2: Problem Definition
- Select specific supply chain function (e.g., demand planning, logistics)
- Define user personas and use cases
- Establish success metrics

### Phase 3: Solution Design
- Design agent architecture (RAG + tool use)
- Define data sources and integration points
- Create interaction model

### Phase 4: Prototype & Evaluation
- Build proof-of-concept with synthetic/sample data
- Evaluate against defined metrics
- Document findings and limitations

---

## Questions for Advisor

1. **Scope:** Should I focus on one SC function or demonstrate cross-functional capability?
2. **Industry:** Any preference for industry vertical (retail, manufacturing, pharma)?
3. **Prototype:** Is a working demo required, or is architectural design sufficient?
4. **Collaboration:** Should I seek industry partner for validation?
5. **Timeline:** What milestones align with the course schedule?

---

## My Background Fit

- **Major in AI/GenAI/Agentic AI** → Technical depth for agent design
- **Supply Chain course** → Domain context and business understanding
- **Intersection** → Unique position to bridge both fields

---

## Requested Feedback

- Does this scope seem appropriate for the course?
- Any specific methodologies or frameworks you recommend?
- Suggestions for narrowing or expanding the focus?
