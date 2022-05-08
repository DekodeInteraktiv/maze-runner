package server

import (
	"log"
	"net/http"
	"time"

	"github.com/PeterBooker/maze-game-server/internal/assets"
	"github.com/PeterBooker/maze-game-server/internal/client"
	"github.com/PeterBooker/maze-game-server/internal/config"
	"github.com/PeterBooker/maze-game-server/internal/game"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	Router *chi.Mux
	Client *http.Client
	Game   *game.Game
}

// New returns a pointer to the main server struct
func New(l *log.Logger, c *config.Config) *Server {
	s := &Server{
		Config: c,
		Logger: l,
		Client: client.New(),
		Game:   game.New(),
	}

	return s
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	s.Router = chi.NewRouter()

	// Middleware Stack
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.RedirectSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.Router.Use(middleware.Timeout(15 * time.Second))

	s.Router.Use(s.VerifyToken())

	s.Router.Handle("/static/*", http.FileServer(http.FS(assets.Content)))

	s.routes()

	s.startHTTP()
}
