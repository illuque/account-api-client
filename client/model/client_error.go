package model

import "net/http"

type ClientError struct {
	Code      int64
	Message   string
	Retryable bool
}

func (m ClientError) Error() string {
	return m.Message
}

func NewUnknownError(message string) error {
	return ClientError{
		Code:      http.StatusInternalServerError,
		Message:   message,
		Retryable: false,
	}
}

func NewBadRequest(message string) error {
	return ClientError{
		Code:      http.StatusBadRequest,
		Message:   message,
		Retryable: false,
	}
}

func NewNotFound(message string) error {
	return ClientError{
		Code:      http.StatusNotFound,
		Message:   message,
		Retryable: true,
	}
}

func NewConflict(message string) error {
	return ClientError{
		Code:      http.StatusConflict,
		Message:   message,
		Retryable: false,
	}
}
