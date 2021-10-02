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

// TODO:I manejar resto de errores
