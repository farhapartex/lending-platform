import type { Metadata } from "next";
import { BorrowView } from "@/components/borrow/BorrowView";

export const metadata: Metadata = {
  title: "Borrow",
  description:
    "Borrow USDC against WETH collateral, with a live safety score, a price-drop simulator, and warnings well before liquidation.",
};

export default function BorrowPage() {
  return <BorrowView />;
}
