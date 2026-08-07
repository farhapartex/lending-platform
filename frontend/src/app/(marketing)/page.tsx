import { Hero } from "@/components/marketing/Hero";
import { ProtocolStatsStrip } from "@/components/marketing/ProtocolStatsStrip";
import { ValueProps } from "@/components/marketing/ValueProps";
import { HowItWorksSteps } from "@/components/marketing/HowItWorksSteps";
import { FeeTransparencyTeaser } from "@/components/marketing/FeeTransparencyTeaser";
import { TrustSignals } from "@/components/marketing/TrustSignals";
import { PracticeModeCta } from "@/components/marketing/PracticeModeCta";

export default function Home() {
  return (
    <>
      <Hero />
      <ProtocolStatsStrip />
      <ValueProps />
      <HowItWorksSteps />
      <FeeTransparencyTeaser />
      <TrustSignals />
      <PracticeModeCta />
    </>
  );
}
