package server

import (
	"net/http"

	"github.com/DekodeInteraktiv/maze-runner/internal/assets"
)

func (s *Server) leaderboardIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		data, err := assets.Content.ReadFile("leaderboard/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, err := assets.Content.ReadFile("leaderboard/favicon.ico")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardRobots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		data, err := assets.Content.ReadFile("leaderboard/robots.txt")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("leaderboard/manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardAssetManifest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := assets.Content.ReadFile("leaderboard/asset-manifest.json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardLogo192() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("leaderboard/logo192.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) leaderboardLogo512() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		data, err := assets.Content.ReadFile("leaderboard/logo512.png")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}
