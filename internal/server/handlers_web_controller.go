package server

import (
	"net/http"

	"github.com/PeterBooker/maze-game-server/internal/assets"
)

func (s *Server) controllerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		data, err := assets.Content.ReadFile("controller/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, err := assets.Content.ReadFile("controller/favicon.ico")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerRobots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		data, err := assets.Content.ReadFile("controller/robots.txt")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("controller/manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerAssetManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("controller/asset-manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerLogo192() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("controller/logo192.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) controllerLogo512() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("controller/logo512.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}
