"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";
import { AppRoute, ButtonSize, ButtonVariant, IconName, WelcomeStepKey } from "@/lib/enums";
import { markOnboardingComplete } from "@/lib/storage";
import { welcomePageContent, welcomeSteps } from "@/content/welcome";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { StepProgress } from "@/components/welcome/StepProgress";
import { TourStep } from "@/components/welcome/TourStep";

export function OnboardingWizard() {
  const router = useRouter();
  const [index, setIndex] = useState(0);

  const step = welcomeSteps[index];
  const isFirstStep = index === 0;
  const isLastStep = index === welcomeSteps.length - 1;

  const complete = (route: AppRoute) => {
    markOnboardingComplete();
    router.push(route);
  };

  if (step === undefined) {
    return null;
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-wrap items-center justify-between gap-4">
        <StepProgress current={index + 1} total={welcomeSteps.length} />

        <Button
          variant={ButtonVariant.Ghost}
          size={ButtonSize.Sm}
          onClick={() => complete(AppRoute.Markets)}
          ariaLabel={welcomePageContent.skipLabel}
        >
          {welcomePageContent.skipLabel}
        </Button>
      </div>

      <Card className="p-6 sm:p-8">
        <TourStep key={step.key} step={step} />
      </Card>

      <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <Button
          variant={ButtonVariant.Secondary}
          disabled={isFirstStep}
          onClick={() => setIndex((current) => Math.max(0, current - 1))}
        >
          {welcomePageContent.backLabel}
        </Button>

        {isLastStep ? (
          <div className="flex flex-col gap-3 sm:flex-row">
            <Button variant={ButtonVariant.Secondary} onClick={() => complete(AppRoute.Markets)}>
              {welcomePageContent.finishLabel}
            </Button>
            <Button trailingIcon={IconName.ArrowRight} onClick={() => complete(AppRoute.Practice)}>
              {welcomePageContent.practiceLabel}
            </Button>
          </div>
        ) : (
          <Button
            trailingIcon={IconName.ArrowRight}
            onClick={() => setIndex((current) => Math.min(welcomeSteps.length - 1, current + 1))}
          >
            {step.key === WelcomeStepKey.Wallet ? "I understand" : welcomePageContent.nextLabel}
          </Button>
        )}
      </div>
    </div>
  );
}
