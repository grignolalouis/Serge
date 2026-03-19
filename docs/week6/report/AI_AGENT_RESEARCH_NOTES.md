# AI Agent Architectures: Research Notes for Literature Review

**Research Date:** 2026-03-12
**Purpose:** Literature review material for SE618-1 Week 6 Midterm Report

---

## 1. Large Language Models (LLMs) — Foundational Concepts

### 1.1 Transformer Architecture

The Transformer architecture was introduced by Vaswani et al. (2017) in "Attention Is All You Need," a paper that has since been cited over 173,000 times. The architecture relies entirely on self-attention mechanisms, dispensing with recurrence and convolutions. It enables parallel processing of input sequences by computing pairwise relationships between all tokens (n² attention), which is both its strength (capturing long-range dependencies) and its limitation (quadratic cost with sequence length).

**Citation:** Vaswani, A., Shazeer, N., Parmar, N., Uszkoreit, J., Jones, L., Gomez, A. N., Kaiser, L., & Polosukhin, I. (2017). Attention is all you need. *Advances in Neural Information Processing Systems*, 30, 5998–6008.
- Paper: https://arxiv.org/abs/1706.03762

### 1.2 Scaling Laws and Emergence

Kaplan et al. (2020) established that language model performance scales as a power-law with model size, dataset size, and compute, with trends spanning more than seven orders of magnitude. Key finding: larger models are significantly more sample-efficient; optimal compute-efficient training involves training very large models on relatively modest data and stopping before convergence.

Hoffmann et al. (2022) ("Chinchilla") later showed that Kaplan et al. underestimated the training data needed for compute-optimal models, leading to the "Chinchilla scaling laws."

Emergent abilities follow a phase-transition pattern: performance is near-random until a critical scale threshold, after which it increases substantially (Wei et al., 2022).

**Citations:**
- Kaplan, J., McCandlish, S., Henighan, T., et al. (2020). Scaling laws for neural language models. *arXiv:2001.08361*. https://arxiv.org/abs/2001.08361
- Hoffmann, J., Borgeaud, S., Mensch, A., et al. (2022). Training compute-optimal large language models. *arXiv:2203.15556*.

### 1.3 Key Models (as of early 2026)

| Model Family | Organization | Key Features |
|---|---|---|
| **GPT-4 / GPT-5** | OpenAI | Closed-source; GPT-5.2 has 400K context, 100% AIME 2025 score |
| **Claude (Opus 4.6, Sonnet 4)** | Anthropic | Constitutional AI; Sonnet 4 has 1M token context; Opus 4.6 (Feb 2026) is a new architecture focused on agentic performance |
| **Llama 4** | Meta | Open-source; first Llama with mixture-of-experts (MoE) architecture (April 2025) |
| **Mistral Large 3** | Mistral AI | Open-weight; 675B total params (MoE); European AI regulation compliance |

LLMs serve as the **reasoning engine** for AI agents. Without tools, they can only generate text. With tools, they become capable of perceiving environments, making decisions, taking actions, and learning from feedback.

---

## 2. Agent Architectures Beyond ReAct

### 2.1 ReAct (Baseline)

ReAct (Yao et al., 2022) interleaves reasoning and acting in a Thought → Action → Observation loop. Each step requires an LLM call. It is the foundational paradigm for tool-using agents.

**Limitation:** Single-path reasoning; each action requires a full LLM inference call; no separation between planning and execution.

### 2.2 Plan-and-Execute

Plan-and-Execute architectures separate reasoning from action: the agent plans the entire strategy first, then executes sequentially. Advantages over ReAct:
- **Faster execution**: sub-tasks can be performed without additional LLM calls (or with lighter-weight models)
- **Better task completion**: forces the planner to explicitly think through all required steps
- **Cost efficiency**: fewer LLM calls for multi-step workflows

**Source:** https://blog.langchain.com/planning-agents/

### 2.3 Tree of Thoughts (ToT) and Graph of Thoughts (GoT)

