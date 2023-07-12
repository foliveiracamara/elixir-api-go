package app_error

import "net/http"

type AppErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *AppErr) Error() string {
	return r.Message
}

func NewAppErr(message, err string, code int, causes []Causes) *AppErr {
	return &AppErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbinddenError(message string) *AppErr {
	return &AppErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
