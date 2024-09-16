package models

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewCustomError(message error, code int) *CustomError {
	return &CustomError{
		Message: message.Error(),
		Code:    code,
	}
}
