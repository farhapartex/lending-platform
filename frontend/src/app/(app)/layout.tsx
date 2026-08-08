import type { ReactNode } from "react";
import { SectionId } from "@/lib/enums";
import { SkipLink } from "@/components/ui/SkipLink";
import { PracticeModeBanner } from "@/components/app/PracticeModeBanner";
import { TopNav } from "@/components/app/TopNav";
import { AppFooter } from "@/components/app/AppFooter";

export default function AppLayout({ children }: { children: ReactNode }) {
  return (
    <>
      <SkipLink />
      <PracticeModeBanner />
      <TopNav />
      <main id={SectionId.MainContent} className="flex-1">
        {children}
      </main>
      <AppFooter />
    </>
  );
}
