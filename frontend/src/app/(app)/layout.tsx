import type { ReactNode } from "react";
import { SkipLink } from "@/components/ui/SkipLink";
import { AdminAuthGuard } from "@/components/admin/AdminAuthGuard";
import { AdminShell } from "@/components/admin/AdminShell";
import { PracticeModeBanner } from "@/components/app/PracticeModeBanner";
import { WrongNetworkBanner } from "@/components/app/WrongNetworkBanner";

export default function AppLayout({ children }: { children: ReactNode }) {
  return (
    <>
      <SkipLink />
      <AdminAuthGuard>
        <AdminShell>
          <PracticeModeBanner />
          <WrongNetworkBanner />
          {children}
        </AdminShell>
      </AdminAuthGuard>
    </>
  );
}
