import type { Metadata } from "next";
import { DocKey } from "@/lib/enums";
import { findDocPage } from "@/content/learn";
import { DocPageView } from "@/components/learn/DocPageView";

const page = findDocPage(DocKey.Liquidation);

export const metadata: Metadata = {
  title: page.title,
  description: page.summary,
};

export default function DocRoute() {
  return <DocPageView page={page} />;
}
