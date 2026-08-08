import { AppRoute, IconName } from "@/lib/enums";
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

      <div className="flex flex-col gap-2 rounded-card border border-line bg-surface-muted p-5">
        <span className="text-sm font-medium text-ink">Coming from a bank or a broker?</span>
        <p className="text-sm leading-relaxed text-ink-soft">
          The biggest difference is that there is no institution holding your money and no support desk who can undo a
          mistake. Everything you can do here, you do yourself.
        </p>
        <TextLink href={AppRoute.LearnFaq} trailingIcon={IconName.ArrowRight}>
          Read the honest answers about custody
        </TextLink>
      </div>
    </div>
  );
}
