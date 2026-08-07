import { AppRoute, ButtonVariant, IconName, SectionId, SectionTone } from "@/lib/enums";
import { feeItems } from "@/content/protocol";
import { Button } from "@/components/ui/Button";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function FeeTransparencyTeaser() {
  return (
    <Section id={SectionId.Fees} tone={SectionTone.Muted} bordered>
      <div className="grid gap-10 lg:grid-cols-[minmax(0,22rem)_minmax(0,1fr)] lg:gap-16">
        <div className="flex flex-col gap-6">
          <SectionHeading
            sectionId={SectionId.Fees}
            eyebrow="Fees"
            title="Every fee, before you deposit."
            description="There is no fee schedule hidden behind a signup. This is the complete list."
          />
          <Button href={AppRoute.LearnFees} variant={ButtonVariant.Secondary} trailingIcon={IconName.ArrowRight}>
            Full fee disclosure
          </Button>
        </div>

        <dl className="divide-y divide-line overflow-hidden rounded-card border border-line bg-surface">
          {feeItems.map((fee) => (
            <div key={fee.kind} className="flex flex-col gap-1.5 p-5 sm:p-6">
              <div className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-1">
                <dt className="text-base font-medium text-ink">{fee.label}</dt>
                <dd className="text-base font-semibold text-brand-ink tabular-nums">{fee.value}</dd>
              </div>
              <p className="max-w-xl text-sm leading-relaxed text-ink-soft">{fee.description}</p>
            </div>
          ))}
        </dl>
      </div>
    </Section>
  );
}
