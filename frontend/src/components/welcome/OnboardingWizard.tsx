"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";
import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { markOnboardingComplete } from "@/lib/storage";
import { welcomePageContent, welcomeSteps } from "@/content/welcome";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
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
      <div className="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
        <StepProgress current={index + 1} onStepSelect={setIndex} />

        <Button
          variant={ButtonVariant.Ghost}
          size={ButtonSize.Sm}
          className="self-start"
          onClick={() => complete(AppRoute.Markets)}
        >
          {welcomePageContent.skipLabel}
        </Button>
      </div>

      <Card className="p-6 sm:p-8">
        <TourStep key={step.key} step={step} />
      </Card>

      <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        {isFirstStep ? (
          <span className="hidden sm:block" />
        ) : (
          <Button
            variant={ButtonVariant.Secondary}
            onClick={() => setIndex((current) => Math.max(0, current - 1))}
          >
            {welcomePageContent.backLabel}
          </Button>
        )}

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
            size={ButtonSize.Lg}
            trailingIcon={IconName.ArrowRight}
            onClick={() => setIndex((current) => Math.min(welcomeSteps.length - 1, current + 1))}
          >
            {welcomePageContent.nextLabel}
          </Button>
        )}
      </div>

      <p className="flex items-center justify-center gap-2 text-xs text-ink-faint">
        <Icon name={IconName.Lock} className="size-3.5 text-mint" />
        {welcomePageContent.reassurance}
      </p>
    </div>
  );
}
