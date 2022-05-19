package server

import (
	"fmt"
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

// viewerIndex ...
func (s *Server) viewerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		data, _ := assets.Content.ReadFile("viewer/index.html")

		w.Write(data)
	}
}

// webIndex ...
func (s *Server) webIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		w.Write([]byte("Welcome to the Maze Game..."))
	}
}

// webFavicon ...
func (s *Server) webFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, _ := assets.Content.ReadFile("static/favicon.ico")

		w.Write(data)
	}
}

// webDocs ...
func (s *Server) webDocs() http.HandlerFunc {

	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		URL:     "https://maze-game-server.dev",
	}

	page := Page{
		Name:        "docs",
		Title:       "Documentation - Maze Game Server",
		Description: "Documentation for the Maze Game.",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		page.URLPath = r.URL.Path
		page.Time = time.Now()

		meta := struct {
			Page Page
			App  App
		}{
			page,
			app,
		}

		fmt.Println(meta)
	}
}
