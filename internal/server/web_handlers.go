package server

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"time"
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

// webIndex ...
func (s *Server) webIndex() http.HandlerFunc {

	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		URL:     "https://maze-game-server.dev",
	}

	page := Page{
		Name:        "index",
		Title:       "Maze Game",
		Description: "Maze Game for developers.",
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

// webFavicon ...
func (s *Server) webFavicon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		img := s.Game.GetImage()

		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		if err != nil {
			log.Panicf("Failed to encode image: %v\n", err)
		}

		w.Write(buf.Bytes())
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
