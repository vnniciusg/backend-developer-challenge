package responseshttp

import "net/http"

type RestSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
}

func NewSuccess(message string, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Data:    data,
	}
}

func NewCreated(message string, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Data:    data,
		Code:    http.StatusCreated,
	}
}

func NewNoContent(message string) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Code:    http.StatusNoContent,
	}
}

func NewAccepted(message string) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Code:    http.StatusAccepted,
	}
}

func NewOk(message string, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
	}
}
