import type { Metadata } from "next";
import { lendPageContent } from "@/content/lend";
import { LendView } from "@/components/lend/LendView";
import { PortalPage } from "@/components/portal/PortalPage";

export const metadata: Metadata = {
  title: "Lend",
  description: "Deposit USDC to earn interest that accrues every second, and withdraw at any time.",
};

export default function LendPage() {
  return (
    <PortalPage title={lendPageContent.title} description={lendPageContent.description}>
      <LendView />
    </PortalPage>
  );
}
