import type { Metadata } from "next";
import { LendView } from "@/components/lend/LendView";

export const metadata: Metadata = {
  title: "Lend",
  description: "Deposit USDC to earn interest that accrues every second, and withdraw at any time.",
};

export default function LendPage() {
  return <LendView />;
}
