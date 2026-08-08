import { IconName } from "@/lib/enums";
import type { WelcomeStep } from "@/content/welcome";
import { welcomePageContent } from "@/content/welcome";
import { Alert } from "@/components/ui/Alert";
import { Icon } from "@/components/ui/Icon";

type ConceptCardProps = {
  step: WelcomeStep;
};

export function ConceptCard({ step }: ConceptCardProps) {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col gap-4">
        <span className="grid size-14 place-items-center rounded-panel bg-gradient-to-br from-brand-soft to-brand-muted text-brand">
          <Icon name={step.icon} className="size-7" />
        </span>

        <div className="flex flex-col gap-2">
          <span className="text-xs font-semibold uppercase tracking-[0.14em] text-brand">{step.eyebrow}</span>
          <h2 className="text-balance text-2xl font-semibold tracking-tight text-ink sm:text-3xl">{step.title}</h2>
        </div>
      </div>

      <p className="flex items-start gap-2.5 rounded-card border border-mint-border bg-mint-soft px-4 py-3 text-sm font-medium text-mint-ink">
        <Icon name={IconName.Check} className="mt-0.5 size-4 shrink-0" />
        <span>
          <span className="sr-only">{welcomePageContent.takeawayLabel}: </span>
          {step.takeaway}
        </span>
      </p>

      <div className="flex flex-col gap-4">
        {step.paragraphs.map((paragraph) => (
          <p key={paragraph} className="text-base leading-relaxed text-ink-soft">
            {paragraph}
          </p>
        ))}
      </div>

      {step.bullets === undefined ? null : (
        <ul className="flex flex-col gap-2.5 rounded-card bg-surface-muted p-5">
          {step.bullets.map((item) => (
            <li key={item} className="flex gap-3 text-sm leading-relaxed text-ink-soft">
              <Icon name={IconName.Check} className="mt-0.5 size-4 shrink-0 text-brand" />
              <span>{item}</span>
            </li>
          ))}
        </ul>
      )}

      {step.callout === undefined ? null : (
        <Alert title={step.callout.title} tone={step.callout.tone} icon={step.callout.icon}>
          {step.callout.body}
        </Alert>
      )}
    </div>
  );
}
