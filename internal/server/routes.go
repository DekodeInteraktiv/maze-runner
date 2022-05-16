package server

import (
	"github.com/go-chi/chi"
)

func (s *Server) routes() {
	// Add Routes
	s.Router.Get("/", s.webIndex())
	s.Router.Get("/favicon.ico", s.webFavicon())
	s.Router.Get("/docs", s.webDocs())

	// Add API v1 routes
	s.Router.Mount("/api/v1", s.apiRoutes())

	// Handle NotFound
	//s.Router.NotFound(s.notFound())
}

func (s *Server) apiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/game", func(r chi.Router) {
		r.Get("/create", s.gameCreate())

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
