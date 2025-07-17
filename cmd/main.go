package main

import (
	"fmt"
	"github.com/wim07101993/zitadelpoc/internal/auth"
	"github.com/wim07101993/zitadelpoc/internal/logging"
	"github.com/wim07101993/zitadelpoc/internal/server"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
var cfg Config
var tokenIntrospectionJwt []byte

func main() {
	if err := cfg.FromFile(); err != nil {
		log.Fatal(err)
	}
	if err := cfg.FromEnv(); err != nil {
		log.Fatal(err)
	}
	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}
	logger.Debug("starting application with config", slog.Any("config", cfg))

	jwt, err := os.ReadFile(cfg.TokenIntrospectionJwtFile)
	if err != nil {
		log.Fatal(err)
	}
	tokenIntrospectionJwt = jwt

	srv := server.NewHttpServer(
		logging.NewMiddleware(logger),
		auth.NewMiddleware(cfg.TokenIntrospectionUrl, tokenIntrospectionJwt))
	srv.RegisterRoutes()

	addr := fmt.Sprintf(":%d", cfg.HttpServerPort)
	logger.Info("start listening for http requests", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Error("failed to serve score scoresIndex",
			slog.Any("error", err))
	}
}
