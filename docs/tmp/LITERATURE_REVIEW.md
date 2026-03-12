# SERGE Project -- Literature Review for Midterm Report

> Last updated: 2026-03-11
> Covers 6 research areas with verified references from web search.

---

## 1. AI Agents: Architectures and Paradigms

### 1.1 Core Concepts

The modern LLM-agent paradigm decomposes autonomous systems into three pillars: **planning**, **memory**, and **tool use** (Weng, 2023). Agents interleave reasoning with action execution, grounding their decisions in external environments rather than relying solely on parametric knowledge.

### 1.2 Key References

**ReAct: Synergizing Reasoning and Acting in Language Models**
- Authors: Yao et al.
- Venue: ICLR 2023
- URL: https://arxiv.org/abs/2210.03629
- Contribution: Introduced the ReAct framework where LLMs generate interleaved *reasoning traces* and *task-specific actions*. Reasoning helps induce, track, and update action plans; actions allow interfacing with external sources (e.g., Wikipedia, knowledge bases). Evaluated on HotPotQA, FEVER, ALFWorld, and WebShop benchmarks, outperforming pure reasoning or pure acting baselines. ReAct also improves human interpretability and trustworthiness.

**Toolformer: Language Models Can Teach Themselves to Use Tools**
- Authors: Schick, Dwivedi-Yu, Dessi, Raileanu, Lomeli, Zettlemoyer, Cancedda, Scialom (Meta AI)
- Venue: NeurIPS 2023
- URL: https://arxiv.org/abs/2302.04761
- Contribution: Trained LLMs in a self-supervised fashion to decide which APIs to call, when to call them, what arguments to pass, and how to incorporate results. Covers calculator, Q&A systems, search engines, translation, and calendar APIs. Achieves competitive zero-shot performance with much larger models.

**HuggingGPT: Solving AI Tasks with ChatGPT and its Friends in Hugging Face**
- Authors: Shen et al. (Zhejiang University / Microsoft Research Asia)
- Venue: NeurIPS 2023
- URL: https://arxiv.org/abs/2303.17580
- Contribution: Four-stage workflow -- task planning, model selection, task execution, response generation -- using ChatGPT as a controller to orchestrate 24+ specialized AI models from Hugging Face. Demonstrates the "LLM as controller" paradigm for multi-model orchestration.

**LLM Powered Autonomous Agents (Blog Post)**
- Author: Lilian Weng (OpenAI)
- Date: June 2023
- URL: https://lilianweng.github.io/posts/2023-06-23-agent/
- Contribution: Highly-cited overview decomposing LLM agents into planning (CoT, ToT, task decomposition), memory (short-term / long-term / external), and tool use modules. References AutoGPT, GPT-Engineer, BabyAGI as case studies.

**Gorilla: Large Language Model Connected with Massive APIs**
- Authors: Patil et al. (UC Berkeley)
- Venue: NeurIPS 2024
- URL: https://arxiv.org/abs/2305.15334
- Contribution: A finetuned LLaMA model that surpasses GPT-4 on writing API calls across 1,600+ APIs. Introduced Retriever Aware Training (RAT) to adapt to test-time API documentation changes. Created APIBench and the Berkeley Function-Calling Leaderboard (BFCL).

**LLM-Agent-UMF: LLM-Based Agent Unified Modeling Framework**
- URL: https://arxiv.org/pdf/2409.11393
- Contribution: Proposes a modular architecture centered on a core-agent with five modules: planning, memory, profile, action, and security. Designed with a technology-agnostic approach enabling plug-and-play module composition.

**Survey: Large Language Model Agent: A Survey on Methodology, Applications and Challenges**
- URL: https://arxiv.org/abs/2503.21460
- Contribution: Covers 100+ papers (2020-2024), organizes agents into four categories: reasoning-enhanced, tool-augmented, multi-agent systems, and memory-augmented. Covers applications in software engineering, scientific research, robotics, and web automation.

**Survey: Agentic AI: A Comprehensive Survey of Architectures, Applications, and Future Directions**
- Venue: Artificial Intelligence Review, Springer, 2025
- URL: https://link.springer.com/article/10.1007/s10462-025-11422-4
- Contribution: Traces evolution from early LMs to the Agentic AI era (2022-present), covering AutoGPT, CrewAI, AutoGen paradigms. Identifies multi-agent collaboration as the current frontier.

