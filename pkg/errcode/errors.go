package errcode

var (
	// common errors
	ErrSuccess             = New(0, "ok")
	ErrBadRequest          = New(400, "Bad Request")
	ErrUnauthorized        = New(401, "Unauthorized")
	ErrNotFound            = New(404, "Not Found")
	ErrInternalServerError = New(500, "Internal Server Error")

	// biz errors
	ErrEmailAlreadyUse = New(1001, "The email is already in use.")
)
