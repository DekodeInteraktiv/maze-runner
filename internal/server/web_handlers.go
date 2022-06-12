package server

import (
	"net/http"
	"time"

	"github.com/PeterBooker/maze-game-server/internal/assets"
)

// App ...
type App struct {
	Name    string
	Version string
	URL     string
}

// Page ...
type Page struct {
	Name        string
	Title       string
	URLPath     string
	Description string
	Time        time.Time
}

func (s *Server) webIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		message := "<p>Welcome to " + s.Config.Name + "</p>" +
			"<p>Version: " + s.Config.Version + "</p>" +
			"<p>Commit:  " + s.Config.Commit + "</p>" +
			"<p>Date:    " + s.Config.Date + "</p>"

		w.Write([]byte(message))
	}
}

func (s *Server) webFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, err := assets.Content.ReadFile("static/favicon.ico")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("File not found."))
			return
		}

		w.Write(data)
	}
}

func (s *Server) webDocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")
		message := "<p>Welcome to the docs page</p>"

		w.Write([]byte(message))
	}
}
