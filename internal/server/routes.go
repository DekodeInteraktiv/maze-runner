package server

import (
	"net/http"
	"time"

	"github.com/PeterBooker/maze-game-server/internal/assets"
	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
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
		r.Get("/viewer/robots.txt", s.viewerRobots())
		r.Get("/viewer/manifest.json", s.viewerManifest())
		r.Get("/viewer/asset-manifest.json", s.viewerAssetManifest())
		r.Get("/viewer/logo192.png", s.viewerLogo192())
		r.Get("/viewer/logo512.png", s.viewerLogo512())
		r.Handle("/viewer/static/*", http.FileServer(http.FS(assets.Content)))
	})

	// Controller App
	controllerCreds := map[string]string{"controller": "happyhour"}
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("controller", controllerCreds))
		r.Get("/controller", s.controllerIndex())
		r.Get("/controller/", s.controllerIndex())
		r.Get("/controller/favicon.ico", s.controllerFavicon())
		r.Get("/controller/robots.txt", s.controllerRobots())
		r.Get("/controller/manifest.json", s.controllerManifest())
		r.Get("/controller/asset-manifest.json", s.controllerAssetManifest())
		r.Get("/controller/logo192.png", s.controllerLogo192())
		r.Get("/controller/logo512.png", s.controllerLogo512())
		r.Handle("/controller/static/*", http.FileServer(http.FS(assets.Content)))
	})

	// Leaderboard App
	leaderboardCreds := map[string]string{"leaderboard": "d3kode"}
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("leaderboard", leaderboardCreds))
		r.Get("/leaderboard", s.leaderboardIndex())
		r.Get("/leaderboard/", s.leaderboardIndex())
		r.Get("/leaderboard/favicon.ico", s.leaderboardFavicon())
		r.Get("/leaderboard/robots.txt", s.leaderboardRobots())
		r.Get("/leaderboard/manifest.json", s.leaderboardManifest())
		r.Get("/leaderboard/asset-manifest.json", s.leaderboardAssetManifest())
		r.Get("/leaderboard/logo192.png", s.leaderboardLogo192())
		r.Get("/leaderboard/logo512.png", s.leaderboardLogo512())
		r.Handle("/leaderboard/static/*", http.FileServer(http.FS(assets.Content)))
	})

	// Register App
	registerCreds := map[string]string{"register": "d3kode"}
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("register", registerCreds))
		r.Get("/register", s.registerIndex())
		r.Get("/register/", s.registerIndex())
		r.Get("/register/favicon.ico", s.registerFavicon())
		r.Get("/register/robots.txt", s.registerRobots())
		r.Get("/register/manifest.json", s.registerManifest())
		r.Get("/register/asset-manifest.json", s.registerAssetManifest())
		r.Get("/register/logo192.png", s.registerLogo192())
		r.Get("/register/logo512.png", s.registerLogo512())
		r.Handle("/register/static/*", http.FileServer(http.FS(assets.Content)))
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
		r.Get("/changelog", s.gameChangelog())
		r.Get("/manual/{password}", s.gameManual())
		r.Get("/statistics", s.gameStatistics())

		r.Route("/{gameID}", func(r chi.Router) {
			r.Get("/info", s.gameStatus())
			r.Get("/status", s.gameStatus())
			r.Get("/start", s.gameStart())
			r.Get("/reset", s.gameReset())
			r.Get("/stop", s.gameStop())
			r.Get("/tile/gift", s.gameTileGift())

			r.Route("/player", func(r chi.Router) {
				limiter := tollbooth.NewLimiter(10, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute * 5})
				r.Use(RateLimiter(limiter))

				r.Post("/register/{password}", s.playerCreate())

				r.Post("/move", s.playerMove())
				r.Get("/status", s.playerStatus())

				r.Get("/ability/bomb", s.playerAbilityBomb())
				r.Post("/ability/shoot", s.playerAbilityShoot())
			})
		})
	})

	return r
}
