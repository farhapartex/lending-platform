import type { Metadata } from "next";
import { dashboardContent } from "@/content/dashboard";
import { DashboardPositions } from "@/components/dashboard/DashboardPositions";
import { PriceDropDrawer } from "@/components/dashboard/PriceDropDrawer";
import { QuickActionBar } from "@/components/dashboard/QuickActionBar";
import { PortalPage } from "@/components/portal/PortalPage";
import { PriceStalenessWarning } from "@/components/markets/PriceStalenessWarning";

export const metadata: Metadata = {
  title: "Dashboard",
  description: "Your deposits, loan, collateral, and safety score in a single view.",
};

export default function DashboardPage() {
  return (
    <PortalPage title={dashboardContent.title} description={dashboardContent.description} actions={<PriceDropDrawer />}>
      <PriceStalenessWarning />
      <QuickActionBar />
      <DashboardPositions />
    </PortalPage>
  );
}
