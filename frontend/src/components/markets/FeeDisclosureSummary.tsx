import { SurfaceElevation } from "@/lib/enums";
import { feeItems } from "@/content/protocol";
import { Card } from "@/components/ui/Card";
import { MetricRow } from "@/components/ui/MetricRow";

export function FeeDisclosureSummary() {
  return (
    <Card elevation={SurfaceElevation.Flat} className="px-5">
      <dl className="divide-y divide-line">
        {feeItems.map((fee) => (
          <MetricRow key={fee.kind} label={fee.label} value={fee.value} hint={fee.description} emphasised />
        ))}
      </dl>
    </Card>
  );
}
