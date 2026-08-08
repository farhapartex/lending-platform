import type { ReactNode } from "react";
import { Container } from "@/components/ui/Container";
import { DocsSidebar } from "@/components/learn/DocsSidebar";

export default function LearnLayout({ children }: { children: ReactNode }) {
  return (
    <Container className="grid gap-10 py-12 lg:grid-cols-[minmax(0,14rem)_minmax(0,1fr)] lg:gap-14">
      <aside className="lg:sticky lg:top-24 lg:self-start">
        <DocsSidebar />
      </aside>
      <div className="min-w-0">{children}</div>
    </Container>
  );
}
