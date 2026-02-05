package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "The email is already in use.")
)

// StatusCode returns the appropriate HTTP status code for the error.
// For common HTTP errors (400-599), it returns the BizCode directly.
// For business errors (>= 1000), it returns 400 (Bad Request).
func (e *Error) StatusCode() int {
	if e.BizCode >= 400 && e.BizCode < 600 {
		return e.BizCode
	}
	return 400
}
