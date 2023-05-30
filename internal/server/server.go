package server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/DekodeInteraktiv/maze-runner/internal/assets"
	"github.com/DekodeInteraktiv/maze-runner/internal/client"
	"github.com/DekodeInteraktiv/maze-runner/internal/config"
	"github.com/DekodeInteraktiv/maze-runner/internal/game"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	Router *chi.Mux
	Client *http.Client
	Games  []*game.Game
	sync.RWMutex
}

// New returns a pointer to the main server struct
func New(l *log.Logger, c *config.Config) *Server {
	s := &Server{
		Config: c,
		Logger: l,
		Client: client.New(),
		Games:  make([]*game.Game, 0, 50),
	}

	return s
}

func (s *Server) CreateGame(size int, distribution float64, timelimit uint, protected bool) *game.Game {
	g := game.New(size, distribution, timelimit, protected)

	s.Lock()
	defer s.Unlock()

	s.Games = append(s.Games, g)

	return g
}

func (s *Server) GetGameByID(id int) *game.Game {
	s.Lock()
	defer s.Unlock()

	if id > len(s.Games) {
		return nil
	}

	for _, g := range s.Games {
		if g.ID == id {
			return g
		}
	}

	return nil
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
	s.Router.Use(middleware.Timeout(5 * time.Second))
	s.Router.Use(s.VerifyToken())

	s.Router.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.Router.Handle("/static/*", http.FileServer(http.FS(assets.Content)))

	s.routes()

	if s.Config.Env == "local" {
		s.startHTTP()
	} else {
		s.startHTTPS()
	}
}
