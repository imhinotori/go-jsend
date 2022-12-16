package jsend

const (
	StatusError   = "error"
	StatusFail    = "fail"
	StatusSuccess = "success"
)

type Body struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func Success(data interface{}) Body {
	return Body{
		Status: StatusSuccess,
		Data:   data,
	}
}

func Fail(data interface{}) Body {
	return Body{
		Status: StatusFail,
		Data:   data,
	}
}

func Error(msg string) Body {
	return Body{
		Status:  StatusError,
		Message: msg,
	}
}

func ErrorWithCode(msg string, code int) Body {
	return Body{
		Status:  StatusError,
		Message: msg,
		Code:    code,
	}
}

func ErrorWithData(msg string, data interface{}) Body {
	return Body{
		Status:  StatusError,
		Message: msg,
		Data:    data,
	}
}

func ErrorFull(msg string, code int, data interface{}) Body {
	return Body{
		Status:  StatusError,
		Message: msg,
		Data:    data,
		Code:    code,
	}
}
