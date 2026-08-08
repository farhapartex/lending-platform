import { AppRoute, IconName, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { basisPoints } from "@/lib/health";
import { liquidationBonusBps } from "@/content/protocol";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { TextLink } from "@/components/ui/TextLink";

export function HowLiquidationWorksCallout() {
  return (
    <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-4 p-6 sm:flex-row sm:items-center sm:gap-6">
      <span className="grid size-10 shrink-0 place-items-center rounded-tile bg-brand-soft text-brand">
        <Icon name={IconName.Receipt} className="size-5" />
      </span>

      <div className="flex flex-1 flex-col gap-1.5">
        <h2 className="text-base font-semibold tracking-tight text-ink">You do not need a bot to take part</h2>
        <p className="max-w-3xl text-sm leading-relaxed text-ink-soft">
          Pick a position, repay its USDC loan, and receive the borrower&apos;s WETH plus a{" "}
          {formatValue(Number(liquidationBonusBps) / Number(basisPoints), ValueFormat.Percent)} bonus. The rules are the
          same for everyone, and the reward is published rather than negotiated.
        </p>
      </div>

      <TextLink href={AppRoute.LearnLiquidation} trailingIcon={IconName.ArrowRight} className="shrink-0">
        Liquidation rules
      </TextLink>
    </Card>
  );
}
