import { SurfaceElevation } from "@/lib/enums";
import { rateExplainerPoints } from "@/content/markets";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";

export function RateExplainer() {
  return (
    <ul className="grid gap-4 md:grid-cols-3">
      {rateExplainerPoints.map((point) => (
        <li key={point.key}>
          <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col gap-2.5 p-5">
            <span className="grid size-9 place-items-center rounded-tile bg-accent-soft text-accent">
              <Icon name={point.icon} className="size-4.5" />
            </span>
            <h3 className="text-sm font-semibold tracking-tight text-ink">{point.title}</h3>
            <p className="text-sm leading-relaxed text-ink-soft">{point.description}</p>
          </Card>
        </li>
      ))}
    </ul>
  );
}
