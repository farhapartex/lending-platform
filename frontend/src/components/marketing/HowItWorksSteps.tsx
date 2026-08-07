import { AppRoute, ButtonVariant, IconName, SectionId, SectionTone, TextAlign } from "@/lib/enums";
import { howItWorksSteps } from "@/content/landing";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export function HowItWorksSteps() {
  return (
    <Section id={SectionId.HowItWorks} tone={SectionTone.Surface} bordered>
      <SectionHeading
        sectionId={SectionId.HowItWorks}
        eyebrow="How it works"
        title="Three steps from wallet to position."
        description="No application, no credit check, and nothing to wait for. The whole flow takes a couple of minutes."
        align={TextAlign.Center}
      />

      <ol className="mt-12 grid gap-6 md:grid-cols-3">
        {howItWorksSteps.map((step, index) => (
          <li key={step.key} className="relative flex flex-col gap-4 rounded-card border border-line bg-canvas p-6">
            <div className="flex items-center gap-3">
              <span className="grid size-9 place-items-center rounded-pill bg-brand text-sm font-semibold text-white tabular-nums">
                {index + 1}
              </span>
              <span className="grid size-9 place-items-center rounded-tile bg-brand-soft text-brand">
                <Icon name={step.icon} className="size-4.5" />
              </span>
            </div>
            <h3 className="text-lg font-semibold tracking-tight text-ink">{step.title}</h3>
            <p className="text-sm leading-relaxed text-ink-soft">{step.description}</p>
          </li>
        ))}
      </ol>

      <div className="mt-10 flex justify-center">
        <Button href={AppRoute.LearnHowItWorks} variant={ButtonVariant.Subtle} trailingIcon={IconName.ArrowRight}>
          Read the full walkthrough
        </Button>
      </div>
    </Section>
  );
}
