import { AppRoute, IconName } from "@/lib/enums";
import { Icon } from "@/components/ui/Icon";
import { TextLink } from "@/components/ui/TextLink";
import { ConceptCard } from "@/components/welcome/ConceptCard";
import type { WelcomeStep } from "@/content/welcome";

type WalletBasicsCardProps = {
  step: WelcomeStep;
};

export function WalletBasicsCard({ step }: WalletBasicsCardProps) {
  return (
    <div className="flex flex-col gap-5">
      <ConceptCard step={step} />

      <div className="flex gap-3 rounded-card border border-line bg-surface-muted p-5">
        <Icon name={IconName.Info} className="mt-0.5 size-4 shrink-0 text-brand" />
        <div className="flex flex-col gap-2">
          <span className="text-sm font-medium text-ink">Used to banks or brokers?</span>
          <p className="text-sm leading-relaxed text-ink-soft">
            The main difference is that you are in charge rather than an institution. That means no one can lock you out,
            and it also means you keep your own backup.
          </p>
          <TextLink href={AppRoute.LearnFaq} trailingIcon={IconName.ArrowRight}>
            Common questions answered
          </TextLink>
        </div>
      </div>
    </div>
  );
}