### 1.3 Relevance to SERGE

SERGE adopts the ReAct-style reasoning-then-acting loop. As an LLM-agnostic agent, it aligns with the UMF framework's modular design principle: the core reasoning engine is decoupled from the LLM backend, while MCP servers serve as the standardized tool interface. The four-stage HuggingGPT pattern (plan -> select -> execute -> summarize) maps directly to SERGE's diagnostic query workflow.

---

## 2. LLMs and Tool Use: MCP and Function Calling

### 2.1 Core Concepts

Tool use in LLMs has evolved from prompt-based API insertion (Toolformer) to structured function calling (OpenAI, Anthropic) to the open **Model Context Protocol (MCP)**. The "Three Ws" of tool learning are: **Whether** (is a tool call necessary?), **Which** (which tool to invoke?), and **How** (what parameters, how to integrate results?).

### 2.2 Key References

**Model Context Protocol (MCP)**
- Organization: Anthropic / The Linux Foundation
- Released: November 2024
- URL: https://modelcontextprotocol.io/specification/2025-11-25
- Announcement: https://www.anthropic.com/news/model-context-protocol
- Contribution: Open standard for connecting LLM applications to external data sources and tools. Three core primitives: **tools** (callable functions), **resources** (data endpoints), **prompts** (reusable templates). SDKs available in Python, TypeScript, C#, Java. Architecture follows a client-server model where MCP servers expose capabilities and MCP clients (LLM applications) connect to them. Adopted by Replit, Sourcegraph, and major IDEs. Hosted by the Linux Foundation. By early 2025, 1,000+ MCP servers existed in the ecosystem.

