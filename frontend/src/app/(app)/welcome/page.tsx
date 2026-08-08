import type { Metadata } from "next";
import { Container } from "@/components/ui/Container";
import { welcomePageContent } from "@/content/welcome";
import { OnboardingWizard } from "@/components/welcome/OnboardingWizard";

export const metadata: Metadata = {
  title: "Getting started",
  description: welcomePageContent.description,
};

export default function WelcomePage() {
  return (
    <Container className="max-w-3xl py-12 sm:py-16">
      <div className="flex flex-col gap-8">
        <header className="flex flex-col gap-3">
          <h1 className="text-3xl font-semibold tracking-tight text-ink sm:text-4xl">{welcomePageContent.title}</h1>
          <p className="text-pretty text-base leading-relaxed text-ink-soft">{welcomePageContent.description}</p>
        </header>

        <OnboardingWizard />
      </div>
    </Container>
  );
}
