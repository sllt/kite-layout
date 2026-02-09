package errcode

// Response is used only for Swagger documentation to describe
// the standard API response format rendered by kite.
// @Description Standard API response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
