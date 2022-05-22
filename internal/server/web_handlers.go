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

// viewerFavicon ...
func (s *Server) viewerFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, _ := assets.Content.ReadFile("viewer/favicon.ico")

		w.Write(data)
	}
}

// controllerIndex ...
func (s *Server) controllerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		data, _ := assets.Content.ReadFile("controller/index.html")

		w.Write(data)
	}
}

// controllerFavicon ...
func (s *Server) controllerFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		data, _ := assets.Content.ReadFile("controller/favicon.ico")

		w.Write(data)
	}
}

// webIndex ...
func (s *Server) webIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Vary", "Accept-Encoding")

		message := "<p>Welcome to the " + s.Config.Name + "</p>" +
			"<p>Version: " + s.Config.Version + "</p>" +
			"<p>Commit:  " + s.Config.Commit + "</p>" +
			"<p>Date:    " + s.Config.Date + "</p>"

		w.Write([]byte(message))
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
