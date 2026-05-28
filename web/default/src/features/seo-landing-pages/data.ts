/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import {
  Boxes,
  BrainCircuit,
  CloudCog,
  GitBranch,
  KeyRound,
  LineChart,
  LockKeyhole,
  Network,
  Route,
  ServerCog,
  ShieldCheck,
  Shuffle,
  WalletCards,
  Zap,
} from 'lucide-react'
import type { LucideIcon } from 'lucide-react'

export type SeoLandingPageKey =
  | 'openai-compatible-api-gateway'
  | 'claude-api-gateway'
  | 'gemini-api-gateway'
  | 'ai-model-router'
  | 'self-hosted-ai-gateway'

export interface SeoLandingPageContent {
  slug: SeoLandingPageKey
  eyebrow: string
  title: string
  description: string
  primaryCta: string
  secondaryCta: string
  introTitle: string
  intro: string
  highlights: Array<{
    icon: LucideIcon
    title: string
    description: string
  }>
  workflowTitle: string
  workflow: Array<{
    title: string
    description: string
  }>
  useCasesTitle: string
  useCases: string[]
  faq: Array<{
    question: string
    answer: string
  }>
}

export const seoLandingPages: Record<
  SeoLandingPageKey,
  SeoLandingPageContent
> = {
  'openai-compatible-api-gateway': {
    slug: 'openai-compatible-api-gateway',
    eyebrow: 'OpenAI-compatible API gateway',
    title: 'OpenAI-compatible API gateway for every AI provider',
    description:
      'Route OpenAI-compatible requests to multiple upstream model providers while keeping one API contract, one key layer, and one place to monitor usage.',
    primaryCta: 'Start routing requests',
    secondaryCta: 'Read the docs',
    introTitle: 'Keep the OpenAI API shape, expand the provider backend',
    intro:
      'All-LLMs lets your applications keep familiar chat completion and compatible API workflows while you connect OpenAI, Claude, Gemini, DeepSeek, Qwen, Llama, and other model providers behind a single gateway.',
    highlights: [
      {
        icon: Route,
        title: 'One compatible endpoint',
        description:
          'Send requests through a unified API surface instead of wiring every application to every provider separately.',
      },
      {
        icon: KeyRound,
        title: 'Centralized API keys',
        description:
          'Issue, rotate, and revoke application keys without exposing upstream provider credentials to every client.',
      },
      {
        icon: LineChart,
        title: 'Usage and cost visibility',
        description:
          'Track token usage, latency, and spend across providers from one operational dashboard.',
      },
    ],
    workflowTitle: 'A cleaner path from app code to model providers',
    workflow: [
      {
        title: 'Point your app at All-LLMs',
        description:
          'Use an OpenAI-compatible base URL and authentication pattern that existing SDKs and tools already understand.',
      },
      {
        title: 'Connect upstream providers',
        description:
          'Add provider channels, model mappings, groups, and routing rules without changing application code.',
      },
      {
        title: 'Operate with shared controls',
        description:
          'Apply rate limits, quotas, billing rules, and logs across all traffic that passes through the gateway.',
      },
    ],
    useCasesTitle: 'Built for teams standardizing AI access',
    useCases: [
      'SaaS products that need one AI API layer across multiple providers.',
      'Internal developer platforms consolidating provider credentials.',
      'AI applications that want fallback capacity without changing SDKs.',
    ],
    faq: [
      {
        question: 'Do I need to rewrite my OpenAI-compatible client?',
        answer:
          'In most cases, no. You point the client to the All-LLMs base URL, use a gateway key, and configure the upstream channels in the dashboard.',
      },
      {
        question: 'Can one gateway route to non-OpenAI providers?',
        answer:
          'Yes. The gateway is designed to connect compatible and cross-provider routes so applications can use a consistent integration layer.',
      },
    ],
  },
  'claude-api-gateway': {
    slug: 'claude-api-gateway',
    eyebrow: 'Claude API gateway',
    title: 'Claude API gateway with unified access control',
    description:
      'Connect Claude workflows through a managed gateway for keys, quotas, routing, billing, and observability across your AI applications.',
    primaryCta: 'Connect Claude workflows',
    secondaryCta: 'Explore model pricing',
    introTitle: 'Make Claude access easier to manage at scale',
    intro:
      'All-LLMs helps teams centralize Claude access alongside other model providers, so product, automation, and agent workloads can share consistent authentication, monitoring, and usage controls.',
    highlights: [
      {
        icon: BrainCircuit,
        title: 'Claude-ready routing',
        description:
          'Route assistant and agent workloads through a gateway layer designed for multi-provider AI operations.',
      },
      {
        icon: ShieldCheck,
        title: 'Governed access',
        description:
          'Control which users, teams, or applications can access Claude-backed routes and how much they can consume.',
      },
      {
        icon: WalletCards,
        title: 'Transparent billing',
        description:
          'Turn provider usage into clearer cost tracking and billing records for your own users or teams.',
      },
    ],
    workflowTitle: 'Operate Claude alongside the rest of your model stack',
    workflow: [
      {
        title: 'Add Claude as a channel',
        description:
          'Configure upstream credentials and model availability in the gateway instead of scattering secrets across services.',
      },
      {
        title: 'Expose controlled routes',
        description:
          'Publish gateway keys and groups with the access, quotas, and limits each product or team needs.',
      },
      {
        title: 'Monitor every request',
        description:
          'Review logs, latency, token usage, and spend from the same dashboard used for other providers.',
      },
    ],
    useCasesTitle: 'Common Claude gateway use cases',
    useCases: [
      'Agent platforms that need controlled Claude access for many users.',
      'Internal tools that must keep upstream API keys out of client apps.',
      'Products that combine Claude with OpenAI, Gemini, or open models.',
    ],
    faq: [
      {
        question: 'Can Claude live beside other providers in the same gateway?',
        answer:
          'Yes. All-LLMs is built for multi-provider routing, so Claude access can be managed with other AI channels.',
      },
      {
        question: 'Can I track Claude usage per user or application?',
        answer:
          'Yes. Gateway logs and billing controls help separate traffic by keys, users, groups, and configured routing rules.',
      },
    ],
  },
  'gemini-api-gateway': {
    slug: 'gemini-api-gateway',
    eyebrow: 'Gemini API gateway',
    title: 'Gemini API gateway for multi-provider AI applications',
    description:
      'Use All-LLMs to connect Gemini-backed routes with shared API keys, monitoring, limits, and fallback options across your AI stack.',
    primaryCta: 'Route Gemini traffic',
    secondaryCta: 'View docs',
    introTitle: 'Bring Gemini into a unified AI gateway',
    intro:
      'Instead of managing Gemini access separately from the rest of your model stack, All-LLMs gives you a single control plane for model access, credentials, usage records, and operational policies.',
    highlights: [
      {
        icon: CloudCog,
        title: 'Provider abstraction',
        description:
          'Keep provider details in the gateway so applications can focus on product logic rather than integration sprawl.',
      },
      {
        icon: Shuffle,
        title: 'Flexible routing',
        description:
          'Route by model, group, or operational needs and combine Gemini with other providers behind one access layer.',
      },
      {
        icon: Zap,
        title: 'Fast operational changes',
        description:
          'Adjust channels, availability, and limits centrally without redeploying every application that uses AI.',
      },
    ],
    workflowTitle: 'Connect Gemini without creating another silo',
    workflow: [
      {
        title: 'Configure provider credentials',
        description:
          'Store upstream access in the gateway and keep application keys separate from provider keys.',
      },
      {
        title: 'Publish gateway routes',
        description:
          'Give applications a stable route while the gateway handles provider-specific mapping and policies.',
      },
      {
        title: 'Review usage in one place',
        description:
          'Use shared logs and usage analytics to compare Gemini traffic with the rest of your AI providers.',
      },
    ],
    useCasesTitle: 'Where Gemini gateway routing helps',
    useCases: [
      'Applications that test Gemini against other providers.',
      'Teams building AI features with centralized model governance.',
      'Platforms that need provider fallback and cost monitoring.',
    ],
    faq: [
      {
        question: 'Why use a gateway instead of calling Gemini directly?',
        answer:
          'A gateway gives you shared keys, usage controls, routing rules, and observability across providers instead of one-off integrations.',
      },
      {
        question: 'Can I combine Gemini with OpenAI-compatible routes?',
        answer:
          'Yes. The platform is designed for mixed provider environments where applications need a consistent operational layer.',
      },
    ],
  },
  'ai-model-router': {
    slug: 'ai-model-router',
    eyebrow: 'AI model router',
    title: 'AI model router for reliability, cost control, and scale',
    description:
      'Route AI requests across providers and models with centralized policies for fallback, quotas, cost tracking, and observability.',
    primaryCta: 'Build your routing layer',
    secondaryCta: 'Compare models',
    introTitle: 'One routing layer for a changing model landscape',
    intro:
      'AI model availability, latency, and pricing change constantly. All-LLMs gives teams a routing layer that can evolve without forcing every product integration to change with it.',
    highlights: [
      {
        icon: Network,
        title: 'Multi-provider routing',
        description:
          'Send traffic to the right upstream channel based on provider availability, model fit, or team policy.',
      },
      {
        icon: Boxes,
        title: 'Model catalog control',
        description:
          'Manage which models are exposed, grouped, priced, and available to each application or user segment.',
      },
      {
        icon: LineChart,
        title: 'Operational analytics',
        description:
          'Use logs, cost data, and latency signals to understand how model routing performs in production.',
      },
    ],
    workflowTitle: 'How routing becomes easier',
    workflow: [
      {
        title: 'Create provider channels',
        description:
          'Connect upstream providers and define the models or endpoints each channel should expose.',
      },
      {
        title: 'Apply routing rules',
        description:
          'Use groups, model mappings, quotas, and access controls to determine how requests move through the gateway.',
      },
      {
        title: 'Tune with real usage',
        description:
          'Review logs and cost patterns, then adjust routing without changing every consuming application.',
      },
    ],
    useCasesTitle: 'Routing scenarios All-LLMs supports',
    useCases: [
      'Fallback routing when a provider is slow or unavailable.',
      'Cost-aware routing across premium and efficient models.',
      'Team-specific model access for internal AI platforms.',
    ],
    faq: [
      {
        question: 'Is an AI model router only for large teams?',
        answer:
          'No. Even small teams benefit when provider keys, model availability, billing, and logs are managed in one layer.',
      },
      {
        question: 'Can routing decisions change over time?',
        answer:
          'Yes. The gateway lets you adjust provider channels, groups, and model exposure as usage patterns change.',
      },
    ],
  },
  'self-hosted-ai-gateway': {
    slug: 'self-hosted-ai-gateway',
    eyebrow: 'Self-hosted AI gateway',
    title: 'Self-hosted AI gateway for teams that need control',
    description:
      'Deploy a unified AI API gateway in your own environment to manage providers, users, keys, billing, and usage logs with more operational control.',
    primaryCta: 'Deploy your gateway',
    secondaryCta: 'Review documentation',
    introTitle: 'Own the control plane for your AI traffic',
    intro:
      'All-LLMs can be self-hosted so teams can keep gateway operations closer to their infrastructure, integrate with internal policies, and standardize AI access without handing every workflow to a third-party proxy.',
    highlights: [
      {
        icon: ServerCog,
        title: 'Deploy in your environment',
        description:
          'Run the gateway where your team manages infrastructure and connect it to the providers your applications need.',
      },
      {
        icon: LockKeyhole,
        title: 'Keep access centralized',
        description:
          'Reduce scattered provider keys by issuing gateway keys and enforcing usage controls from one place.',
      },
      {
        icon: GitBranch,
        title: 'Open-source foundation',
        description:
          'Build on a community-driven gateway that can be extended, audited, and adapted to your deployment model.',
      },
    ],
    workflowTitle: 'A practical self-hosted setup',
    workflow: [
      {
        title: 'Install the gateway',
        description:
          'Deploy All-LLMs with your preferred database and cache setup, then configure system settings.',
      },
      {
        title: 'Connect providers and users',
        description:
          'Add upstream providers, create access groups, and issue application keys for your teams or customers.',
      },
      {
        title: 'Monitor and refine',
        description:
          'Use logs, quotas, billing rules, and provider controls to keep AI traffic observable and manageable.',
      },
    ],
    useCasesTitle: 'Best-fit self-hosted use cases',
    useCases: [
      'Companies that need tighter control over AI provider credentials.',
      'Developer platforms that want an internal AI access layer.',
      'Products that need custom billing, quotas, or routing policies.',
    ],
    faq: [
      {
        question: 'Why self-host an AI gateway?',
        answer:
          'Self-hosting gives teams more control over deployment, access policies, observability, and integration with internal operations.',
      },
      {
        question: 'Can a self-hosted gateway still use cloud AI providers?',
        answer:
          'Yes. The gateway can run in your environment while routing requests to the upstream providers you configure.',
      },
    ],
  },
}
