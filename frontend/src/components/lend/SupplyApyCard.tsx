import { AppRoute, IconName, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { supplyApyRate } from "@/content/protocol";
import { lendPageContent } from "@/content/lend";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { TextLink } from "@/components/ui/TextLink";

export function SupplyApyCard() {
  return (
    <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-3 p-6">
      <span className="text-sm text-ink-soft">Current supply APY</span>
      <span className="text-3xl font-semibold tracking-tight text-ink tabular-nums">
        {formatValue(supplyApyRate, ValueFormat.Percent)}
      </span>
      <p className="flex items-start gap-2 text-sm leading-relaxed text-ink-soft">
        <Icon name={IconName.Info} className="mt-0.5 size-4 text-brand" />
        {lendPageContent.compoundingNote}
      </p>
      <TextLink href={AppRoute.Markets} trailingIcon={IconName.ArrowRight}>
        See how rates are set
      </TextLink>
    </Card>
  );
}
