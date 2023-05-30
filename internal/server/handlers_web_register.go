package server

import (
	"net/http"

	"github.com/DekodeInteraktiv/maze-runner/internal/assets"
)

func (s *Server) registerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		data, err := assets.Content.ReadFile("register/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, err := assets.Content.ReadFile("register/favicon.ico")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerRobots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		data, err := assets.Content.ReadFile("register/robots.txt")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("register/manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerAssetManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("register/asset-manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerLogo192() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("register/logo192.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) registerLogo512() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("register/logo512.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}
