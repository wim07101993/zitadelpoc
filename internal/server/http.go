package server

import (
	"github.com/wim07101993/zitadelpoc/internal/auth"
	"github.com/wim07101993/zitadelpoc/internal/logging"
	"net/http"
)

type HttpServer struct {
	Logger *logging.Middleware
	Auth   *auth.Middleware
}

func NewHttpServer(log *logging.Middleware, auth *auth.Middleware) *HttpServer {
	return &HttpServer{
		Logger: log,
		Auth:   auth,
	}
}

func (serv *HttpServer) RegisterRoutes() {
	http.HandleFunc("/healthz", cors(serv.Logger.Log(healthz)))
	http.HandleFunc("/secret", cors(serv.Logger.Log(serv.Auth.Authenticate(secret))))
}

func healthz(res http.ResponseWriter, _ *http.Request) error {
	res.WriteHeader(200)
	_, err := res.Write([]byte("OK"))
	return err
}

func secret(res http.ResponseWriter, _ *http.Request) error {
	res.WriteHeader(200)
	_, err := res.Write([]byte("THIS IS SECRET"))
	return err
}
