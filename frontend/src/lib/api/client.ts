import { ApiErrorCode } from "@/lib/enums";
import { apiRequestTimeoutMs, apiUrl } from "@/lib/api/config";
import { ApiError, isApiError } from "@/lib/api/errors";
import type { WireErrorEnvelope } from "@/lib/api/wire";

const requestIdHeader = "X-Request-Id";

const codesByWireValue: Record<string, ApiErrorCode> = {
  [ApiErrorCode.BadRequest]: ApiErrorCode.BadRequest,
  [ApiErrorCode.NotFound]: ApiErrorCode.NotFound,
  [ApiErrorCode.InternalError]: ApiErrorCode.InternalError,
};

const codesByStatus: Record<number, ApiErrorCode> = {
  400: ApiErrorCode.BadRequest,
  404: ApiErrorCode.NotFound,
};

const messagesByCode: Record<ApiErrorCode, string> = {
  [ApiErrorCode.BadRequest]: "That request could not be read.",
  [ApiErrorCode.NotFound]: "We could not find that.",
  [ApiErrorCode.InternalError]: "Something went wrong on the server.",
  [ApiErrorCode.Unreachable]: "We could not reach the server.",
  [ApiErrorCode.Timeout]: "The server took too long to respond.",
  [ApiErrorCode.MalformedResponse]: "The server sent a response we could not read.",
};

type TimedRequest = {
  signal: AbortSignal;
  timedOut: () => boolean;
  dispose: () => void;
};

function withTimeout(external: AbortSignal | undefined, timeoutMs: number): TimedRequest {
  const controller = new AbortController();
  let expired = false;

  const timer = setTimeout(() => {
    expired = true;
    controller.abort();
  }, timeoutMs);

  const forwardAbort = () => controller.abort();

  if (external !== undefined) {
    if (external.aborted) {
      controller.abort();
    } else {
      external.addEventListener("abort", forwardAbort, { once: true });
    }
  }

  return {
    signal: controller.signal,
    timedOut: () => expired,
    dispose: () => {
      clearTimeout(timer);
      external?.removeEventListener("abort", forwardAbort);
    },
  };
}

function codeFor(wireCode: string | undefined, status: number): ApiErrorCode {
  if (wireCode !== undefined && wireCode in codesByWireValue) {
    return codesByWireValue[wireCode];
  }

  return codesByStatus[status] ?? ApiErrorCode.InternalError;
}

async function readEnvelope(response: Response): Promise<WireErrorEnvelope | undefined> {
  try {
    return (await response.json()) as WireErrorEnvelope;
  } catch {
    return undefined;
  }
}

async function toApiError(response: Response): Promise<ApiError> {
  const envelope = await readEnvelope(response);
  const code = codeFor(envelope?.error?.code, response.status);

  return new ApiError({
    code,
    message: envelope?.error?.message ?? messagesByCode[code],
    status: response.status,
    requestId: envelope?.request_id ?? response.headers.get(requestIdHeader) ?? undefined,
  });
}

function wasAborted(error: unknown): boolean {
  return error instanceof Error && error.name === "AbortError";
}

function asThrowable(error: unknown, timedOut: boolean): unknown {
  if (isApiError(error)) {
    return error;
  }

  if (timedOut) {
    return new ApiError({ code: ApiErrorCode.Timeout, message: messagesByCode[ApiErrorCode.Timeout] });
  }

  if (wasAborted(error)) {
    return error;
  }

  return new ApiError({ code: ApiErrorCode.Unreachable, message: messagesByCode[ApiErrorCode.Unreachable] });
}

async function readJson<T>(response: Response): Promise<T> {
  try {
    return (await response.json()) as T;
  } catch {
    throw new ApiError({
      code: ApiErrorCode.MalformedResponse,
      message: messagesByCode[ApiErrorCode.MalformedResponse],
      status: response.status,
      requestId: response.headers.get(requestIdHeader) ?? undefined,
    });
  }
}

export async function getJson<T>(path: string, signal?: AbortSignal): Promise<T> {
  const request = withTimeout(signal, apiRequestTimeoutMs);

  try {
    const response = await fetch(apiUrl(path), {
      method: "GET",
      headers: { Accept: "application/json" },
      signal: request.signal,
    });

    if (!response.ok) {
      throw await toApiError(response);
    }

    return await readJson<T>(response);
  } catch (error) {
    throw asThrowable(error, request.timedOut());
  } finally {
    request.dispose();
  }
}
