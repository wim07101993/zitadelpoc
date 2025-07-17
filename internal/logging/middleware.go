package logging

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

const CorrelationIdKey = "correlationId"

type Middleware struct {
	Logger *slog.Logger
}

func NewMiddleware(logger *slog.Logger) *Middleware {
	return &Middleware{
		Logger: logger,
	}
}

func (m *Middleware) Log(handler func(res http.ResponseWriter, req *http.Request) error) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		correlation := uuid.New().String()
		m.Logger.Info("handle http request",
			slog.String("method", req.Method),
			slog.String("pattern", req.Pattern),
			slog.String("uri", req.RequestURI),
			slog.String("correlationId", correlation))

		req = req.WithContext(context.WithValue(req.Context(), CorrelationIdKey, correlation))
		loggingRes := NewResponseWriter(m.Logger, res, correlation)
		defer loggingRes.Flush()

		loggingRes.Err = handler(loggingRes, req)
	}
}
