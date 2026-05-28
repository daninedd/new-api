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
import { Link } from '@tanstack/react-router'
import {
  ArrowRight,
  BrainCircuit,
  CloudCog,
  Network,
  Route,
  ServerCog,
} from 'lucide-react'
import { AnimateInView } from '@/components/animate-in-view'

const solutions = [
  {
    title: 'OpenAI-compatible API Gateway',
    description:
      'Keep familiar OpenAI-compatible workflows while routing to multiple providers through one gateway.',
    to: '/openai-compatible-api-gateway',
    icon: Route,
  },
  {
    title: 'Claude API Gateway',
    description:
      'Centralize Claude access with shared keys, quotas, billing, and request logs.',
    to: '/claude-api-gateway',
    icon: BrainCircuit,
  },
  {
    title: 'Gemini API Gateway',
    description:
      'Connect Gemini-backed routes without creating another isolated provider integration.',
    to: '/gemini-api-gateway',
    icon: CloudCog,
  },
  {
    title: 'AI Model Router',
    description:
      'Route traffic across providers for reliability, cost control, and operational visibility.',
    to: '/ai-model-router',
    icon: Network,
  },
  {
    title: 'Self-hosted AI Gateway',
    description:
      'Deploy your own control plane for users, keys, providers, billing, and logs.',
    to: '/self-hosted-ai-gateway',
    icon: ServerCog,
  },
] as const

export function Solutions() {
  return (
    <section className='border-border/40 relative z-10 border-t px-6 py-20 md:py-24'>
      <div className='mx-auto max-w-6xl'>
        <AnimateInView className='mb-10 max-w-2xl'>
          <p className='text-muted-foreground mb-3 text-xs font-medium tracking-widest uppercase'>
            Solutions
          </p>
          <h2 className='text-2xl leading-tight font-bold tracking-tight md:text-3xl'>
            Explore focused AI gateway use cases
          </h2>
          <p className='text-muted-foreground mt-4 max-w-xl text-sm leading-relaxed md:text-base'>
            Practical guides for routing, governing, and self-hosting AI model
            access with All-LLMs.
          </p>
        </AnimateInView>

        <div className='grid gap-4 md:grid-cols-2 lg:grid-cols-3'>
          {solutions.map((solution, index) => {
            const Icon = solution.icon

            return (
              <AnimateInView
                key={solution.to}
                delay={index * 80}
                animation='fade-up'
              >
                <Link
                  to={solution.to}
                  className='border-border/50 bg-background group hover:border-border hover:bg-muted/20 flex h-full min-h-[168px] flex-col rounded-lg border p-5 transition-colors'
                >
                  <div className='mb-4 flex items-center justify-between gap-4'>
                    <div className='border-border/50 bg-muted/30 text-muted-foreground group-hover:text-foreground flex size-10 shrink-0 items-center justify-center rounded-lg border transition-colors'>
                      <Icon className='size-5' strokeWidth={1.5} />
                    </div>
                    <ArrowRight className='text-muted-foreground size-4 shrink-0 transition-transform duration-200 group-hover:translate-x-0.5' />
                  </div>
                  <h3 className='text-base leading-snug font-semibold'>
                    {solution.title}
                  </h3>
                  <p className='text-muted-foreground mt-2 text-sm leading-relaxed'>
                    {solution.description}
                  </p>
                </Link>
              </AnimateInView>
            )
          })}
        </div>
      </div>
    </section>
  )
}
