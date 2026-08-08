import { BadgeTone, HealthTier, IconName } from "@/lib/enums";

export const healthLabels: Record<HealthTier, string> = {
  [HealthTier.Safe]: "Safe",
  [HealthTier.Caution]: "Caution",
  [HealthTier.AtRisk]: "At risk",
  [HealthTier.Liquidatable]: "Liquidation possible now",
};

export const healthTones: Record<HealthTier, BadgeTone> = {
  [HealthTier.Safe]: BadgeTone.Positive,
  [HealthTier.Caution]: BadgeTone.Caution,
  [HealthTier.AtRisk]: BadgeTone.Critical,
  [HealthTier.Liquidatable]: BadgeTone.Critical,
};

export const healthIcons: Record<HealthTier, IconName> = {
  [HealthTier.Safe]: IconName.ShieldCheck,
  [HealthTier.Caution]: IconName.Info,
  [HealthTier.AtRisk]: IconName.Warning,
  [HealthTier.Liquidatable]: IconName.Warning,
};

export const healthExplanations: Record<HealthTier, string> = {
  [HealthTier.Safe]: "Your collateral comfortably covers your loan. There is room for the price to move against you.",
  [HealthTier.Caution]:
    "Still safe, but the buffer is thinner than we would recommend. A further price fall would push this into risky territory.",
  [HealthTier.AtRisk]:
    "A small further drop in the WETH price would make your position eligible for liquidation. Add collateral or repay part of the loan now.",
  [HealthTier.Liquidatable]:
    "Your position can be liquidated right now. Repaying or adding collateral immediately is the only way to stop it.",
};

export const healthTextClasses: Record<HealthTier, string> = {
  [HealthTier.Safe]: "text-mint-ink",
  [HealthTier.Caution]: "text-amber-ink",
  [HealthTier.AtRisk]: "text-rose-ink",
  [HealthTier.Liquidatable]: "text-rose-ink",
};

export const healthFillClasses: Record<HealthTier, string> = {
  [HealthTier.Safe]: "bg-mint",
  [HealthTier.Caution]: "bg-amber",
  [HealthTier.AtRisk]: "bg-rose",
  [HealthTier.Liquidatable]: "bg-rose",
};
