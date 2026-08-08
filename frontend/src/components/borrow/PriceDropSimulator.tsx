import { SectionId } from "@/lib/enums";
import { borrowPageContent } from "@/content/borrow";
import { Card } from "@/components/ui/Card";
import { PriceDropSimulatorBody } from "@/components/borrow/PriceDropSimulatorBody";

export function PriceDropSimulator() {
  return (
    <Card className="flex flex-col gap-5 p-6 sm:p-7">
      <div className="flex flex-col gap-2">
        <h2 id={`${SectionId.BorrowSimulator}-heading`} className="text-lg font-semibold tracking-tight text-ink">
          {borrowPageContent.simulatorTitle}
        </h2>
        <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.simulatorDescription}</p>
      </div>

      <PriceDropSimulatorBody />
    </Card>
  );
}
