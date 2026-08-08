"use client";

import { AppRoute, BadgeTone, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { errorContent } from "@/content/system";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";

type ErrorViewProps = {
  digest?: string;
  onRetry: () => void;
};

export function ErrorView({ digest, onRetry }: ErrorViewProps) {
  return (
    <Container className="max-w-2xl py-20">
      <div className="flex flex-col items-start gap-6">
        <div className="flex flex-col gap-3">
          <h1 className="text-balance text-3xl font-semibold tracking-tight text-ink sm:text-4xl">
            {errorContent.title}
          </h1>
          <p className="text-pretty text-base leading-relaxed text-ink-soft">{errorContent.description}</p>
        </div>

        <Alert title="Your funds are safe" tone={BadgeTone.Positive} icon={IconName.ShieldCheck}>
          {errorContent.fundsNote}
        </Alert>

        <div className="flex flex-col gap-3 sm:flex-row">
          <Button size={ButtonSize.Lg} onClick={onRetry}>
            {errorContent.retryCta}
          </Button>
          <Button href={AppRoute.Markets} size={ButtonSize.Lg} variant={ButtonVariant.Secondary}>
            {errorContent.homeCta}
          </Button>
        </div>

        {digest === undefined ? null : (
          <p className="text-xs text-ink-faint">
            {errorContent.digestLabel}: <span className="font-mono">{digest}</span>
          </p>
        )}
      </div>
    </Container>
  );
}
