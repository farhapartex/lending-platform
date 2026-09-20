import type { Metadata } from "next";
import Link from "next/link";
import { AppRoute, IconName } from "@/lib/enums";
import { loginContent } from "@/content/login";
import { Icon } from "@/components/ui/Icon";
import { Logo } from "@/components/ui/Logo";
import { MetaMaskOption } from "@/components/login/MetaMaskOption";

export const metadata: Metadata = {
  title: "Connect a wallet",
  description:
    "Connect a wallet to use Fusion. There is no sign up, no password and no email — your wallet is your account.",
};

export default function LoginPage() {
  return (
    <div className="grid min-h-screen lg:grid-cols-2">
      <aside className="relative hidden flex-col justify-between overflow-hidden bg-brand-ink px-10 py-12 lg:flex xl:px-14">
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -left-24 -top-24 size-[30rem] rounded-full bg-brand/40 blur-3xl"
        />
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -bottom-32 -right-20 size-[26rem] rounded-full bg-brand/25 blur-3xl"
        />

        <div className="relative w-fit">
          <Logo labelClassName="!text-white" />
        </div>

        <div className="relative flex max-w-md flex-col gap-8">
          <div className="flex flex-col gap-4">
            <h1 className="text-balance text-4xl font-semibold leading-[1.15] tracking-tight text-white">
              {loginContent.title}
            </h1>
            <p className="text-pretty text-base leading-relaxed text-white/70">{loginContent.description}</p>
          </div>

          <ul className="flex flex-col gap-5">
            {loginContent.points.map((point) => (
              <li key={point.title} className="flex gap-3">
                <span className="mt-0.5 grid size-7 shrink-0 place-items-center rounded-tile bg-white/10 text-white">
                  <Icon name={IconName.ShieldCheck} className="size-3.5" />
                </span>
                <div className="flex flex-col gap-1">
                  <span className="text-sm font-semibold text-white">{point.title}</span>
                  <span className="text-sm leading-relaxed text-white/60">{point.body}</span>
                </div>
              </li>
            ))}
          </ul>
        </div>

        <p className="relative text-xs leading-relaxed text-white/45">{loginContent.footnote}</p>
      </aside>

      <main className="flex flex-col bg-canvas px-6 py-10 sm:px-10">
        <div className="flex items-center justify-between lg:hidden">
          <Logo />
        </div>

        <div className="flex flex-1 items-center justify-center py-12">
          <div className="w-full max-w-sm">
            <div className="flex flex-col gap-1.5">
              <h2 className="text-2xl font-semibold tracking-tight text-ink">{loginContent.panelTitle}</h2>
              <p className="text-sm leading-relaxed text-ink-soft">{loginContent.panelDescription}</p>
            </div>

            <div className="mt-8">
              <MetaMaskOption />
            </div>

            <p className="mt-8 text-xs leading-relaxed text-ink-faint lg:hidden">{loginContent.footnote}</p>
          </div>
        </div>

        <div className="flex justify-center">
          <Link
            href={AppRoute.Home}
            className="rounded-pill px-3 py-2 text-sm text-ink-faint transition-colors hover:text-ink outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
          >
            Back to the site
          </Link>
        </div>
      </main>
    </div>
  );
}
