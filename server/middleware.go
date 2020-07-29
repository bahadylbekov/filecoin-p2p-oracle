package server

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func (s *Server) SetRequestID(ctx *fasthttp.RequestCtx) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		id := uuid.New().String()
		ctx.Response.Header.Add("X-Request-ID", id)
		ctx.SetUserValue("ctxKeyRequestID", id)
	}
}

func (s *Server) logRequest(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		s.SetRequestID(ctx)
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": ctx.RemoteAddr(),
			"request_id":  ctx.Value("ctxKeyRequestID"),
		})
		logger.Infof("started %s %s", ctx.Method, ctx.Request.RequestURI)
		start := time.Now()
		handler(ctx)

		var level logrus.Level
		switch {
		case ctx.Response.StatusCode() >= 500:
			level = logrus.ErrorLevel
		case ctx.Response.StatusCode() >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			ctx.Response.StatusCode(),
			http.StatusText(ctx.Response.StatusCode()),
			time.Since(start),
		)
	}
}
