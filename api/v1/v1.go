package v1

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	BizCode int
	Message string
}

func newError(code int, msg string) error {
	return &Error{
		BizCode: code,
		Message: msg,
	}
}

func (e *Error) Error() string {
	return e.Message
}

// Code implements kite's http.CodeResponder interface.
// Returns the business error code for the JSON response "code" field.
func (e *Error) Code() int {
	return e.BizCode
}
