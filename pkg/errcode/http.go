package errcode

import (
	"errors"
	"net/http"

	kiteHTTP "github.com/sllt/kite/pkg/kite/http"
)

// AsError normalizes arbitrary errors into layout's public error contract.
//
// Handlers and services should return *Error values directly when the error is
// safe to expose. Unknown errors are intentionally converted to the generic
// internal-server error so HTTP middleware does not leak implementation details.
func AsError(err error) *Error {
	if err == nil {
		return nil
	}

	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr
	}

	return ErrInternalServerError
}

// WriteHTTPError renders the same response envelope that Kite handlers use:
//
//	{"code": <business-code>, "data": null, "message": "..."}
//
// Use this helper in net/http middleware, where errors cannot be returned to
// Kite's normal handler responder.
func WriteHTTPError(w http.ResponseWriter, r *http.Request, err error) {
	method := http.MethodGet
	if r != nil {
		method = r.Method
	}

	kiteHTTP.NewResponder(w, method).Respond(nil, AsError(err))
}