**Tree of Thoughts** (Yao et al., 2023) extends Chain-of-Thought by introducing a search over the space of thoughts. Instead of a single reasoning path, ToT generates and evaluates multiple possible next steps at each stage, organized as a tree structure where each node is a partial solution.

**Graph of Thoughts** (Besta et al., 2023) goes further, allowing arbitrary connections between thoughts — aggregation (combining best parts of several ideas) and iterative refinement loops.

**Integration:** Language Agent Tree Search (LATS) combines ToT with ReAct and planning in LangGraph.

### 2.4 ReWOO (Reasoning Without Observation)

Separates the planning phase entirely from the execution phase. The planner generates a full plan with placeholders for tool outputs, then all tools are executed, and finally results are synthesized. Reduces LLM calls significantly.

### 2.5 Modern Unified Taxonomy

Recent surveys (arXiv:2601.12560, Jan 2026) propose a unified taxonomy breaking agents into:
- **Perception** — input processing
- **Brain** — core reasoning
- **Planning** — strategy formation
- **Action** — execution
- **Tool Use** — external capabilities
- **Collaboration** — multi-agent coordination

Critical architectural dimensions now include: orchestration strategy, prompt implementation (ReAct vs function calling), memory architecture, and thinking tool integration.

**Key survey citations:**
- "Agentic Artificial Intelligence: Architectures, Taxonomies, and Evaluation of Large Language Model Agents" — arXiv:2601.12560 (Jan 2026). https://arxiv.org/html/2601.12560v1
- "A Review of Prominent Paradigms for LLM-Based Agents: Tool Use, Planning, and Feedback Learning" — CoLing 2025. https://arxiv.org/html/2406.05804v4
- "LLM-Based Agents for Tool Learning: A Survey" — Springer, 2025. https://link.springer.com/article/10.1007/s41019-025-00296-9
- Lilian Weng, "LLM Powered Autonomous Agents" (June 2023, updated). https://lilianweng.github.io/posts/2023-06-23-agent/

### 2.6 Test-Time Compute Scaling

A major 2025 insight: optimal scaling of test-time compute can be more effective than scaling model parameters. Extended thinking (deep reasoning at inference time) offers better generalization through enhanced flexibility in prompt and workflow design.

---

## 3. Claude Code as an Agent Architecture

### 3.1 Core Architecture: The Agentic Loop

Claude Code implements a **single-threaded master loop** with a minimal core pattern:

```
while(tool_call) → execute tool → feed results → repeat
```

The loop operates in three blended phases:
1. **Gather context** — search files, read code, understand the codebase
2. **Take action** — edit files, run commands, make changes
3. **Verify results** — run tests, check output, validate changes

Claude Code serves as the **agentic harness** around Claude: it provides the tools, context management, and execution environment that turn a language model into a capable agent.

**Key source:** https://code.claude.com/docs/en/how-claude-code-works

### 3.2 Tool Categories

| Category | Capabilities |
|---|---|
| **File operations** | Read, edit, create, rename files |
| **Search** | Glob (pattern matching), Grep (regex content search) |
| **Execution** | Bash commands, build tools, git, package managers |
| **Web** | WebSearch, WebFetch for documentation/research |
| **Code intelligence** | Type errors, jump to definitions, find references |

### 3.3 Context Engineering (Not Just Prompt Engineering)

Anthropic distinguishes **context engineering** from prompt engineering:
- **Prompt engineering** = crafting effective instructions for discrete tasks
- **Context engineering** = strategically managing ALL tokens during inference (system instructions, tools, external data, message history)

Key insight: models experience "context rot" — degradation in recall accuracy as context grows, because transformers create n² pairwise token relationships. Context is "a finite resource with diminishing marginal returns."

Techniques for managing context:
1. **Compaction** — summarize conversation history and reinitiate
2. **Structured note-taking** — agents maintain persistent notes outside the context window
3. **Sub-agent architectures** — focused tasks with clean context windows, returning condensed summaries

**Source:** https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents

### 3.4 Sub-Agent Spawning

Sub-agents are child Claude instances that a parent agent spawns for discrete tasks:
- Each sub-agent runs in its **own isolated context window**
- Does NOT share the parent's conversation history
- Receives a specific task, does its work, reports back a summary
- Keeps the main agent's context window clean
- Operates under **strict depth limitations** to prevent recursive spawning

