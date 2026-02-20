# Week 2 - Progress Report: AI Agent for Supply Chain Operations

LOUIS GRIGNOLA
GAUTHIER SCHMITT
MAEL TALUS
MATTEO SCHWARZ

**Course:** SE618-1 Logistics
**Focus:** Agentic AI applied to Supply Chain Management

## This Week's Progress

### Research

We conducted research on supply chain modeling to better understand the domain we are building for: data flows, system boundaries (SRM, WMS, OMS, TMS), and how information silos create operational friction.

### Roadmap Definition

We established a clear roadmap for the coming weeks (see below).

### Tutor Meeting

We contacted and scheduled a first meeting with our tutor professor, Ajarn Warut (SIIT), to align on expectations and validate our approach.

## Roadmap

### Phase 1 - Supply Chain Modeling

Design and implement a simulated supply chain with realistic test data across the four systems (SRM, WMS, OMS, TMS). This fake but coherent dataset will serve as the foundation for all agent interactions.

### Phase 2 - Agent Approach Exploration

We will explore and compare **two approaches**:

| Approach | Description |
|----------|-------------|
| **A - Custom Agent** | Build our own complete agent using tRPC-Agent-Go, with full control over reasoning, tool routing, and orchestration |
| **B - MCP Integration** | Build MCP (Model Context Protocol) servers to expose supply chain data, and leverage existing agents (Claude, Cowork, Claude Code, etc.) to interact with it |

### Goal

Demonstrate that AI agents can perform **long-range, multi-step tasks** by:
- **Retrieving** and **finding** data across different siloed sources
- **Debugging** supply chain problems (e.g., delayed orders, stock discrepancies)
- Providing **read-only access** across all data systems
- **Reasoning** over combined data to produce actionable answers
