package client_error

import "net/http"

type ErrorData struct {
	Code      int64  `json:"code,required"`
	Message   string `json:"message,required"`
	Retryable bool   `json:"retryable,required"`
}

func NewUnknownClientError(message string) *ErrorData {
	return &ErrorData{
		Code:      http.StatusInternalServerError,
		Message:   message,
		Retryable: false,
	}
}

func NewBadRequest(message string) *ErrorData {
	return &ErrorData{
		Code:      http.StatusBadRequest,
		Message:   message,
		Retryable: false,
	}
}

func NewNotFound(message string) *ErrorData {
	return &ErrorData{
		Code:      http.StatusNotFound,
		Message:   message,
		Retryable: true,
	}
}

func NewConflict(message string) *ErrorData {
	return &ErrorData{
		Code:      http.StatusConflict,
		Message:   message,
		Retryable: false,
	}
}

func NewFromApiError(code int, message string) *ErrorData {
	switch code {
	case http.StatusNotFound:
		return NewNotFound(message)
	case http.StatusBadRequest:
		return NewBadRequest(message)
	case http.StatusConflict:
		return NewConflict(message)
	default:
		return newUnknownApiError(message)
	}
}

func newUnknownApiError(message string) *ErrorData {
	return &ErrorData{
		Code:      http.StatusInternalServerError,
		Message:   message,
		Retryable: true, // according to API spec, errors are retryable
	}
}

// TODO:I manejar resto de errores
