package server

import (
	"net/http"
	"time"
)

func (s *Server) startHTTP() {
	http := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      s.Router,
		Addr:         ":9071",
	}

	go func() { s.Logger.Fatal(http.ListenAndServe()) }()
}
