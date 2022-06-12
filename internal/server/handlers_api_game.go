package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/PeterBooker/maze-game-server/internal/game"
	"github.com/go-chi/chi"
)

// gameCreate sets up a new game.
func (s *Server) gameCreate() http.HandlerFunc {
	type Payload struct {
		Size         int     `json:"size"`
		Distribution float64 `json:"distribution"`
		TimeLimit    uint    `json:"timelimit"`
	}

	type CreateGameResponse struct {
		ID           int                   `json:"id"`
		Password     string                `json:"password"`
		Token        string                `json:"token"`
		Status       string                `json:"status"`
		Timer        uint                  `json:"timer"`
		TimeLimit    uint                  `json:"time_limit"`
		Active       chan bool             `json:"-"`
		Players      []*game.Player        `json:"players"`
		Size         int                   `json:"size"`
		Maze         [][]game.MazeTileType `json:"maze"`
		Claims       [][]game.ClaimType    `json:"claims"`
		Objects      []*game.Object        `json:"objects"`
		ActionLog    []*game.Action        `json:"-"`
		sync.RWMutex `json:"-"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JSON encoded body data.
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid POST data.",
			}
			writeJSON(w, data, http.StatusBadRequest)
			return
		}

		if payload.Size == 0 || payload.Size < 5 || payload.Size > 200 {
			payload.Size = 20
		}

		if payload.Distribution == 0 || payload.Distribution < -0.8 || payload.Distribution > 0.8 {
			payload.Distribution = -0.35
		}

		if payload.TimeLimit == 0 || payload.TimeLimit < 15 || payload.TimeLimit > 900 {
			payload.TimeLimit = 60
		}

		g := s.CreateGame(payload.Size, payload.Distribution, payload.TimeLimit)

		g.RLock()
		defer g.RUnlock()
		newGame := (*CreateGameResponse)(g)

		writeJSON(w, newGame, http.StatusAccepted)
	}
}

// gameStatus gets the status for a specific game.
func (s *Server) gameStatus() http.HandlerFunc {
	type StatusGameResponse struct {
		ID           int                   `json:"id"`
		Password     string                `json:"password"`
		Token        string                `json:"-"`
		Status       string                `json:"status"`
		Timer        uint                  `json:"timer"`
		TimeLimit    uint                  `json:"time_limit"`
		Active       chan bool             `json:"-"`
		Players      []*game.Player        `json:"players"`
		Size         int                   `json:"size"`
		Maze         [][]game.MazeTileType `json:"maze"`
		Claims       [][]game.ClaimType    `json:"claims"`
		Objects      []*game.Object        `json:"objects"`
		ActionLog    []*game.Action        `json:"log"`
		sync.RWMutex `json:"-"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(id)
		if g == nil {
			data := struct {
				Error string
			}{
				fmt.Sprintf("No game found with ID: %d.", id),
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		if g == nil {
			data := struct {
				Error string
			}{
				"Game not found",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g.RLock()
		defer g.RUnlock()
		newGame := (*StatusGameResponse)(g)
		//newGame.ActionLog = newGame.ActionLog[0:25]

		writeJSON(w, newGame, http.StatusAccepted)
	}
}

// gameStart schedules the start of the game.
func (s *Server) gameStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(id)
		if g == nil {
			data := struct {
				Error string
			}{
				fmt.Sprintf("No game found with ID: %d.", id),
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g.Start()

		g.RLock()
		data := struct {
			Message string
		}{
			fmt.Sprintf("Game (ID: %d) is starting in 5 seconds.", g.ID),
		}
		g.RUnlock()

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameReset resets the game to its initial state.
func (s *Server) gameReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(id)
		if g == nil {
			data := struct {
				Error string
			}{
				fmt.Sprintf("No game found with ID: %d.", id),
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g.Reset()

		data := struct {
			Message string
		}{
			fmt.Sprintf("Game (ID: %d) has been reset.", id),
		}

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameStop finishes the game.
func (s *Server) gameStop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(id)
		if g == nil {
			data := struct {
				Error string
			}{
				fmt.Sprintf("No game found with ID: %d.", id),
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g.Stop()

		data := struct {
			Message string
		}{
			fmt.Sprintf("Game (ID: %d) has been stopped.", id),
		}

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameTileGift sets up a new game.
func (s *Server) gameTileGift() http.HandlerFunc {
	type Payload struct {
		X    int    `json:"x"`
		Y    int    `json:"y"`
		Team uint16 `json:"team"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JSON encoded body data.
		var payload Payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid POST data.",
			}
			writeJSON(w, data, http.StatusBadRequest)
			return
		}

		// Get the Game ID.
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		// Find the game.
		g := s.GetGameByID(id)
		if g == nil {
			data := struct {
				Error string
			}{
				fmt.Sprintf("No game found with ID: %d.", id),
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		// Check the X/Y cords are valid.
		if payload.X >= (g.Size-1) || payload.Y >= (g.Size-1) || payload.X < 0 || payload.Y < 0 {
			data := struct {
				Error string
			}{
				"Coordinates are out of bounds.",
			}

			writeJSON(w, data, http.StatusBadRequest)
			return
		}

		g.RLock()
		g.Claims[payload.X][payload.Y] = game.ClaimType(payload.Team)
		g.RUnlock()

		data := struct {
			Message string
		}{
			"Tile successfully gifted.",
		}

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameChangelog returns a URL to the changelog for the game.
func (s *Server) gameChangelog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Message string
		}{
			"https://docs.google.com/document/d/1ikyNSb1o0u0czTJ7WjmVH_arV51gtnG8uIOlOmdMIKA",
		}

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameManual returns a URL to the changelog for the game.
func (s *Server) gameManual() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := chi.URLParam(r, "password")

		if password != "Co70qCEl6t8J" {
			data := struct {
				Error string
			}{
				"Invalid password",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		data := struct {
			Message string
		}{
			"https://docs.google.com/presentation/d/1DkklzvDPjxUcVdR1-_bOtjbO0HeyUVgGgL7fWnBELtc",
		}

		writeJSON(w, data, http.StatusAccepted)
	}
}

// gameStatistics returns a URL to the changelog for the game.
func (s *Server) gameStatistics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Error string
		}{
			"Endpoint no longer supported due to GDPR concerns",
		}

		writeJSON(w, data, http.StatusForbidden)
	}
}
