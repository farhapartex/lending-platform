import type { Metadata } from "next";
import { BadgeTone, SectionId, SectionTone } from "@/lib/enums";
import { practicePageContent } from "@/content/practice";
import { Badge } from "@/components/ui/Badge";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { NetworkSwitchCard } from "@/components/practice/NetworkSwitchCard";
import { PracticeModeExplainer } from "@/components/practice/PracticeModeExplainer";
import { ReturnToLiveCta } from "@/components/practice/ReturnToLiveCta";
import { TestnetFaucetCard } from "@/components/practice/TestnetFaucetCard";
import { ThingsToTryList } from "@/components/practice/ThingsToTryList";

export const metadata: Metadata = {
  title: "Practice mode",
  description: practicePageContent.description,
};

export default function PracticePage() {
  return (
    <>
      <PageHeader
        title={practicePageContent.title}
        description={practicePageContent.description}
        badge={<Badge tone={BadgeTone.Positive}>No real money</Badge>}
      />

      <Section id={SectionId.PracticeSetup} tone={SectionTone.Surface} bordered>
        <SectionHeading
          sectionId={SectionId.PracticeSetup}
          eyebrow="Setup"
          title={practicePageContent.setupTitle}
          description={practicePageContent.setupDescription}
        />

        <div className="mt-8 grid gap-6 lg:grid-cols-2">
          <NetworkSwitchCard />
          <TestnetFaucetCard />
        </div>
      </Section>

      <PracticeModeExplainer />

      <ThingsToTryList />

      <ReturnToLiveCta />
    </>
  );
}
