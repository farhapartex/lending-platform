const fallbackBaseUrl = "http://localhost:8080/api/v1";

function resolveBaseUrl(): string {
  const configured = process.env.NEXT_PUBLIC_API_BASE_URL?.trim();

  if (configured === undefined || configured === "") {
    return fallbackBaseUrl;
  }

  return configured.replace(/\/+$/, "");
}

export const apiBaseUrl = resolveBaseUrl();

export const apiRequestTimeoutMs = 8_000;

export function apiUrl(path: string): string {
  return path.startsWith("/") ? `${apiBaseUrl}${path}` : `${apiBaseUrl}/${path}`;
}
