package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/lab259/cors"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

type ctxKey int8

// Server ...
type Server struct {
	router       *router.Router
	logger       *logrus.Logger
	sessionStore sessions.Store
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	TLSConfig    *tls.Config
}

// Start function starts server using provided config
func Start(config *Config) error {
	sessionStore := cookie.NewStore([]byte(config.SessionKey))
	_, handler := NewServer(sessionStore)

	return fasthttp.ListenAndServe(config.BindAddress, handler)
}

func NewServer(sessionStore sessions.Store) (*Server, fasthttp.RequestHandler) {
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	s := &Server{
		router:       router.New(),
		logger:       logrus.New(),
		sessionStore: sessionStore,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
	}

	handler := s.configureRouter()

	return s, handler
}

func (s *Server) configureRouter() fasthttp.RequestHandler {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:8080"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
	})

	s.router.GET("/", s.logRequest(s.Index))
	handler := cors.Handler(s.router.Handler)

	return handler
}

func (s *Server) Index(ctx *fasthttp.RequestCtx) {
	_, err := ctx.WriteString("Welcome!")
	if err != nil {
		respondWithError(ctx, errInternalServerError, http.StatusInternalServerError)
	}
}
