import type { Metadata } from "next";
import { BadgeTone, DataStatus, IconName, SectionId, SectionTone, WalletGatePurpose } from "@/lib/enums";
import { historyDataStatus, historyPageContent } from "@/content/history";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { WalletGate } from "@/components/app/WalletGate";
import { HistoryPanel } from "@/components/history/HistoryPanel";

export const metadata: Metadata = {
  title: "History",
  description: "Every deposit, withdrawal, borrow, repayment, and liquidation for your wallet.",
};

export default function HistoryPage() {
  return (
    <>
      <PageHeader title={historyPageContent.title} description={historyPageContent.description} />

      <Section id={SectionId.HistoryList} tone={SectionTone.Canvas}>
        <h2 id={`${SectionId.HistoryList}-heading`} className="sr-only">
          {historyPageContent.listTitle}
        </h2>

        {historyDataStatus === DataStatus.Unavailable ? (
          <Alert title={historyPageContent.unavailableTitle} tone={BadgeTone.Caution} icon={IconName.Warning}>
            {historyPageContent.unavailableDescription}
          </Alert>
        ) : (
          <WalletGate purpose={WalletGatePurpose.PersonalData}>
            <HistoryPanel />
          </WalletGate>
        )}
      </Section>

      <Container className="pb-4">
        <p className="text-xs text-ink-faint">
          Rebuilt from indexed contract events. Exporting your history arrives in a later phase.
        </p>
      </Container>
    </>
  );
}
