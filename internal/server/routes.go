package server

import (
	"net/http"

	"github.com/PeterBooker/maze-game-server/internal/assets"
	"github.com/go-chi/chi"
)

func (s *Server) routes() {
	// Add Routes
	s.Router.Get("/", s.webIndex())
	s.Router.Get("/favicon.ico", s.webFavicon())

	// Viewer
	s.Router.Get("/viewer", s.viewerIndex())
	s.Router.Get("/viewer/favicon.ico", s.viewerFavicon())
	s.Router.Handle("/viewer/static/*", http.FileServer(http.FS(assets.Content)))

	// Controller
	s.Router.Get("/controller", s.controllerIndex())
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
			})
		})
	})

	return r
}
