package web

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (l Log) Error() string {
	return l.Level + " - " + l.Message
}

func (r Response) Error() string {
	return r.Message
}

func (r Response) Header() int {
	return r.Code
}

func ErrorLog(msg string) Log {
	return Log{
		Level:   "ERROR",
		Message: msg,
	}
}

func InfoLog(msg string) Log {
	return Log{
		Level:   "INFO",
		Message: msg,
	}
}

func WarnLog(msg string) Log {
	return Log{
		Level:   "WARN",
		Message: msg,
	}
}

func Success(msg string, data interface{}) Response {
	return Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func Created(msg string, data interface{}) Response {
	return Response{
		Code:    201,
		Message: msg,
		Data:    data,
	}
}

func NotFound(msg string) Response {
	return Response{
		Code:    404,
		Message: msg,
	}
}

func BadRequest(msg string) Response {
	return Response{
		Code:    400,
		Message: msg,
	}
}

func InternalServerError(msg string) Response {
	return Response{
		Code:    500,
		Message: msg,
	}
}

func Unauthorized(msg string) Response {
	return Response{
		Code:    401,
		Message: msg,
	}
}
