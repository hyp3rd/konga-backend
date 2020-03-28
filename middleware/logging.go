package middleware

import (
	"github.com/go-kit/kit/log"

	"gitlab.com/hyperd/konga-backend"
)

// Middleware describes the titanic service (as opposed to endpoint) middleware.
type Middleware func(konga.Service) konga.Service

// LoggingMiddleware provides basic logging Middleware
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next konga.Service) konga.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   konga.Service
	logger log.Logger
}
