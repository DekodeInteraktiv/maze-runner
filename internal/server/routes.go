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

	//r.Post("/register/{name}", s.register())

	//r.Get("/player/get", s.getPlayer())

	//r.Get("/get/maze", s.imageExample())

	r.Route("/game", func(r chi.Router) {
		r.Get("/create", s.gameCreate())
		r.Get("/info/{id}", s.gameInfo())
	})

	return r
}
