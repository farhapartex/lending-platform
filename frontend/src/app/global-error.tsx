"use client";

import { errorContent } from "@/content/system";

export default function GlobalError({ error, reset }: { error: Error & { digest?: string }; reset: () => void }) {
  return (
    <html lang="en">
      <body style={{ margin: 0, backgroundColor: "#f8fafd", color: "#191d2b", fontFamily: "system-ui, sans-serif" }}>
        <main style={{ margin: "0 auto", maxWidth: "36rem", padding: "5rem 1.5rem" }}>
          <h1 style={{ fontSize: "1.75rem", fontWeight: 600, letterSpacing: "-0.01em", margin: "0 0 0.75rem" }}>
            {errorContent.title}
          </h1>
          <p style={{ lineHeight: 1.6, color: "#545c72", margin: "0 0 1rem" }}>{errorContent.description}</p>
          <p style={{ lineHeight: 1.6, color: "#0a684a", margin: "0 0 1.5rem" }}>{errorContent.fundsNote}</p>
          <button
            type="button"
            onClick={reset}
            style={{
              height: "2.75rem",
              padding: "0 1.25rem",
              borderRadius: "9999px",
              border: "none",
              backgroundColor: "#5b5bd6",
              color: "#ffffff",
              fontSize: "0.875rem",
              fontWeight: 500,
              cursor: "pointer",
            }}
          >
            {errorContent.retryCta}
          </button>
          {error.digest === undefined ? null : (
            <p style={{ marginTop: "1.5rem", fontSize: "0.75rem", color: "#848ca1" }}>
              {errorContent.digestLabel}: {error.digest}
            </p>
          )}
        </main>
      </body>
    </html>
  );
}
