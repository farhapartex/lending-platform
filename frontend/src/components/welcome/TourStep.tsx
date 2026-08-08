import { WelcomeStepKey } from "@/lib/enums";
import type { WelcomeStep } from "@/content/welcome";
import { ConceptCard } from "@/components/welcome/ConceptCard";
import { PracticeModeCard } from "@/components/welcome/PracticeModeCard";
import { WalletBasicsCard } from "@/components/welcome/WalletBasicsCard";

type TourStepProps = {
  step: WelcomeStep;
};

export function TourStep({ step }: TourStepProps) {
  if (step.key === WelcomeStepKey.Wallet) {
    return <WalletBasicsCard step={step} />;
  }

  if (step.key === WelcomeStepKey.Ready) {
    return (
      <div className="flex flex-col gap-5">
        <ConceptCard step={step} />
        <PracticeModeCard />
      </div>
    );
  }

  return <ConceptCard step={step} />;
}
