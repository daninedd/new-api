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
import { ArrowRight, Check, FileText, Rocket } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { PublicLayout } from '@/components/layout'
import { Footer } from '@/components/layout/components/footer'
import { seoLandingPages, type SeoLandingPageKey } from './data'

interface SeoLandingPageProps {
  page: SeoLandingPageKey
}

export function SeoLandingPage(props: SeoLandingPageProps) {
  const content = seoLandingPages[props.page]

  return (
    <PublicLayout showMainContainer={false}>
      <main className='overflow-hidden'>
        <section className='border-border/70 bg-background pt-28 pb-16 md:pt-32 md:pb-20'>
          <div className='mx-auto grid max-w-7xl gap-10 px-4 md:px-6 lg:grid-cols-[minmax(0,1fr)_24rem] lg:items-center'>
            <div className='max-w-3xl space-y-7'>
              <div className='border-border bg-muted/40 text-muted-foreground inline-flex rounded-full border px-3 py-1 text-sm'>
                {content.eyebrow}
              </div>
              <div className='space-y-5'>
                <h1 className='text-4xl leading-tight font-semibold tracking-normal text-balance md:text-6xl'>
                  {content.title}
                </h1>
                <p className='text-muted-foreground max-w-2xl text-lg leading-8 md:text-xl'>
                  {content.description}
                </p>
              </div>
              <div className='flex flex-wrap gap-3'>
                <Button size='lg' render={<Link to='/sign-up' />}>
                  <Rocket />
                  {content.primaryCta}
                </Button>
                <Button
                  variant='outline'
                  size='lg'
                  render={<a href='https://docs.all-llms.com/' />}
                >
                  <FileText />
                  {content.secondaryCta}
                </Button>
              </div>
            </div>

            <div className='border-border bg-card/80 rounded-lg border p-5 shadow-sm'>
              <div className='space-y-4'>
                <div className='bg-muted/50 rounded-lg p-4'>
                  <div className='text-muted-foreground text-sm'>
                    Gateway flow
                  </div>
                  <div className='mt-3 grid gap-2 text-sm'>
                    {[
                      'Application request',
                      'All-LLMs policy layer',
                      'Provider channel',
                      'Usage and billing log',
                    ].map((item, index) => (
                      <div
                        key={item}
                        className='border-border bg-background flex items-center gap-3 rounded-md border px-3 py-2'
                      >
                        <span className='bg-primary/10 text-primary flex size-6 items-center justify-center rounded-full text-xs font-medium'>
                          {index + 1}
                        </span>
                        <span>{item}</span>
                      </div>
                    ))}
                  </div>
                </div>
                <div className='grid grid-cols-2 gap-3 text-sm'>
                  {['Keys', 'Routing', 'Logs', 'Billing'].map((item) => (
                    <div
                      key={item}
                      className='border-border bg-background rounded-md border px-3 py-2'
                    >
                      <Check className='text-primary mb-2 size-4' />
                      {item}
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </section>

        <section className='border-border/70 border-t py-16 md:py-20'>
          <div className='mx-auto max-w-7xl px-4 md:px-6'>
            <div className='max-w-3xl space-y-4'>
              <h2 className='text-3xl font-semibold tracking-normal md:text-4xl'>
                {content.introTitle}
              </h2>
              <p className='text-muted-foreground text-lg leading-8'>
                {content.intro}
              </p>
            </div>
            <div className='mt-10 grid gap-4 md:grid-cols-3'>
              {content.highlights.map((item) => {
                const Icon = item.icon
                return (
                  <article
                    key={item.title}
                    className='border-border bg-card rounded-lg border p-5'
                  >
                    <Icon className='text-primary size-6' />
                    <h3 className='mt-5 text-lg font-semibold'>
                      {item.title}
                    </h3>
                    <p className='text-muted-foreground mt-2 leading-7'>
                      {item.description}
                    </p>
                  </article>
                )
              })}
            </div>
          </div>
        </section>

        <section className='bg-muted/30 border-border/70 border-t py-16 md:py-20'>
          <div className='mx-auto grid max-w-7xl gap-10 px-4 md:px-6 lg:grid-cols-[0.85fr_1fr]'>
            <div className='space-y-4'>
              <h2 className='text-3xl font-semibold tracking-normal md:text-4xl'>
                {content.workflowTitle}
              </h2>
              <p className='text-muted-foreground leading-7'>
                {content.useCasesTitle}
              </p>
              <ul className='space-y-3'>
                {content.useCases.map((item) => (
                  <li key={item} className='flex gap-3'>
                    <Check className='text-primary mt-1 size-4 shrink-0' />
                    <span className='text-muted-foreground'>{item}</span>
                  </li>
                ))}
              </ul>
            </div>
            <div className='grid gap-4'>
              {content.workflow.map((item, index) => (
                <article
                  key={item.title}
                  className='border-border bg-background rounded-lg border p-5'
                >
                  <div className='flex gap-4'>
                    <span className='bg-primary text-primary-foreground flex size-8 shrink-0 items-center justify-center rounded-full text-sm font-medium'>
                      {index + 1}
                    </span>
                    <div>
                      <h3 className='font-semibold'>{item.title}</h3>
                      <p className='text-muted-foreground mt-2 leading-7'>
                        {item.description}
                      </p>
                    </div>
                  </div>
                </article>
              ))}
            </div>
          </div>
        </section>

        <section className='border-border/70 border-t py-16 md:py-20'>
          <div className='mx-auto max-w-5xl px-4 md:px-6'>
            <h2 className='text-3xl font-semibold tracking-normal md:text-4xl'>
              Questions teams ask before adopting an AI gateway
            </h2>
            <div className='mt-8 grid gap-4 md:grid-cols-2'>
              {content.faq.map((item) => (
                <article
                  key={item.question}
                  className='border-border rounded-lg border p-5'
                >
                  <h3 className='font-semibold'>{item.question}</h3>
                  <p className='text-muted-foreground mt-2 leading-7'>
                    {item.answer}
                  </p>
                </article>
              ))}
            </div>
            <div className='mt-10 flex flex-wrap gap-3'>
              <Button size='lg' render={<Link to='/sign-up' />}>
                Get started with All-LLMs
                <ArrowRight />
              </Button>
              <Button variant='outline' size='lg' render={<Link to='/pricing' />}>
                Explore model marketplace
              </Button>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </PublicLayout>
  )
}
