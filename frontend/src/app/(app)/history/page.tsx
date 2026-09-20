import type { Metadata } from "next";
import { historyPageContent } from "@/content/history";
import { HistoryPanel } from "@/components/history/HistoryPanel";
import { PortalPage } from "@/components/portal/PortalPage";

export const metadata: Metadata = {
  title: "History",
  description: "Every deposit, withdrawal, borrow, repayment, and liquidation for your wallet.",
};

export default function HistoryPage() {
  return (
    <PortalPage title={historyPageContent.title} description={historyPageContent.description}>
      <HistoryPanel />

      <p className="text-xs leading-relaxed text-ink-faint">
        Rebuilt from indexed contract events. Exporting your history arrives in a later phase.
      </p>
    </PortalPage>
  );
}
