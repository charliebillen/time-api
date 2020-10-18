package api

import (
	"encoding/json"
	"net/http"
)

// Server encapsulates the server dependencies
type Server struct {
	TimeProvider TimeProvider
}

// ServeHTTP implements http.Handler for *Server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/time":
		handleGetTime(w, s.TimeProvider)
	default:
		handleNotFound(w)
	}
}

func handleNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func handleGetTime(w http.ResponseWriter, tp TimeProvider) {
	t := tp()
	rsp := response{
		Hour:   t.Hour(),
		Minute: t.Minute(),
		Second: t.Second(),
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}

type response struct {
	Hour   int
	Minute int
	Second int
}
