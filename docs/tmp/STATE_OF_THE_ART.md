# State of the Art: Agentic AI in Supply Chain Management and Logistics

## Executive Summary

The application of agentic AI -- particularly Large Language Model (LLM)-based multi-agent systems -- to supply chain management (SCM) is a rapidly emerging research area. While industry players such as Walmart, Siemens, SAP, and JD.com are already experimenting with or deploying agentic supply chain solutions, academic research remains in its early stages, fragmented across computer science, operations research, and manufacturing engineering. This review synthesizes 14 relevant research papers published between 2023 and 2026 that explore LLM agents, multi-agent systems, retrieval-augmented generation (RAG), and related approaches for supply chain optimization, inventory management, disruption monitoring, and warehouse operations.

---

## 1. Catalogue of Research Papers (2023--2026)

### Paper 1: Agentic LLMs in the Supply Chain: Towards Autonomous Multi-Agent Consensus-Seeking

| Field | Details |
|-------|---------|
| **Authors** | Valeria Jannelli, Stefan Schoepf, Matthias Bickel, Torbjorn Netland, Alexandra Brintrup |
| **Year** | 2024 (preprint Nov 2024); published 2025 |
| **Venue** | arXiv:2411.10184 / *International Journal of Production Research* (Taylor & Francis), DOI: 10.1080/00207543.2025.2604311 |
| **URL** | [arXiv](https://arxiv.org/abs/2411.10184) -- [Taylor & Francis](https://www.tandfonline.com/doi/full/10.1080/00207543.2025.2604311) |

**Contribution:** Introduces supply chain-specific consensus-seeking frameworks for LLM agents. Multiple agents represent different companies in a supply chain and balance selfish goals with systemic outcomes through natural language conversation. The work demonstrates that LLM-based consensus-seeking reduces the bullwhip effect and that LLM agents equipped with appropriate tools can minimize bullwhip better than traditional restocking policies and centralized demand approaches.

**Evaluation:** Case study in inventory management using a multi-echelon supply chain simulation (beer distribution game setting). Metrics include bullwhip effect ratio, total supply chain cost, and consensus convergence rate. Code is open-sourced.

---

### Paper 2: Rethinking Supply Chain Planning: A Generative Paradigm (Leveraging LLM-Based Agents for Intelligent Supply Chain Planning)

| Field | Details |
|-------|---------|
| **Authors** | Jiaheng Yin, Yongzhi Qi, Jianshen Zhang, Dongyang Geng, Zhengyu Chen, Hao Hu, Wei Qi, Zuo-Jun Max Shen |
| **Year** | 2025 (submitted Sep 2025; revised Jan 2026) |
| **Venue** | arXiv:2509.03811 |
| **URL** | [arXiv](https://arxiv.org/abs/2509.03811) |

**Contribution:** Proposes the Supply Chain Planning Agent (SCPA) framework, a GenAI-powered agentic system that bridges strategic intent with operational action. The framework integrates semantic interpretation, autonomous task decomposition, chain-of-thought reasoning, and iterative refinement. Deployed within JD.com's large-scale retail supply chain operations covering millions of SKUs.

**Evaluation:** Real-world deployment at JD.com. Key metrics: approximately 22% improvement in plan accuracy, 2--3% improvement in stock fulfillment rates, approximately 40% reduction in weekly data processing time. Evaluation is based on operational deployment rather than synthetic benchmarks.

---

### Paper 3: Automating Supply Chain Disruption Monitoring via an Agentic AI Approach

| Field | Details |
|-------|---------|
| **Authors** | Sara AlMahri, Liming Xu, Alexandra Brintrup |
| **Year** | 2026 (submitted Jan 2026) |
| **Venue** | arXiv:2601.09680 / also on SSRN |
| **URL** | [arXiv](https://arxiv.org/abs/2601.09680) |

**Contribution:** Introduces a minimally supervised agentic AI framework with seven specialized agents powered by LLMs and deterministic tools. The agents jointly detect disruption signals from unstructured news, map them to multi-tier supplier networks, evaluate exposure based on network structure, and recommend mitigations such as alternative sourcing. Addresses the critical problem of limited visibility beyond Tier-1 suppliers.

**Evaluation:** Tested across 30 synthesized scenarios covering three automotive manufacturers and five disruption classes. F1 scores between 0.962 and 0.991 across core analytical tasks. End-to-end analyses completed in mean 3.83 minutes at $0.0836 per disruption event -- a reduction of more than three orders of magnitude in response time compared to traditional multi-day analyst-driven assessments. Real-world validation via case study on the 2022 Russia-Ukraine conflict and its impact on the automotive sector.

---

### Paper 4: AI Agent Systems for Supply Chains: Structured Decision Prompts and Memory Retrieval

| Field | Details |
|-------|---------|
| **Authors** | Konosuke Yoshizato, Kazuma Shimizu, Ryota Higa, Takanobu Otsuka |
| **Year** | 2026 (submitted Feb 2026) |
| **Venue** | arXiv:2602.05524 / Accepted at AAMAS 2026 (25th International Conference on Autonomous Agents and Multiagent Systems) |
| **URL** | [arXiv](https://arxiv.org/abs/2602.05524) |

**Contribution:** Investigates LLM-based multi-agent systems (MAS) for multi-echelon inventory management. Introduces AIM-RM (Agent for Inventory Management with Retrieval Memory), which uses similarity matching to access historical ordering experiences. This is the first study to demonstrate that an LLM-based MAS can achieve optimal solutions for multi-echelon inventory management problems without requiring prompt adjustments. Encodes fixed-ordering strategy prompts with stepwise processes and safe-stock strategies.

**Evaluation:** Empirical experiments across various supply chain scenarios comparing AIM-RM against benchmark methods (including traditional ordering policies). Results demonstrate optimal ordering decisions in restricted scenarios and superior performance of AIM-RM across diverse supply chain configurations.

---

### Paper 5: InvAgent: A Large Language Model based Multi-Agent System for Inventory Management in Supply Chains

| Field | Details |
|-------|---------|
| **Authors** | Yinzhu Quan, Zefang Liu |
| **Year** | 2024 (submitted Jul 2024; revised Jan 2025) |
| **Venue** | arXiv:2407.11384 |
| **URL** | [arXiv](https://arxiv.org/abs/2407.11384) -- [GitHub](https://github.com/zefang-liu/InvAgent) |

**Contribution:** Presents InvAgent, a zero-shot multi-agent inventory management system. Each agent manages a different node in the supply chain using LLMs for reasoning, decision-making, and inter-agent communication. Key innovations include: zero-shot learning without prior training data, chain-of-thought explainability, and dynamic adaptability to varying demand scenarios. Demonstrates that LLM-based agents can handle inventory decisions without reinforcement learning fine-tuning.

**Evaluation:** Extensive evaluations across multiple demand scenarios (stable demand, variable demand, seasonal patterns). Metrics: inventory holding costs, stockout rates, total supply chain cost, and adaptability to demand fluctuations. Compared against heuristic baselines (e.g., (s,S) policies, base-stock policies). Code is open-sourced on GitHub.

---

### Paper 6: Agentic AI Sustainability Assessment for Supply Chain Document Insights

| Field | Details |
|-------|---------|
| **Authors** | Diego Gosmar, Anna Chiara Pallotta, Giovanni Zenezini |
| **Year** | 2025 (submitted Nov 2025) |
| **Venue** | arXiv:2511.07097 |
| **URL** | [arXiv](https://arxiv.org/abs/2511.07097) |

**Contribution:** Proposes an ESG-oriented sustainability assessment framework for agentic AI applied to supply chain document processing. Compares three scenarios: fully manual (human-only), AI-assisted (human-in-the-loop/HITL), and multi-agent agentic AI workflows with parsers and verifiers. Demonstrates that agentic AI can achieve dramatic reductions in resource consumption while maintaining output quality.

**Evaluation:** Comparative assessment across three operational scenarios. Metrics: energy consumption (70--90% reduction), CO2 emissions (90--97% reduction), water usage (89--98% reduction) versus manual processes. Includes a complete replicability use case with real-world document extraction tasks. Framework integrates performance, energy, and emission indicators into a unified ESG methodology.

---

### Paper 7: From Unstructured Communication to Intelligent RAG: Multi-Agent Automation for Supply Chain Knowledge Bases

| Field | Details |
|-------|---------|
| **Authors** | Yao Zhang, Zaixi Shang, Silpan Patel, Mikel Zuniga |
| **Year** | 2025 (submitted Jun 2025) |
| **Venue** | arXiv:2506.17484 / 1st Workshop on AI for Supply Chain @ ACM SIGKDD 2025 (KDD '25), Toronto, Canada |
| **URL** | [arXiv](https://arxiv.org/abs/2506.17484) -- [Amazon Science](https://www.amazon.science/publications/from-unstructured-communication-to-intelligent-rag-multi-agent-automation-for-supply-chain-knowledge-bases) |

**Contribution:** Introduces an offline-first methodology that transforms unstructured supply chain communications (support tickets, emails, chat logs) into structured knowledge bases. Uses a multi-agent system with three specialized agents: Category Discovery (taxonomy creation), Categorization (ticket grouping), and Knowledge Synthesis (article generation). The prebuilt knowledge base is then used for RAG-based question answering.

**Evaluation:** Real-world support tickets with resolution notes from supply chain operations. Metrics: helpful answer rate (48.74% vs. 38.60% for traditional RAG), 77.4% reduction in unhelpful responses, knowledge base compression to 3.4% of original data volume, approximately 50% automatic resolution rate for future tickets.

---

### Paper 8: Leveraging Knowledge Graphs and LLM Reasoning to Identify Operational Bottlenecks for Warehouse Planning Assistance

| Field | Details |
|-------|---------|
| **Authors** | Rishi Parekh, Saisubramaniam Gopalakrishnan, Zishan Ahmad, Anirudh Deodhar |
| **Year** | 2025 (submitted Jul 2025) |
| **Venue** | arXiv:2507.17273 |
| **URL** | [arXiv](https://arxiv.org/abs/2507.17273) |

**Contribution:** Proposes a framework integrating Knowledge Graphs (KGs) with LLM-based agents to analyze Discrete Event Simulation (DES) output data from warehouse operations. Transforms raw DES data into a semantically rich KG capturing relationships between simulation events and entities. The LLM agent uses iterative reasoning, generates interdependent sub-questions, creates Cypher queries for KG interaction, and self-reflects to correct errors -- mimicking human expert analysis.

**Evaluation:** Tested on warehouse scenarios with equipment breakdowns and process irregularities. Two question categories: operational questions (near-perfect pass rates in pinpointing inefficiencies) and complex investigative questions (superior diagnostic ability for subtle, interconnected issues). Compared against baseline methods for DES analysis.

---

### Paper 9: Integrating RAG (Retrieval-Augmented Generation) with Prompt Engineering for Knowledge-Driven Supply Chain Solutions

| Field | Details |
|-------|---------|
| **Authors** | (Authors not fully confirmed via open access; published on ResearchGate) |
| **Year** | 2025 (published Jul 2025) |
| **Venue** | ResearchGate publication (ID: 393479215) |
| **URL** | [ResearchGate](https://www.researchgate.net/publication/393479215_Integrating_RAG_Retrieval-Augmented_Generation_with_Prompt_Engineering_for_Knowledge-Driven_Supply_Chain_Solutions) |

**Contribution:** Presents a hybrid framework where prompt-engineered queries guide retrieval from structured and unstructured supply chain data. The RAG system dynamically accesses external knowledge bases to enhance factual accuracy and contextual relevance, while prompt engineering ensures domain-specific reasoning. Addresses automated compliance monitoring, inventory anomaly detection, and logistical disruption response.

**Evaluation:** Experimental evaluations using real-world datasets. Metrics: decision latency, contextual precision, and system adaptability. Demonstrates improvements across all measured dimensions relative to baseline LLM approaches.

---

### Paper 10: LLMs in Supply Chain Management: Opportunities and a Case Study

| Field | Details |
|-------|---------|
| **Authors** | Ge Zheng, Sara AlMahri, Liming Xu, Maria Minaricova |
| **Year** | 2025 |
| **Venue** | IFAC-PapersOnLine (11th IFAC Conference on Manufacturing Modelling, Management and Control -- MIM 2025, Trondheim, Norway) / ScienceDirect |
| **URL** | [ScienceDirect](https://www.sciencedirect.com/science/article/pii/S2405896325012595) |

**Contribution:** Provides a comprehensive exploration of LLM potential in SCM, examining adoption challenges and opportunities. Presents a pipeline case study demonstrating the integration of LLMs with a decentralized agent-based system for SCM task execution, using a delivery delay prediction example. Discusses how LLMs can integrate with real-time market monitoring systems for dynamic forecast adjustment.

**Evaluation:** Pipeline case study with delivery delay prediction. Qualitative assessment of LLM integration feasibility with existing decentralized agent-based supply chain systems.

---

### Paper 11: Large Language Models for Supply Chain Optimization (OptiGuide)

| Field | Details |
|-------|---------|
| **Authors** | Beibin Li, Konstantina Mellou, Bo Zhang, Jeevan Pathuri, Ishai Menache |
| **Year** | 2023 (submitted Jul 2023) |
| **Venue** | arXiv:2307.03875 |
| **URL** | [arXiv](https://arxiv.org/abs/2307.03875) |

**Contribution:** Introduces OptiGuide, a framework that bridges LLMs with combinatorial optimization for supply chain management. The system accepts plain text queries and outputs insights about optimization outcomes, enabling non-technical users to interact with complex optimization models. Preserves data privacy by avoiding transmission of proprietary data to LLMs. Addresses what-if analysis such as cost implications of switching suppliers.

**Evaluation:** General evaluation benchmark developed for assessing LLM output accuracy. Real-world case study on server placement within Microsoft's cloud supply chain. Metrics focused on accuracy of natural language interpretation and correctness of optimization insights.

---

### Paper 12: Enhancing Supply Chain Resilience with Multi-Agent Systems and Machine Learning: A Framework for Adaptive Decision-Making

| Field | Details |
|-------|---------|
| **Authors** | (Published in The American Journal of Engineering and Technology) |
| **Year** | 2025 |
| **Venue** | *The American Journal of Engineering and Technology* |
| **URL** | [Journal](https://theamericanjournals.com/index.php/tajet/article/view/5919) -- [ResearchGate](https://www.researchgate.net/publication/389525426_Enhancing_supply_chain_resilience_with_multi-agent_systems_and_machine_learning_a_framework_for_adaptive_decision-making) |

**Contribution:** Proposes a MAS-ML (Multi-Agent Systems coupled with Machine Learning) framework for adaptive decision-making in supply chains. Each agent handles a specific supply chain function (demand forecasting, inventory management, production planning, logistics) while leveraging real-time ML-driven predictions. The decentralized architecture enables resilience against demand variability and transportation disruptions.

**Evaluation:** Framework evaluation through simulation scenarios involving demand variability, transportation disruptions, and pandemic-like events. Metrics: supply chain flexibility, adaptability, cost optimization, and predictiveness.

---

### Paper 13: Intelligent Human-Machine Partnership for Manufacturing: Enhancing Warehouse Planning through Simulation-Driven Knowledge Graphs and LLM Collaboration

| Field | Details |
|-------|---------|
| **Authors** | (See arXiv listing) |
| **Year** | 2025 (submitted Dec 2025) |
| **Venue** | arXiv:2512.18265 |
| **URL** | [arXiv](https://arxiv.org/abs/2512.18265) |

**Contribution:** Extends the KG+LLM approach for warehouse planning by incorporating simulation-driven knowledge graphs and human-machine collaboration. Focuses on how LLM agents can partner with human warehouse planners through simulation-driven insights, enabling more effective decision-making in complex manufacturing warehouse environments.

**Evaluation:** Warehouse planning scenarios using simulation-driven data integrated into knowledge graphs. Assessed through comparison of human-AI collaborative planning outcomes versus manual-only approaches.

---

### Paper 14: Supply Chain Mapping through Retrieval-Augmented Generation: Applications to the Electronics Industry

| Field | Details |
|-------|---------|
| **Authors** | (See Taylor & Francis listing; associated with Brintrup research group) |
| **Year** | 2026 (published online Jan 2026) |
| **Venue** | *Journal of the Operational Research Society* (Taylor & Francis), DOI: 10.1080/01605682.2025.2608868 |
| **URL** | [Taylor & Francis](https://www.tandfonline.com/doi/full/10.1080/01605682.2025.2608868) |

**Contribution:** Presents a novel methodology for automated multi-tier supply chain mapping using RAG and network science. The RAG-based approach extracts supplier-customer relationships from unstructured public data sources including SEC 10-K filings and earnings calls. Extracted entities are structured into directed supply chain graphs and analyzed using network science metrics (centrality, modularity, path length).

**Evaluation:** Applied to the electronics industry. Data sources include SEC filings and earnings call transcripts. Evaluation uses network science metrics (centrality, modularity, path length) and accuracy of extracted supply chain relationships against known ground-truth supplier networks.

---

## 2. Evaluation Approaches by Paper

| # | Paper (Short Title) | Evaluation Type | Key Metrics | Data Source |
|---|---------------------|-----------------|-------------|-------------|
| 1 | Agentic LLMs Consensus-Seeking | Simulation case study | Bullwhip ratio, total SC cost, consensus convergence | Multi-echelon SC simulation (beer game) |
| 2 | SCPA / JD.com | Real-world deployment | Planning accuracy (+22%), stock fulfillment (+2--3%), processing time (-40%) | JD.com operational data (millions of SKUs) |
| 3 | Disruption Monitoring | Synthetic scenarios + real case study | F1 score (0.962--0.991), response time (3.83 min), cost ($0.0836/event) | 30 synthesized scenarios; Russia-Ukraine conflict case |
| 4 | AI Agent Systems (AIM-RM) | Empirical benchmarking | Optimal ordering policy achievement, comparative performance | Multi-echelon inventory simulation |
| 5 | InvAgent | Multi-scenario simulation | Inventory cost, stockout rate, total SC cost | Simulated demand patterns (stable, variable, seasonal) |
| 6 | Sustainability Assessment | Comparative scenario analysis | Energy (-70--90%), CO2 (-90--97%), water (-89--98%) | Real-world document extraction tasks |
| 7 | Multi-Agent RAG (Amazon/KDD) | Real-world ticket data | Helpful answer rate (48.74%), unhelpful reduction (77.4%), auto-resolution (50%) | Supply chain support tickets |
| 8 | KG+LLM Warehouse Bottlenecks | Simulated warehouse scenarios | Pass rate (near-perfect for operational Qs), diagnostic accuracy | DES warehouse simulation data |
| 9 | RAG + Prompt Engineering | Experimental evaluation | Decision latency, contextual precision, adaptability | Real-world SC datasets |
| 10 | LLMs in SCM (IFAC MIM) | Pipeline case study | Delivery delay prediction accuracy | Decentralized agent-based SCM system |
| 11 | OptiGuide (Microsoft) | Case study + benchmark | Output accuracy, query interpretation correctness | Microsoft cloud SC server placement |
| 12 | MAS-ML Resilience | Simulation framework | Flexibility, adaptability, cost optimization | Simulated disruption scenarios |
| 13 | Human-Machine Warehouse KG | Simulation + comparison | Human-AI collaborative planning outcomes | Simulation-driven KG data |
| 14 | SC Mapping via RAG | Industry application | Network metrics (centrality, modularity), relationship extraction accuracy | SEC 10-K filings, earnings calls |

---

## 3. Common Evaluation Methodologies in the Field

### 3.1 Simulation-Based Evaluation

The most prevalent approach. Researchers construct supply chain simulations (often based on classical settings such as the beer distribution game or multi-echelon inventory networks) and measure agent performance against traditional baselines.

- **Typical baselines:** (s,S) inventory policies, base-stock policies, centralized demand sharing approaches, heuristic ordering rules.
- **Typical metrics:** Total supply chain cost, bullwhip effect ratio, stockout rate, inventory holding cost, fill rate, order fulfillment time.
- **Strengths:** Controlled environment, reproducibility, ability to test edge cases.
- **Limitations:** May not capture real-world complexity; results depend heavily on simulation fidelity.

### 3.2 Real-World Deployment and Case Studies

A smaller but growing number of studies validate their approaches through real-world deployment, typically in partnership with industry.

- **Examples:** JD.com retail supply chain (Paper 2), Microsoft cloud supply chain (Paper 11), automotive sector disruption (Paper 3).
- **Typical metrics:** Operational efficiency gains (processing time reduction, accuracy improvement), cost per event, response time, stock fulfillment rates.
- **Strengths:** High external validity; demonstrates practical feasibility.
- **Limitations:** Proprietary data limits reproducibility; results may be company-specific.

### 3.3 Synthetic Scenario Evaluation

Some papers construct carefully designed synthetic scenarios to test specific capabilities of their systems.

- **Examples:** 30 synthesized disruption scenarios across manufacturers and disruption classes (Paper 3), warehouse equipment breakdown scenarios (Paper 8).
- **Typical metrics:** F1 score, precision, recall, pass rate, diagnostic accuracy.
- **Strengths:** Targeted evaluation of specific capabilities; can cover edge cases systematically.
- **Limitations:** Scenario design may introduce bias; limited ecological validity.

### 3.4 Comparative Framework Analysis

Several papers compare different operational paradigms (e.g., manual vs. AI-assisted vs. fully agentic) rather than measuring absolute performance.

- **Examples:** Manual vs. HITL vs. agentic AI for document processing (Paper 6), traditional RAG vs. prebuilt KB RAG (Paper 7).
- **Typical metrics:** Relative improvement ratios, resource consumption differences, user satisfaction.
- **Strengths:** Demonstrates value proposition clearly; intuitive for practitioners.
- **Limitations:** Baseline selection affects perceived improvement.

### 3.5 Benchmark Development

An emerging trend is the creation of evaluation benchmarks specifically for LLM-based supply chain systems.

- **Examples:** OptiGuide general evaluation benchmark (Paper 11), structured decision prompt evaluation protocols (Paper 4).
- **Current state:** No widely accepted community benchmark exists yet. Most papers use custom evaluation setups.
- **Gap:** The field lacks standardized benchmarks comparable to those in NLP (e.g., GLUE, SuperGLUE) or computer vision (e.g., ImageNet).

---

## 4. Key Themes and Research Gaps

### 4.1 Emerging Research Themes

1. **Multi-Agent Consensus and Coordination:** How multiple LLM agents representing different supply chain stakeholders can reach consensus on shared decisions (inventory levels, delivery schedules) while balancing competing objectives (Papers 1, 4, 5).

2. **RAG for Supply Chain Knowledge Management:** Using retrieval-augmented generation to ground LLM outputs in real enterprise data, including unstructured communications, regulatory documents, and operational databases (Papers 7, 9, 14).

3. **Disruption Detection and Resilience:** Autonomous monitoring of unstructured data sources (news, social media, regulatory filings) to detect and respond to supply chain disruptions before they cascade (Papers 3, 12).

4. **Warehouse and Logistics Optimization:** Applying LLM-based agents and knowledge graphs to warehouse bottleneck identification, layout planning, and operational efficiency (Papers 8, 13).

5. **Sustainability and ESG Integration:** Assessing the environmental footprint of agentic AI deployments and integrating sustainability metrics into evaluation frameworks (Paper 6).

6. **Human-AI Collaboration:** Designing systems where AI agents augment rather than replace human planners, with appropriate levels of autonomy and oversight (Papers 2, 10, 13).

### 4.2 Identified Research Gaps

| Gap | Description | Opportunity |
|-----|-------------|-------------|
| **Standardized benchmarks** | No community-accepted benchmark for evaluating LLM agents in SC | Develop open-source SC simulation benchmarks |
| **Cross-domain generalizability** | Most studies focus on a single SC function (inventory, disruption, etc.) | End-to-end multi-function agentic systems |
| **Hallucination in critical decisions** | LLM agents may generate factually incorrect SC recommendations | Verification mechanisms, guardrails, human-in-the-loop |
| **Data privacy across tiers** | Multi-tier SC coordination requires data sharing | Federated learning, privacy-preserving agent protocols |
| **Long-horizon planning** | Most evaluations cover short time windows | Longitudinal evaluation over months/years |
| **Scalability evidence** | Limited evidence on performance with large-scale, real SC networks | Industry-scale validation beyond pilot studies |
| **Reproducibility** | Many papers use proprietary data or closed systems | Open-source code, datasets, and evaluation protocols |

---

## 5. Industry Adoption Context

| Company | Initiative | Status |
|---------|-----------|--------|
| **Walmart** | Agentic LLMs for supply chain automation | Active pilots |
| **Siemens** | Agentic AI for manufacturing and logistics | Development |
| **SAP** | SAP Joule -- AI copilot for supply chain | Released |
| **Microsoft** | Copilot for Supply Chain Management; OptiGuide | GA / Research |
| **Google** | Supply chain agent platforms | Development |
| **JD.com** | SCPA framework in production | Deployed |
| **Amazon** | Multi-agent RAG for SC knowledge bases | Research/Deployment |

Taylor & Francis has launched a dedicated special issue on ["The Agentic Supply Chain"](https://think.taylorandfrancis.com/special_issues/agentic-supply-chain/) in the *International Journal of Production Research*, signaling the maturation of this research area.

---

## 6. Bibliography

### Primary Research Papers

1. Jannelli, V., Schoepf, S., Bickel, M., Netland, T., & Brintrup, A. (2024). Agentic LLMs in the Supply Chain: Towards Autonomous Multi-Agent Consensus-Seeking. *arXiv:2411.10184*. Published in *International Journal of Production Research*, 2025. DOI: 10.1080/00207543.2025.2604311. [arXiv](https://arxiv.org/abs/2411.10184)

2. Yin, J., Qi, Y., Zhang, J., Geng, D., Chen, Z., Hu, H., Qi, W., & Shen, Z.-J. M. (2025). Rethinking Supply Chain Planning: A Generative Paradigm. *arXiv:2509.03811*. [arXiv](https://arxiv.org/abs/2509.03811)

3. AlMahri, S., Xu, L., & Brintrup, A. (2026). Automating Supply Chain Disruption Monitoring via an Agentic AI Approach. *arXiv:2601.09680*. [arXiv](https://arxiv.org/abs/2601.09680)

4. Yoshizato, K., Shimizu, K., Higa, R., & Otsuka, T. (2026). AI Agent Systems for Supply Chains: Structured Decision Prompts and Memory Retrieval. *arXiv:2602.05524*. Accepted at AAMAS 2026. [arXiv](https://arxiv.org/abs/2602.05524)

5. Quan, Y. & Liu, Z. (2024). InvAgent: A Large Language Model based Multi-Agent System for Inventory Management in Supply Chains. *arXiv:2407.11384*. [arXiv](https://arxiv.org/abs/2407.11384)

6. Gosmar, D., Pallotta, A. C., & Zenezini, G. (2025). Agentic AI Sustainability Assessment for Supply Chain Document Insights. *arXiv:2511.07097*. [arXiv](https://arxiv.org/abs/2511.07097)

7. Zhang, Y., Shang, Z., Patel, S., & Zuniga, M. (2025). From Unstructured Communication to Intelligent RAG: Multi-Agent Automation for Supply Chain Knowledge Bases. *arXiv:2506.17484*. Presented at 1st Workshop on AI for Supply Chain @ ACM SIGKDD 2025, Toronto. [arXiv](https://arxiv.org/abs/2506.17484)

8. Parekh, R., Gopalakrishnan, S., Ahmad, Z., & Deodhar, A. (2025). Leveraging Knowledge Graphs and LLM Reasoning to Identify Operational Bottlenecks for Warehouse Planning Assistance. *arXiv:2507.17273*. [arXiv](https://arxiv.org/abs/2507.17273)

9. Integrating RAG (Retrieval-Augmented Generation) with Prompt Engineering for Knowledge-Driven Supply Chain Solutions. (2025). *ResearchGate Publication 393479215*. [ResearchGate](https://www.researchgate.net/publication/393479215)

10. Zheng, G., AlMahri, S., Xu, L., & Minaricova, M. (2025). LLMs in Supply Chain Management: Opportunities and a Case Study. *IFAC-PapersOnLine* (11th IFAC MIM 2025, Trondheim, Norway). [ScienceDirect](https://www.sciencedirect.com/science/article/pii/S2405896325012595)

11. Li, B., Mellou, K., Zhang, B., Pathuri, J., & Menache, I. (2023). Large Language Models for Supply Chain Optimization. *arXiv:2307.03875*. [arXiv](https://arxiv.org/abs/2307.03875)

12. Enhancing Supply Chain Resilience with Multi-Agent Systems and Machine Learning: A Framework for Adaptive Decision-Making. (2025). *The American Journal of Engineering and Technology*. [Journal](https://theamericanjournals.com/index.php/tajet/article/view/5919)

13. Intelligent Human-Machine Partnership for Manufacturing: Enhancing Warehouse Planning through Simulation-Driven Knowledge Graphs and LLM Collaboration. (2025). *arXiv:2512.18265*. [arXiv](https://arxiv.org/abs/2512.18265)

14. Supply Chain Mapping through Retrieval-Augmented Generation: Applications to the Electronics Industry. (2026). *Journal of the Operational Research Society* (Taylor & Francis). DOI: 10.1080/01605682.2025.2608868. [Taylor & Francis](https://www.tandfonline.com/doi/full/10.1080/01605682.2025.2608868)
