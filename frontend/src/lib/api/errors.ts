import { ApiErrorCode } from "@/lib/enums";

type ApiErrorDetails = {
  code: ApiErrorCode;
  message: string;
  status?: number;
  requestId?: string;
};

export class ApiError extends Error {
  readonly code: ApiErrorCode;
  readonly status: number | undefined;
  readonly requestId: string | undefined;

  constructor({ code, message, status, requestId }: ApiErrorDetails) {
    super(message);

    this.name = "ApiError";
    this.code = code;
    this.status = status;
    this.requestId = requestId;
  }
}

export function malformedResponse(field: string): ApiError {
  return new ApiError({
    code: ApiErrorCode.MalformedResponse,
    message: `The server sent a response we could not read (${field}).`,
  });
}

export function isApiError(error: unknown): error is ApiError {
  return error instanceof ApiError;
}

export function isNotFound(error: unknown): boolean {
  return isApiError(error) && error.code === ApiErrorCode.NotFound;
}

export function isBackendUnreachable(error: unknown): boolean {
  return isApiError(error) && (error.code === ApiErrorCode.Unreachable || error.code === ApiErrorCode.Timeout);
}

export function isRetryable(error: unknown): boolean {
  if (!isApiError(error)) {
    return false;
  }

  return (
    error.code === ApiErrorCode.Unreachable ||
    error.code === ApiErrorCode.Timeout ||
    error.code === ApiErrorCode.InternalError
  );
}
