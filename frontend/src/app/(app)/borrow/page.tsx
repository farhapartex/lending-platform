import type { Metadata } from "next";
import { borrowPageContent } from "@/content/borrow";
import { BorrowView } from "@/components/borrow/BorrowView";
import { PortalPage } from "@/components/portal/PortalPage";

export const metadata: Metadata = {
  title: "Borrow",
  description:
    "Borrow USDC against WETH collateral, with a live safety score, a price-drop simulator, and warnings well before liquidation.",
};

export default function BorrowPage() {
  return (
    <PortalPage title={borrowPageContent.title} description={borrowPageContent.description}>
      <BorrowView />
    </PortalPage>
  );
}
