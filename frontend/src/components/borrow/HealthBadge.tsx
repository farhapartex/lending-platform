import { HealthTier } from "@/lib/enums";
import { Badge } from "@/components/ui/Badge";
import { healthIcons, healthLabels, healthTones } from "@/components/borrow/healthPresentation";

type HealthBadgeProps = {
  tier: HealthTier;
  className?: string;
};

export function HealthBadge({ tier, className }: HealthBadgeProps) {
  return (
    <Badge tone={healthTones[tier]} icon={healthIcons[tier]} className={className}>
      {healthLabels[tier]}
    </Badge>
  );
}
