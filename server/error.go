package server

import "github.com/valyala/fasthttp"

var (
	errInternalServerError = "Internal server error"
	errBadRequest          = "Bad request"
	errForbidden           = "Forbidden"
	errNotFound            = "Not Found"
	errGatewayTimeout      = "Gateway Timeout"
)

func respondWithError(ctx *fasthttp.RequestCtx, message string, code int) {
	ctx.Error(message, code)
}
