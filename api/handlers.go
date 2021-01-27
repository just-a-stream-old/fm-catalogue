package api

import "net/http"

func (s *server) getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Welcome"))
	}
}
