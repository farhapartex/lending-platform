import { IconName } from "@/lib/enums";
import type { WelcomeStep } from "@/content/welcome";
import { Alert } from "@/components/ui/Alert";
import { Icon } from "@/components/ui/Icon";

type ConceptCardProps = {
  step: WelcomeStep;
};

export function ConceptCard({ step }: ConceptCardProps) {
  return (
    <div className="flex flex-col gap-5">
      <span className="grid size-12 place-items-center rounded-tile bg-brand-soft text-brand">
        <Icon name={step.icon} className="size-6" />
      </span>

      <div className="flex flex-col gap-2">
        <span className="text-xs font-semibold uppercase tracking-[0.14em] text-brand">{step.eyebrow}</span>
        <h2 className="text-balance text-2xl font-semibold tracking-tight text-ink sm:text-3xl">{step.title}</h2>
      </div>

      <div className="flex flex-col gap-4">
        {step.paragraphs.map((paragraph) => (
          <p key={paragraph} className="text-base leading-relaxed text-ink-soft">
            {paragraph}
          </p>
        ))}
      </div>

      {step.bullets === undefined ? null : (
        <ul className="flex flex-col gap-3">
          {step.bullets.map((item) => (
            <li key={item} className="flex gap-3 text-base leading-relaxed text-ink-soft">
              <Icon name={IconName.Check} className="mt-1 size-4 shrink-0 text-mint" />
              <span>{item}</span>
            </li>
          ))}
        </ul>
      )}

      {step.callout === undefined ? null : (
        <Alert title={step.callout.title} tone={step.callout.tone} icon={IconName.Info}>
          {step.callout.body}
        </Alert>
      )}
    </div>
  );
}
