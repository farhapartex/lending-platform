import { HealthTier } from "@/lib/enums";
import { basisPoints, healthTierLowerBounds, healthTierOrder, healthTierUpperBounds } from "@/lib/health";
import { healthExplanations } from "@/components/borrow/healthPresentation";
import { HealthBadge } from "@/components/borrow/HealthBadge";

function boundToText(value: bigint | null): string {
  return value === null ? "" : (Number(value) / Number(basisPoints)).toFixed(2);
}

function rangeText(tier: HealthTier): string {
  const lower = healthTierLowerBounds[tier];
  const upper = healthTierUpperBounds[tier];

  if (lower !== null && upper === null) {
    return `${boundToText(lower)} and above`;
  }

  if (lower === null && upper !== null) {
    return `Below ${boundToText(upper)}`;
  }

  return `${boundToText(lower)} to ${boundToText(upper)}`;
}

export function HealthTierTable() {
  return (
    <div className="overflow-hidden rounded-card border border-line bg-surface">
      <div className="overflow-x-auto">
        <table className="w-full min-w-[38rem] border-collapse text-left">
          <caption className="sr-only">What each safety level means</caption>
          <thead>
            <tr className="bg-surface-muted">
              <th scope="col" className="whitespace-nowrap px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                Level
              </th>
              <th scope="col" className="whitespace-nowrap px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                Health factor
              </th>
              <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                What it means
              </th>
            </tr>
          </thead>
          <tbody>
            {healthTierOrder.map((tier) => (
              <tr key={tier} className="border-t border-line align-top">
                <td className="whitespace-nowrap px-5 py-4">
                  <HealthBadge tier={tier} compact />
                </td>
                <td className="whitespace-nowrap px-5 py-4 text-sm text-ink tabular-nums">{rangeText(tier)}</td>
                <td className="px-5 py-4 text-sm leading-relaxed text-ink-soft">{healthExplanations[tier]}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
