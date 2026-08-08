import { HealthTier } from "@/lib/enums";
import { Badge } from "@/components/ui/Badge";
import {
  healthIcons,
  healthLabels,
  healthShortLabels,
  healthTones,
} from "@/components/borrow/healthPresentation";

type HealthBadgeProps = {
  tier: HealthTier;
  compact?: boolean;
  className?: string;
};

export function HealthBadge({ tier, compact = false, className }: HealthBadgeProps) {
  return (
    <Badge tone={healthTones[tier]} icon={healthIcons[tier]} className={className}>
      {compact ? healthShortLabels[tier] : healthLabels[tier]}
    </Badge>
  );
}
