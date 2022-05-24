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
