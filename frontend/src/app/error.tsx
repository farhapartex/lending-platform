"use client";

import { ErrorView } from "@/components/system/ErrorView";

export default function RouteError({ error, reset }: { error: Error & { digest?: string }; reset: () => void }) {
  return <ErrorView digest={error.digest} onRetry={reset} />;
}
