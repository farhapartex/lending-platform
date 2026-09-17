import type { Metadata } from "next";
import { SectionId, SectionTone, WalletGatePurpose } from "@/lib/enums";
import { dashboardContent } from "@/content/dashboard";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { WalletGate } from "@/components/app/WalletGate";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";
import { DashboardPositions } from "@/components/dashboard/DashboardPositions";
import { PriceDropDrawer } from "@/components/dashboard/PriceDropDrawer";
import { QuickActionBar } from "@/components/dashboard/QuickActionBar";

export const metadata: Metadata = {
  title: "Dashboard",
  description: "Your deposits, loan, collateral, and safety score in a single view.",
};

export default function DashboardPage() {
  return (
    <>
      <PageHeader title={dashboardContent.title} description={dashboardContent.description} aside={<PriceDropDrawer />}>
        <QuickActionBar />
      </PageHeader>

      <PriceStalenessWarning />

      <Section id={SectionId.DashboardOverview} tone={SectionTone.Canvas}>
        <h2 id={`${SectionId.DashboardOverview}-heading`} className="sr-only">
          {dashboardContent.overviewTitle}
        </h2>

        <WalletGate purpose={WalletGatePurpose.PersonalData}>
          <DashboardPositions />
        </WalletGate>
      </Section>
    </>
  );
}
