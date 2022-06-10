package server

import (
	"net/http"

	"github.com/PeterBooker/maze-game-server/internal/assets"
)

func (s *Server) viewerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		data, err := assets.Content.ReadFile("viewer/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, err := assets.Content.ReadFile("viewer/favicon.ico")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerRobots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		data, err := assets.Content.ReadFile("viewer/robots.txt")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("viewer/manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerAssetManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("viewer/asset-manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerLogo192() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("viewer/logo192.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) viewerLogo512() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("viewer/logo512.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}
