import { SectionId } from "@/lib/enums";
import { SkipLink } from "@/components/ui/SkipLink";
import { MarketingNav } from "@/components/marketing/MarketingNav";
import { Hero } from "@/components/marketing/Hero";
import { ProtocolStatsStrip } from "@/components/marketing/ProtocolStatsStrip";
import { ValueProps } from "@/components/marketing/ValueProps";
import { HowItWorksSteps } from "@/components/marketing/HowItWorksSteps";
import { FeeTransparencyTeaser } from "@/components/marketing/FeeTransparencyTeaser";
import { TrustSignals } from "@/components/marketing/TrustSignals";
import { PracticeModeCta } from "@/components/marketing/PracticeModeCta";
import { MarketingFooter } from "@/components/marketing/MarketingFooter";

export default function Home() {
  return (
    <>
      <SkipLink />
      <MarketingNav />
      <main id={SectionId.MainContent} className="flex-1">
        <Hero />
        <ProtocolStatsStrip />
        <ValueProps />
        <HowItWorksSteps />
        <FeeTransparencyTeaser />
        <TrustSignals />
        <PracticeModeCta />
      </main>
      <MarketingFooter />
    </>
  );
}
