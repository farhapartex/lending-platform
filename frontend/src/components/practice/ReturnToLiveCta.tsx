import { AppRoute, ButtonVariant, IconName } from "@/lib/enums";
import { practicePageContent } from "@/content/practice";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";

export function ReturnToLiveCta() {
  return (
    <Container className="pb-16">
      <div className="flex flex-col gap-4 rounded-panel border border-line bg-surface p-6 sm:flex-row sm:items-center sm:justify-between sm:p-8">
        <div className="flex max-w-2xl flex-col gap-1.5">
          <h2 className="text-lg font-semibold tracking-tight text-ink">{practicePageContent.returnTitle}</h2>
          <p className="text-sm leading-relaxed text-ink-soft">{practicePageContent.returnDescription}</p>
        </div>

        <Button
          href={AppRoute.Markets}
          variant={ButtonVariant.Secondary}
          trailingIcon={IconName.ArrowRight}
          className="shrink-0"
        >
          {practicePageContent.returnCta}
        </Button>
      </div>
    </Container>
  );
}
