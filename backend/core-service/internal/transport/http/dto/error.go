package dto

const (
	CodeInternalError = "INTERNAL_ERROR"
	CodeNotFound      = "NOT_FOUND"
	CodeBadRequest    = "BAD_REQUEST"
)

type ErrorBody struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

type ErrorResponse struct {
	Error     ErrorBody `json:"error"`
	RequestID string    `json:"request_id,omitempty"`
}

func NewErrorResponse(code, message, requestID string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
		RequestID: requestID,
	}
}