Sub-agents inherit: loaded skills, project context (CLAUDE.md), file system access.
Sub-agents do NOT inherit: other sub-agent state, parent's working memory after spawn.

### 3.5 Skills System (SKILL.md)

Skills are organized folders of instructions, scripts, and resources that Claude discovers and loads dynamically. Introduced October 2025, they follow the Agent Skills open standard (https://agentskills.io).

**SKILL.md structure:**
- YAML frontmatter: `name`, `description`, `disable-model-invocation`, `allowed-tools`, `context`, `agent`, etc.
- Markdown body: instructions Claude follows when the skill is invoked

**Key design principle: Progressive Disclosure**
- Skill descriptions are loaded at session start (lightweight)
- Full skill content only loads when invoked (on-demand)
- This mirrors Claude Code's approach to context management

**Invocation modes:**
- **User-invoked**: `/skill-name` slash commands
- **Model-invoked**: Claude loads automatically when relevant
- **Fork mode**: `context: fork` runs in isolated sub-agent

**Bundled skills include:**
- `/simplify` — spawns 3 parallel review agents (code reuse, quality, efficiency)
- `/batch <instruction>` — decomposes into 5-30 independent units, spawns one agent per unit in git worktrees
- `/loop [interval] <prompt>` — recurring scheduled execution

**Source:** https://code.claude.com/docs/en/skills

### 3.6 Claude Agent SDK (formerly Claude Code SDK)

The Claude Agent SDK (renamed September 2025) is Anthropic's open-source framework that exposes the same infrastructure powering Claude Code as a programmable library.

**Timeline:**
- May 2025: Claude Code SDK announced alongside Claude Opus 4 / Sonnet 4
- September 2025: Renamed to Claude Agent SDK (general-purpose agent runtime)
- February 2026: Claude Opus 4.6 released (new architecture focused on agentic performance)

**Key capabilities via SDK:**
- Built-in tools: Read, Write, Edit, Bash, Glob, Grep, WebSearch, WebFetch
- Hooks: PreToolUse, PostToolUse, Stop, SessionStart, SessionEnd
- Sub-agents: define custom agents with specialized instructions
- MCP integration: connect to external systems
- Sessions: maintain context across exchanges, resume/fork

**Source:** https://platform.claude.com/docs/en/agent-sdk/overview

### 3.7 How Claude Code Goes Beyond ReAct

Claude Code is NOT a simple ReAct loop. Key differences:

1. **Integrated deep reasoning**: Extended thinking is built into the model, not bolted on as a prompting technique. The model reasons deeply before and during tool use.

2. **File-system-native**: The agent operates directly in a project directory with full file system access, git awareness, and terminal execution — not via abstract tool descriptions.

3. **Skill-based composition**: Reusable workflows (skills) can be discovered, loaded on demand, and composed. This is more like a library of expert behaviors than a fixed reasoning template.

4. **Sub-agent spawning with context isolation**: The orchestrator-worker pattern is built in, with strict depth limits and clean context boundaries.

5. **Real-time steering**: Users can interrupt and redirect mid-loop without restarting (via the asynchronous dual-buffer queue system).

6. **Context engineering over prompt engineering**: Sophisticated management of the finite context window, including compaction, structured notes, and progressive skill loading.

---

## 4. Anthropic's Agent Design Patterns

### 4.1 Seven Composable Patterns (December 2024)

From Anthropic's "Building Effective Agents" guide:

| Pattern | Type | Description |
|---|---|---|
| **Augmented LLM** | Building block | LLM enhanced with retrieval, tools, and memory |
| **Prompt Chaining** | Workflow | Sequential steps with programmatic gates between them |
| **Routing** | Workflow | Classify input, direct to specialized handlers |
| **Parallelization** | Workflow | Simultaneous work (sectioning or voting) |
| **Orchestrator-Workers** | Workflow | Central LLM dynamically delegates to workers |
| **Evaluator-Optimizer** | Workflow | One LLM generates, another evaluates iteratively |
| **Autonomous Agents** | Advanced | LLMs dynamically direct processes in a loop |

**Core principles:**
1. **Simplicity** — avoid unnecessary complexity; start with single LLM calls
2. **Transparency** — explicitly show planning steps
3. **ACI Excellence** — invest in agent-computer interface quality (tool documentation, testing)

> "Success in the LLM space isn't about building the most sophisticated system. It's about building the right system for your needs."

**Source:** https://www.anthropic.com/research/building-effective-agents

### 4.2 Agent-Computer Interface (ACI)

Anthropic emphasizes that developers should invest as much effort in Agent-Computer Interfaces (ACI) as in Human-Computer Interfaces (HCI). Good tool definitions include:
- Example usage
- Edge cases
- Input format requirements
- Clear boundaries from other tools

### 4.3 Writing Tools for Agents

Anthropic recommends tools that are:
- Self-contained and clear in purpose
- Token-efficient
- Minimal overlapping functionality
- Well-documented with examples

**Source:** https://www.anthropic.com/engineering/writing-tools-for-agents

---

## 5. Model Context Protocol (MCP)

### 5.1 Overview

MCP is an **open standard** introduced by Anthropic in **November 2024** to standardize how AI systems integrate with external tools, systems, and data sources. It solves the "N x M" integration problem (N models x M tools) by providing a universal adapter.

**Analogy:** MCP is to AI tools what USB-C is to peripherals — a universal connector.

### 5.2 Technical Architecture

- **Based on:** JSON-RPC 2.0
- **Inspired by:** Language Server Protocol (LSP)
- **Communication:** Stateful session protocol, bidirectional
- **Transport:** Supports both local (stdio) and remote (HTTP/SSE) mechanisms

**Primitives:**
- Servers expose: **Prompts**, **Resources**, **Tools**
- Clients expose: **Roots**, **Sampling**

### 5.3 Design Principles

1. **Security**: Host controls client connection permissions
2. **Two-way communication**: AI can receive information AND trigger actions
3. **Standardization**: Any compliant model works with any compliant tool
4. **Modularity**: MCP servers can be developed/deployed independently

### 5.4 Adoption

Following its November 2024 announcement, MCP was adopted by major AI providers including OpenAI and Google DeepMind. SDKs available in Python, TypeScript, C#, and Java.

**Specification:** https://modelcontextprotocol.io/specification/2025-11-25

**Sources:**
- Announcement: https://www.anthropic.com/news/model-context-protocol
- Wikipedia: https://en.wikipedia.org/wiki/Model_Context_Protocol
- GitHub: https://github.com/modelcontextprotocol/modelcontextprotocol

---

## 6. Anthropic Research Papers and References

### 6.1 Key Publications

| Title | Date | URL |
|---|---|---|
| Constitutional AI: Harmlessness from AI Feedback | Dec 2022 | https://arxiv.org/abs/2212.08073 |
| Building Effective Agents | Dec 2024 | https://www.anthropic.com/research/building-effective-agents |
| Effective Context Engineering for AI Agents | 2025 | https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents |
| Effective Harnesses for Long-Running Agents | 2025 | https://www.anthropic.com/engineering/effective-harnesses-for-long-running-agents |
| Writing Tools for Agents | 2025 | https://www.anthropic.com/engineering/writing-tools-for-agents |
| Equipping Agents with Agent Skills | 2025 | https://www.anthropic.com/engineering/equipping-agents-for-the-real-world-with-agent-skills |
| Introducing Advanced Tool Use | 2025 | https://www.anthropic.com/engineering/advanced-tool-use |
| Measuring AI Agent Autonomy in Practice | 2025 | https://www.anthropic.com/research/measuring-agent-autonomy |
| Building Agents with the Claude Agent SDK | 2025 | https://www.anthropic.com/engineering/building-agents-with-the-claude-agent-sdk |
| Code Execution with MCP | 2025 | https://www.anthropic.com/engineering/code-execution-with-mcp |

### 6.2 Tool Use Documentation

- Tool use overview: https://platform.claude.com/docs/en/agents-and-tools/tool-use/overview
- How to implement tool use: https://platform.claude.com/docs/en/agents-and-tools/tool-use/implement-tool-use
- Programmatic tool calling: https://platform.claude.com/docs/en/agents-and-tools/tool-use/programmatic-tool-calling
- Agent loop documentation: https://platform.claude.com/docs/en/agent-sdk/agent-loop

### 6.3 Advanced Tool Use Features (2025)

Three beta features for dynamic tool discovery and execution:
1. **Tool Search Tool**: Claude uses search to access thousands of tools without consuming context window
2. **Programmatic Tool Calling**: Claude writes code that calls tools in a code execution container (reduces latency and tokens)
3. **Learning from examples**: Claude learns tool usage patterns from demonstrations

---

## 7. Additional Academic References

### Agent Surveys and Taxonomies

- **Agentic AI: Architectures, Taxonomies, and Evaluation** (Jan 2026). arXiv:2601.12560. https://arxiv.org/html/2601.12560v1
- **Agentic AI: A Comprehensive Survey** (Oct 2025). arXiv:2510.25445. https://arxiv.org/html/2510.25445v1
- **The Rise of Agentic AI: Definitions, Frameworks, Architectures** (2025). MDPI Future Internet, 17(9), 404. https://www.mdpi.com/1999-5903/17/9/404
- **AgentArch: Benchmark for Agent Architectures in Enterprise** (Sep 2025). arXiv:2509.10769. https://arxiv.org/html/2509.10769v1
- **Agentic AI Frameworks: Architectures, Protocols, Design Challenges** (Aug 2025). arXiv:2508.10146. https://arxiv.org/html/2508.10146v1
- **LLM-based Agentic Reasoning Frameworks: A Survey** (Aug 2025). arXiv:2508.17692. https://arxiv.org/html/2508.17692v1
- **Pre-Act: Multi-Step Planning and Reasoning Improves Acting** (May 2025). arXiv:2505.09970. https://arxiv.org/html/2505.09970v2

### Foundational Agent Papers

- **ReAct**: Yao, S., et al. (2022). ReAct: Synergizing reasoning and acting in language models. arXiv:2210.03629.
- **Tree of Thoughts**: Yao, S., et al. (2023). Tree of thoughts: Deliberate problem solving with large language models. arXiv:2305.10601.
- **Graph of Thoughts**: Besta, M., et al. (2023). Graph of thoughts: Solving elaborate problems with large language models. arXiv:2308.09687.

### Industry Trends

- Gartner reported a **1,445% surge** in multi-agent system inquiries from Q1 2024 to Q2 2025
- Google's **Agent2Agent Protocol (A2A)** complements MCP with capabilities for memory management, goal coordination, task invocation, and capability discovery

---

## 8. Key Takeaways for SERGE Project

1. **Architecture choice validated**: Claude Code's agentic loop pattern (gather context → take action → verify) maps directly to SERGE's use case of querying SC systems, reasoning across them, and generating reports.

2. **MCP is the right integration layer**: MCP was designed exactly for the N x M problem SERGE faces (1 agent x 4 SC systems). It is now an industry standard adopted by all major AI providers.

3. **Skills enable domain expertise**: SERGE can package SC-specific workflows as skills (e.g., `/analyze-stockout`, `/supplier-risk-report`) that encode expert knowledge and load on demand.

4. **Sub-agents for parallel analysis**: The sub-agent pattern enables SERGE to query multiple SC systems in parallel (SRM + WMS + OMS + TMS) with isolated context, then synthesize findings.

5. **Read-only by design**: Claude Code's permission system supports read-only tool configurations (`allowed-tools: Read, Grep, Glob`), which aligns perfectly with SERGE's "never write to SC databases" constraint.

6. **Context engineering matters**: For a supply chain agent dealing with large datasets (500 orders, 400 shipments), progressive skill loading and sub-agent context isolation are essential to avoid context rot.

7. **Evaluation**: Anthropic's Evaluator-Optimizer pattern and LLM-as-judge approach provide a framework for evaluating SERGE's outputs against ground truth.
