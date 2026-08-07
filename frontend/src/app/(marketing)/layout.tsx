import type { ReactNode } from "react";
import { SectionId } from "@/lib/enums";
import { SkipLink } from "@/components/ui/SkipLink";
import { MarketingNav } from "@/components/marketing/MarketingNav";
import { MarketingFooter } from "@/components/marketing/MarketingFooter";

export default function MarketingLayout({ children }: { children: ReactNode }) {
  return (
    <>
      <SkipLink />
      <MarketingNav />
      <main id={SectionId.MainContent} className="flex-1">
        {children}
      </main>
      <MarketingFooter />
    </>
  );
}
