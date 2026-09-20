"use client";

import { useEffect, type ReactNode } from "react";
import { useRouter } from "next/navigation";
import { AppRoute, WalletStatus } from "@/lib/enums";
import { useWalletState } from "@/hooks/useWalletState";

export function AdminAuthGuard({ children }: { children: ReactNode }) {
  const router = useRouter();
  const { status, isSettling } = useWalletState();

  const isSignedOut = !isSettling && status === WalletStatus.Disconnected;

  useEffect(() => {
    if (isSignedOut) {
      router.replace(AppRoute.Login);
    }
  }, [isSignedOut, router]);

  if (isSettling || isSignedOut) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-canvas px-6">
        <div className="flex flex-col items-center gap-3">
          <span className="size-8 animate-spin rounded-full border-2 border-shell-line border-t-accent" />
          <p className="text-sm text-ink-faint">
            {isSignedOut ? "Taking you to the connect screen" : "Checking your wallet"}
          </p>
        </div>
      </div>
    );
  }

  return <>{children}</>;
}
