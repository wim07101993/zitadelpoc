package server

import "net/http"

// cors is a middleware which adds the allow all header for cors.
func cors(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		handler(res, req)
	}
}
