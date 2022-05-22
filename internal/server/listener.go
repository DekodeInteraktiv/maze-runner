package server

import (
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
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

func (s *Server) startHTTPS() {
	go func() { s.Logger.Fatal(certmagic.HTTPS([]string{"maze.peterbooker.com"}, s.Router)) }()
}
