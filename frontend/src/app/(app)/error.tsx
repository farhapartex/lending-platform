"use client";

import { ErrorView } from "@/components/system/ErrorView";

export default function AppRouteError({ error, reset }: { error: Error & { digest?: string }; reset: () => void }) {
  return <ErrorView digest={error.digest} onRetry={reset} />;
}
