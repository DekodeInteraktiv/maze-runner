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
	s.Router.Get("/docs", s.webDocs())
	s.Router.Get("/favicon.ico", s.webFavicon())

	// Viewer App
	viewerCreds := map[string]string{"viewer": "d3kode"}
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("viewer", viewerCreds))
		r.Get("/viewer", s.viewerIndex())
		r.Get("/viewer/", s.viewerIndex())
		r.Get("/viewer/{id}", s.viewerIndex())
		r.Get("/viewer/favicon.ico", s.viewerFavicon())
		r.Handle("/viewer/static/*", http.FileServer(http.FS(assets.Content)))
	})

	// Controller App
	controllerCreds := map[string]string{"controller": "happyhour"}
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("controller", controllerCreds))
		r.Get("/controller", s.controllerIndex())
		r.Get("/controller/", s.controllerIndex())
		r.Get("/controller/favicon.ico", s.controllerFavicon())
		r.Handle("/controller/static/*", http.FileServer(http.FS(assets.Content)))
	})

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
				r.Get("/status", s.playerStatus())
			})
		})
	})

	return r
}
