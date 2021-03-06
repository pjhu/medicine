package errors

import (
	"errors"
	"net/http"
)

type ErrorWithCode struct {
	Code string
	Err error
}

type ErrorResponse struct {
	Code string `json:"code"`
	ErrorMessage string `json:"message"`
}

func (r *ErrorWithCode) Error() string {
	return r.Err.Error()
}

// GetHTTPStatus return a HTTP status
func (r *ErrorWithCode) GetHTTPStatus() int {
	return codeMappingWithHTTTPStatus[r.Code]
}

// GetErrorResponse return a HTTP response
func (r *ErrorWithCode) GetErrorResponse() *ErrorResponse {
	
	return &ErrorResponse {
		Code: r.Code, 
		ErrorMessage: r.Err.Error()}
}

// NewErrorWithCode create new error with error code
func NewErrorWithCode(errorCode string, errorMessage string) *ErrorWithCode {
	return &ErrorWithCode {
		Code: errorCode, 
		Err: errors.New(errorMessage)}
}

const (
	// server error code
	SystemInternalError = "SYSTEM_INTERNAL_ERROR"
	DatabaseError = "DATABASE_ERROR"

	// client error code
	InvalidParameter = "INVALID_PARAMETER"
	ValidationFailed = "VALIDATION_FAILED"
	AlreadySignedInError  = "ALREADY_SIGNED_IN_ERROR"
	Unauthorized = "UNAUTHORIZED"
	Forbidden = "FORBIDDEN"
	GetAuthorizerTokenRrror = "GET_AUTHORIZER_TOKEN_ERROR"
	NoAuthorizerError = "NO_AUTHORIZER_ERROR"
)

var codeMappingWithHTTTPStatus = map[string]int {
	SystemInternalError:           http.StatusInternalServerError,
	DatabaseError:                 http.StatusInternalServerError,
	Unauthorized:                  http.StatusInternalServerError,
}


