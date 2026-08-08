import { Container } from "@/components/ui/Container";
import { Skeleton } from "@/components/ui/Skeleton";

export default function AppLoading() {
  return (
    <div aria-busy="true" aria-live="polite">
      <span className="sr-only">Loading</span>

      <div className="border-b border-line bg-surface">
        <Container className="flex flex-col gap-8 py-10 lg:flex-row lg:items-start lg:justify-between">
          <div className="flex w-full flex-col gap-4">
            <Skeleton className="h-9 w-72 max-w-full" />
            <Skeleton className="h-4 w-full max-w-xl" />
            <Skeleton className="h-4 w-2/3 max-w-md" />
          </div>
          <Skeleton className="h-32 w-full rounded-card lg:w-56" />
        </Container>
      </div>

      <Container className="py-16">
        <div className="flex flex-col gap-6">
          <Skeleton className="h-6 w-48" />
          <div className="grid gap-6 sm:grid-cols-3">
            <Skeleton className="h-28 rounded-card" />
            <Skeleton className="h-28 rounded-card" />
            <Skeleton className="h-28 rounded-card" />
          </div>
          <Skeleton className="h-64 rounded-card" />
        </div>
      </Container>
    </div>
  );
}
