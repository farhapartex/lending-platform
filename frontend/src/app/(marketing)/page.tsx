import { SectionId, SectionTone } from "@/lib/enums";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { Hero } from "@/components/marketing/Hero";
import { ProductShowcase } from "@/components/marketing/ProductShowcase";
import { Mechanics } from "@/components/marketing/Mechanics";
import { TrustSignals } from "@/components/marketing/TrustSignals";
import { ClosingCta } from "@/components/marketing/ClosingCta";

export default function Home() {
  return (
    <>
      <Hero />

      <Section id={SectionId.Products} tone={SectionTone.Canvas}>
        <SectionHeading
          sectionId={SectionId.Products}
          eyebrow="Two products"
          title="Pick the one that fits what you are trying to do."
          description="They take the same collateral and answer different questions. One borrows dollars that already exist, the other creates them."
        />
        <div className="mt-10">
          <ProductShowcase />
        </div>
      </Section>

      <Mechanics />
      <TrustSignals />
      <ClosingCta />
    </>
  );
}