**MCP Enterprise Adoption (2025)**
- Microsoft Dynamics 365 ERP MCP Server: launched at Build 2025 for Finance and Supply Chain Management, enabling agents to trigger journal entries, validate transactions, retrieve KPIs without custom code (https://www.microsoft.com/en-us/dynamics-365/blog/it-professional/2025/11/11/dynamics-365-erp-model-context-protocol/)
- Salesforce added native MCP support to Agentforce (July 2025)
- Oracle introduced MCP server support with native database integration
- MCP market projected at $1.8B in 2025

**LLM-Based Agents for Tool Learning: A Survey**
- Venue: Data Science and Engineering, Springer, 2025
- URL: https://link.springer.com/article/10.1007/s41019-025-00296-9
- Contribution: Identifies three dominant paradigms: (1) fine-tuning (Toolformer, Gorilla), (2) in-context learning (Chameleon), (3) orchestration frameworks (HuggingGPT). Documents performance degradation of >27 percentage points when forcing JSON output on reasoning tasks.

**Function Calling in LLMs: Industrial Practices, Challenges**
- URL: https://openreview.net/pdf/d01d50e27f7636724789f2aad6f4ac378749a0e1.pdf
- Contribution: Comprehensive taxonomy covering training data construction, model fine-tuning, deployment strategies, and evaluation across the full lifecycle.

**Berkeley Function-Calling Leaderboard (BFCL)**
- URL: https://gorilla.cs.berkeley.edu/blogs/8_berkeley_function_calling_leaderboard.html
- Contribution: 2,000 question-function-answer pairs across multiple languages and diverse application domains. Standard benchmark for evaluating LLM tool-calling accuracy.

**Unified Tool Integration for LLMs: A Protocol-Agnostic Approach**
- URL: https://arxiv.org/html/2508.02979v1
- Contribution: ToolRegistry -- a lightweight, modular integration layer that is protocol-agnostic, adapting to existing LLM applications while providing tool management abstractions.

### 2.3 Relevance to SERGE

MCP is the backbone of SERGE's architecture. Each supply chain system (SRM, WMS, OMS, TMS) is exposed through a dedicated MCP server, providing a standardized interface that any LLM client can connect to. This makes SERGE inherently LLM-agnostic: swapping the reasoning engine (Claude, GPT, Llama, etc.) requires no changes to the data access layer. The protocol-agnostic ToolRegistry concept validates SERGE's design of decoupling tool definitions from model-specific function-calling formats.

---

## 3. Supply Chain Management Problems

### 3.1 Core Concepts

B2B distribution supply chains are characterized by **fragmented systems** (ERP, WMS, OMS, TMS from different vendors), **data silos**, **lack of end-to-end visibility**, and **manual cross-system reconciliation**. The SCOR Digital Standard provides a process reference framework covering Plan, Source, Transform, Order, Fulfill, Return, and Orchestrate.

### 3.2 Key References

**Supply Chain Visibility Gap**
- Source: GEODIS Survey (cited by Descartes, 2024)
- URL: https://www.descartes.com/resources/knowledge-center/global-supply-chain-visibility-solutions-breaking-down-silos-b2b
- Finding: Almost two-thirds of respondents say their supply chain should be a competitive advantage, but **only 6% report having full supply chain visibility**. Nearly 80% have either no visibility or a restricted view.

**Cross-System Integration Pain Points (ERP/WMS/OMS/TMS)**
- Source: Supply & Demand Chain Executive / AWS, 2024
- URL: https://www.sdcexec.com/sourcing-procurement/procurement-software/article/22948489/amazon-web-services-aws-synchronizing-erp-wms-and-tms-systems-for-supply-chain-optimization
- Finding: Each system "speaks a different language" -- finance in ERP, orders in OMS, inventory in WMS, transport in TMS. Without integration, functions like transport ordering become manual (email + hand-keyed responses). Communication delays create confusion, errors, and inefficiencies that worsen at scale.

**Unified Dashboard: Connecting ERP, OMS, WMS, and TMS**
- Source: Omniful, 2026
- URL: https://www.omniful.ai/blog/unified-dashboard-erp-oms-wms-tms-supply-chain
- Finding: Most companies run separate standalone systems for each operation, creating data silos. Integration through unified dashboards can yield 5-10% freight cost savings.

**SCOR Digital Standard (SCOR-DS)**
- Organization: ASCM (formerly APICS/SCC)
- URL: https://www.ascm.org/corporate-solutions/standards-tools/scor-ds/
- Contribution: Updated in 2022 to address dynamic, asynchronous digital supply chains. Seven process elements with 19 emerging practices for digitization. Emphasizes collaboration, visibility, and market-driver effects. Platform-agnostic by design.

**Breaking Down Silos: Enhancing Supply Chain Visibility**
- Venue: IRJMETS, September 2024
- URL: https://www.irjmets.com/uploadedfiles/paper//issue_9_september_2024/61691/final/fin_irjmets1727091238.pdf
- Contribution: Academic paper specifically addressing how data silos hinder supply chain visibility and proposing integration approaches.

**B2B Distribution Challenges**
- Source: Globe3 / Trading & Distribution Industry Report
- URL: https://www.globe3.com/blog/article/challenges--opportunities-in-the-trading--distribution-industry
- Finding: Rise of marketplace platforms (Amazon) causing "supplier disintermediation." Distributors must differentiate through customized B2B offerings, data-driven insights, and operational excellence.

### 3.3 Relevance to SERGE

SERGE directly addresses the 94% visibility gap identified by GEODIS. By connecting to SRM, WMS, OMS, and TMS through MCP servers and providing a natural-language query interface, SERGE enables cross-system diagnostic queries without requiring full system integration or data warehouse consolidation. The read-only approach means no risk to operational systems while still breaking down information silos for decision support.

---

## 4. AI/LLM Applications in Supply Chain (2023-2026)

### 4.1 Core Concepts

The application of LLMs and AI agents to supply chain management is an emerging field (sharp uptick from <5 papers pre-2021 to 45 in 2024 alone). Key application areas include demand forecasting, inventory optimization, disruption monitoring, multi-agent consensus-seeking, and natural-language interfaces to supply chain data.

### 4.2 Key References

**Agentic LLMs in the Supply Chain: Towards Autonomous Multi-Agent Consensus-Seeking**
- Authors: (Published in International Journal of Production Research, 2025)
- URL: https://www.tandfonline.com/doi/full/10.1080/00207543.2025.2604311
- Also: https://arxiv.org/abs/2411.10184
- Contribution: Introduces supply chain-specific consensus-seeking frameworks where LLM agents represent companies, balancing selfish goals with systemic outcomes through structured conversation. Results show reduced bullwhip effects. When equipped with appropriate tools, LLM agents minimize bullwhip better than restocking policies and centralized demand approaches. Notes that academic agentic SC research is "in its infancy."

**Leveraging LLM-Based Agents for Intelligent Supply Chain Planning (SCPA)**
- Authors: (JD.com deployment)
- URL: https://arxiv.org/html/2509.03811v1
- Contribution: Constructs a Supply Chain Planning Agent (SCPA) framework that understands domain knowledge, comprehends operator needs, and decomposes tasks. Real-world deployment at JD.com demonstrates feasibility of LLM-agent applications in supply chain.

**LLMs in Supply Chain Management: Opportunities and a Case Study**
- Venue: ScienceDirect / IFAC, 2025
- URL: https://www.sciencedirect.com/science/article/pii/S2405896325012595
- Contribution: Explores opportunities for LLM integration into SCM with practical case study demonstrating effectiveness in sequential supply chain inventory replenishment tasks.

**Large Language Model-Driven Supply Chain Diagnosis Method**
- Venue: Springer, 2025
- URL: https://link.springer.com/chapter/10.1007/978-981-95-1106-8_14
- Contribution: Integrates LLM with optimization solver, multi-agent collaboration framework, and fault mode library to achieve automated supply chain diagnosis. Uses LLM to transform natural language into optimization code, automating the path from problem description to solution.

**Large Language Models for Supply Chain Decisions**
- URL: https://arxiv.org/abs/2507.21502
- Contribution: Identifies three key LLM applications for SC: (1) providing insights from data, (2) answering what-if questions, (3) updating SC tools to represent current business environments. Goal: automate decision-making without compromising decision quality.

**Systematic Analysis of Generative AI for Supply Chain Transformation**
- Venue: ScienceDirect, 2025
- URL: https://www.sciencedirect.com/science/article/pii/S2949863525000883
- Contribution: Systematic review of 98 peer-reviewed studies on GAI in SCM. First combined thematic and SCOR model-based mapping. Found ~45% of studies do not fully specify architectures (methodological immaturity). Most evidence remains at prototype level, rarely reporting system-wide KPIs.

**Automating Supply Chain Disruption Monitoring via an Agentic AI Approach**
- URL: https://arxiv.org/html/2601.09680v1
- Contribution: Knowledge Graph Query Agent translates diagnostic questions into graph database queries, mapping multi-tier supplier relationships affected by disruptions. Uses structured supply chain knowledge graph in Neo4j.

**Agentic Digital Twins: Bridging Model-Based and AI-Driven Decision-Making**
- Venue: International Journal of Production Research, 2026
- URL: https://www.tandfonline.com/doi/full/10.1080/00207543.2026.2630277
- Contribution: Bridges digital twin concept with agentic AI for supply chain decision support. Digital twins as real-time, composable data products for AI systems.

**On Implementing Autonomous Supply Chains: A Multi-Agent System Approach**
- Venue: ScienceDirect, 2024
- URL: https://www.sciencedirect.com/science/article/pii/S0166361524000484
- Contribution: Methodology for analysis and design of agent-based supply chain systems with practical implementation of autonomous supply chains.

**The Agentic Supply Chain (Special Issue)**
- Publisher: Taylor & Francis / International Journal of Production Research
- URL: https://think.taylorandfrancis.com/special_issues/agentic-supply-chain/
- Contribution: Dedicated special issue on the agentic supply chain, signaling growing academic recognition of this research area.

**Enterprise Adoption:**
- Walmart and Siemens experimenting with agentic LLMs for SC task automation
- SAP, Microsoft, Google developing supply chain agent platforms
- Microsoft Copilot for ERP/CRM launched March 2023

### 4.3 Relevance to SERGE

SERGE positions itself in the diagnostic/read-only segment of SC-AI, complementing the optimization and planning focus of most existing work. The LLM-driven diagnosis method (Springer, 2025) validates SERGE's core concept of using natural language to query and diagnose supply chain issues. The SCPA framework (JD.com) demonstrates real-world feasibility. The finding that ~45% of GAI-SC studies lack full architecture specification underscores the value of SERGE's explicit, documented MCP-based architecture.

---

## 5. LLM-as-Judge Evaluation

### 5.1 Core Concepts

LLM-as-Judge uses strong LLMs (e.g., GPT-4) to evaluate the outputs of other LLMs or systems, providing scalable, cost-effective, and consistent assessments as an alternative to human expert evaluation. Three evaluation modes: **pointwise** (rate one output), **pairwise** (compare two outputs), **listwise** (rank multiple outputs). Known biases include position bias, verbosity bias, and self-enhancement bias.

### 5.2 Key References

**Judging LLM-as-a-Judge with MT-Bench and Chatbot Arena**
- Authors: Zheng, Chiang et al.
- Venue: NeurIPS 2023 (Datasets and Benchmarks Track)
- URL: https://arxiv.org/abs/2306.05685
- Contribution: Foundational paper for the LLM-as-Judge paradigm. Introduces MT-Bench (multi-turn question set) and Chatbot Arena (crowdsourced battle platform). Strong LLM judges like GPT-4 achieve >80% agreement with human preferences. Examines biases: position, verbosity, self-enhancement, and limited reasoning. Proposes mitigation solutions. Publicly released 3K expert votes and 30K conversations with human preferences.

**A Survey on LLM-as-a-Judge**
- Authors: Gu et al., 2024
- URL: https://arxiv.org/abs/2411.15594
- Contribution: Comprehensive survey addressing how to build reliable LLM-as-Judge systems. Explores strategies to enhance consistency, mitigate biases, and adapt to diverse assessment scenarios.

**LLMs-as-Judges: A Comprehensive Survey on LLM-based Evaluation Methods**
- Authors: Li et al., December 2024
- URL: https://arxiv.org/abs/2412.05579
- GitHub: https://github.com/CSHaitao/Awesome-LLMs-as-Judges
- Contribution: Examines from five perspectives: Functionality, Methodology, Applications, Meta-evaluation, and Limitations. Covers the full taxonomy of evaluation approaches.

**From Generation to Judgment: Opportunities and Challenges of LLM-as-a-Judge**
- Authors: Li et al.
- Venue: EMNLP 2025
- URL: https://aclanthology.org/2025.emnlp-main.138.pdf
- Contribution: Introduces taxonomy along three dimensions: *what* to judge, *how* to judge, *where* to judge. Addresses bias mitigation, prompt engineering, and standardization of evaluation methodologies.

**Practical Guide: LLM-as-a-Judge**
- Source: Evidently AI
- URL: https://www.evidentlyai.com/llm-guide/llm-as-a-judge
- Contribution: Practical implementation guide for LLM-as-Judge evaluation pipelines.

### 5.3 Relevance to SERGE

SERGE's evaluation strategy uses LLM-as-Judge with well-defined ground truth (as per advisor feedback). For diagnostic queries across SRM/WMS/OMS/TMS, ground truth can be established from known data states, and an LLM judge can evaluate whether the agent's response correctly synthesizes cross-system information. The pairwise comparison mode from Zheng et al. is particularly relevant for comparing different LLM backends in SERGE's LLM-agnostic architecture. Key consideration: mitigating self-enhancement bias when the judge LLM is the same model being evaluated.

---

## 6. Research Gaps

### 6.1 Gap 1: No Standardized Supply Chain Agent Benchmarks

Current agent benchmarks (AgentBench, SWE-bench, WebArena) focus on software engineering, web navigation, and general reasoning. **No benchmark exists for evaluating AI agents operating across supply chain systems.**

**Supporting Evidence:**
- AgentBench (Liu et al., ICLR 2024): Evaluates LLMs across 8 environments (code, games, web) but no supply chain scenarios. URL: https://arxiv.org/abs/2308.03688
- Existing benchmarks "silo tasks" (text-only, vision-only, code-only), failing to evaluate the blended capabilities required for real-world SC workflows (source: https://o-mega.ai/articles/the-best-ai-agent-evals-and-benchmarks-full-2025-guide)
- Approximately 45% of GAI-SC studies do not fully specify architectures; most evidence at prototype level (source: https://www.sciencedirect.com/science/article/pii/S2949863525000883)
- "Academic agentic supply chain research is in its infancy, with research fragmented across computer science, operations research, and manufacturing engineering" (source: https://www.tandfonline.com/doi/full/10.1080/00207543.2025.2604311)

### 6.2 Gap 2: No Plug-and-Play, LLM-Agnostic SC Agent Architecture

Existing SC-AI solutions are tightly coupled to specific LLM providers or require custom integration code for each supply chain system.

**Supporting Evidence:**
- Most agent frameworks (LangChain, CrewAI, AutoGen) are general-purpose, not tailored for SC domain requirements
- Enterprise MCP adoption is recent (2025) and SC-specific MCP servers are not yet standardized
- The LLM-Agent-UMF proposes modular architecture but has not been applied to supply chain domains (URL: https://arxiv.org/pdf/2409.11393)
- Data readiness and governance gaps remain key hurdles (BCG, 2025: https://media-publications.bcg.com/The-Widening-AI-Value-Gap-October-2025.pdf)

### 6.3 Gap 3: Read-Only Diagnostic Agents Underexplored

Most SC-AI research focuses on **optimization** (planning, forecasting, replenishment) or **autonomous action** (consensus-seeking, automated ordering). The read-only diagnostic use case -- querying across systems to answer operational questions without modifying data -- is underexplored.

**Supporting Evidence:**
- LLM-driven SC diagnosis exists (Springer 2025) but focuses on automated optimization code generation, not read-only queries
- SCPA (JD.com) focuses on planning tasks, not cross-system diagnostic queries
- Digital twin approaches (query offloading for read-heavy analytics) are related but require full digital twin infrastructure

### 6.4 Gap 4: Lack of SC-Specific Evaluation with Ground Truth

LLM-as-Judge has been validated for general NLP tasks (MT-Bench, Chatbot Arena) but **no evaluation framework exists specifically for supply chain agent responses**.

**Supporting Evidence:**
- Zheng et al. (2023) validates LLM-as-Judge for open-ended conversation, not domain-specific factual SC queries
- SC-specific metrics (SCOR KPIs, order accuracy, inventory visibility) are not captured in existing LLM evaluation frameworks
- The need for domain-specific ground truth in SC evaluation is unaddressed in current literature

### 6.5 Relevance to SERGE

SERGE directly addresses Gaps 2, 3, and 4:
- **Gap 2**: SERGE provides an LLM-agnostic, plug-and-play architecture via MCP servers for SC systems
- **Gap 3**: SERGE is explicitly a read-only diagnostic agent, filling the underexplored niche
- **Gap 4**: SERGE's evaluation plan uses LLM-as-Judge with SC-specific ground truth, contributing a novel evaluation methodology

SERGE can also contribute toward closing Gap 1 by publishing its diagnostic query benchmark as a reusable SC agent evaluation dataset.

---

## Summary Table

| Area | Key Papers | Core Insight | SERGE Relevance |
|------|-----------|--------------|-----------------|
| AI Agents | Yao et al. (ReAct, 2023), Schick et al. (Toolformer, 2023), Shen et al. (HuggingGPT, 2023) | Interleaved reasoning + acting; self-supervised tool learning; LLM-as-controller | ReAct loop for diagnostic reasoning; modular architecture |
| LLMs & Tool Use | Anthropic (MCP, 2024), Patil et al. (Gorilla, 2024), BFCL | MCP standardizes LLM-tool connections; 1000+ servers by 2025 | MCP servers for SRM/WMS/OMS/TMS; LLM-agnostic by design |
| SC Problems | GEODIS survey, SCOR-DS, Omniful | 94% lack full visibility; siloed ERP/WMS/OMS/TMS | Breaks silos via cross-system queries |
| AI in SC | Agentic LLMs in SC (2025), SCPA/JD.com, SC Diagnosis (2025) | Field in infancy; bullwhip reduction; NL-to-optimization | Read-only diagnostic niche; validated by diagnosis research |
| LLM-as-Judge | Zheng et al. (MT-Bench, 2023), Gu et al. (Survey, 2024) | >80% agreement with humans; pointwise/pairwise/listwise | Pairwise comparison of LLM backends; SC ground truth |
| Research Gaps | No SC agent benchmarks; no plug-and-play SC architecture | 45% of GAI-SC papers lack architecture; fragmented research | SERGE fills gaps in architecture, diagnostics, evaluation |

---

## Sources

- [ReAct - Yao et al., ICLR 2023](https://arxiv.org/abs/2210.03629)
- [Toolformer - Schick et al., NeurIPS 2023](https://arxiv.org/abs/2302.04761)
- [HuggingGPT - Shen et al., NeurIPS 2023](https://arxiv.org/abs/2303.17580)
- [Gorilla - Patil et al., NeurIPS 2024](https://arxiv.org/abs/2305.15334)
- [LLM Powered Autonomous Agents - Weng, 2023](https://lilianweng.github.io/posts/2023-06-23-agent/)
- [LLM-Agent-UMF](https://arxiv.org/pdf/2409.11393)
- [LLM Agent Survey - Methodology, Applications, Challenges](https://arxiv.org/abs/2503.21460)
- [Agentic AI Survey - Springer 2025](https://link.springer.com/article/10.1007/s10462-025-11422-4)
- [Model Context Protocol - Anthropic](https://www.anthropic.com/news/model-context-protocol)
- [MCP Specification](https://modelcontextprotocol.io/specification/2025-11-25)
- [Dynamics 365 ERP MCP Server](https://www.microsoft.com/en-us/dynamics-365/blog/it-professional/2025/11/11/dynamics-365-erp-model-context-protocol/)
- [LLM Tool Learning Survey - Springer 2025](https://link.springer.com/article/10.1007/s41019-025-00296-9)
- [Berkeley Function-Calling Leaderboard](https://gorilla.cs.berkeley.edu/blogs/8_berkeley_function_calling_leaderboard.html)
- [Unified Tool Integration - Protocol-Agnostic](https://arxiv.org/html/2508.02979v1)
- [GEODIS Supply Chain Visibility](https://www.descartes.com/resources/knowledge-center/global-supply-chain-visibility-solutions-breaking-down-silos-b2b)
- [Synchronizing ERP, WMS, TMS - AWS](https://www.sdcexec.com/sourcing-procurement/procurement-software/article/22948489/amazon-web-services-aws-synchronizing-erp-wms-and-tms-systems-for-supply-chain-optimization)
- [SCOR Digital Standard - ASCM](https://www.ascm.org/corporate-solutions/standards-tools/scor-ds/)
- [Breaking Down Silos - IRJMETS 2024](https://www.irjmets.com/uploadedfiles/paper//issue_9_september_2024/61691/final/fin_irjmets1727091238.pdf)
- [Agentic LLMs in SC - IJPR 2025](https://www.tandfonline.com/doi/full/10.1080/00207543.2025.2604311)
- [SCPA - JD.com](https://arxiv.org/html/2509.03811v1)
- [LLM-Driven SC Diagnosis - Springer 2025](https://link.springer.com/chapter/10.1007/978-981-95-1106-8_14)
- [LLMs for SC Decisions](https://arxiv.org/abs/2507.21502)
- [Systematic Analysis of GAI for SC - ScienceDirect 2025](https://www.sciencedirect.com/science/article/pii/S2949863525000883)
- [SC Disruption Monitoring - Agentic AI](https://arxiv.org/html/2601.09680v1)
- [Agentic Digital Twins - IJPR 2026](https://www.tandfonline.com/doi/full/10.1080/00207543.2026.2630277)
- [Autonomous SC Multi-Agent - ScienceDirect 2024](https://www.sciencedirect.com/science/article/pii/S0166361524000484)
- [The Agentic Supply Chain - Special Issue](https://think.taylorandfrancis.com/special_issues/agentic-supply-chain/)
- [Zheng et al. - MT-Bench, NeurIPS 2023](https://arxiv.org/abs/2306.05685)
- [Survey on LLM-as-a-Judge - Gu et al. 2024](https://arxiv.org/abs/2411.15594)
- [LLMs-as-Judges Survey - Li et al. 2024](https://arxiv.org/abs/2412.05579)
- [LLM-as-Judge EMNLP 2025](https://aclanthology.org/2025.emnlp-main.138.pdf)
- [AgentBench - ICLR 2024](https://arxiv.org/abs/2308.03688)
- [AI Agent Benchmarks Guide 2025](https://o-mega.ai/articles/the-best-ai-agent-evals-and-benchmarks-full-2025-guide)
- [MCP Enterprise Integration - CData 2026](https://www.cdata.com/blog/2026-year-enterprise-ready-mcp-adoption)
- [BCG AI Value Gap 2025](https://media-publications.bcg.com/The-Widening-AI-Value-Gap-October-2025.pdf)
