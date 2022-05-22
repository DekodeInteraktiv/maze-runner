package server

import (
	"net/http"

	"github.com/PeterBooker/maze-game-server/internal/assets"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) routes() {
	// Add Routes
	s.Router.Get("/", s.webIndex())
	s.Router.Get("/favicon.ico", s.webFavicon())

	creds := map[string]string{"username": "admin", "password": "admin"}

	// Viewer
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("viewer", creds))
		r.Get("/viewer", s.viewerIndex())
		r.Get("/viewer/", s.viewerIndex())
		r.Get("/viewer/{id}", s.viewerIndex())
		r.Get("/viewer/favicon.ico", s.viewerFavicon())
		r.Handle("/viewer/static/*", http.FileServer(http.FS(assets.Content)))
	})

	// Controller
	s.Router.Get("/controller", s.controllerIndex())
	s.Router.Get("/controller/", s.controllerIndex())
	s.Router.Get("/controller/favicon.ico", s.controllerFavicon())
	s.Router.Handle("/controller/static/*", http.FileServer(http.FS(assets.Content)))

	// Add API v1 routes
	s.Router.Mount("/api/v1", s.apiRoutes())

	// Handle NotFound
	//s.Router.NotFound(s.notFound())
}

func (s *Server) apiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/game", func(r chi.Router) {
		r.Post("/create", s.gameCreate())

		r.Route("/{gameID}", func(r chi.Router) {
			r.Get("/info", s.gameInfo())
			r.Get("/start", s.gameStart())

			r.Route("/player", func(r chi.Router) {
				r.Post("/register/{password}", s.playerCreate())

				r.Post("/move", s.playerMove())
				r.Post("/status", s.playerStatus())
			})
		})
	})

	return r
}
