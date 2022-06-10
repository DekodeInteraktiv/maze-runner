package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/PeterBooker/maze-game-server/internal/game"
	"github.com/go-chi/chi"
	"golang.org/x/exp/slices"
)

// playerCreate sets up a new player in a specific game.
func (s *Server) playerCreate() http.HandlerFunc {
	type Payload struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		gamePass := chi.URLParam(r, "password")

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

		// Get the game ID and search for the game.
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(id)

		if gamePass != g.Password {
			data := struct {
				Error string
			}{
				"Invalid game password.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		if len(g.Players) >= 4 {
			data := struct {
				Error string
			}{
				"The game is full.",
			}

			writeJSON(w, data, http.StatusBadRequest)
			return
		}

		// Register the player in the game.
		p := g.RegisterPlayer(payload.Name, payload.Color)

		writeJSON(w, p, 200)
	}
}

// playerMoveOld moves a player.
func (s *Server) playerMoveOld() http.HandlerFunc {
	type Payload struct {
		Direction string `json:"direction"`
		Distance  int    `json:"distance"`
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

		// Get the game ID and search for the game.
		gameIDStr := chi.URLParam(r, "gameID")

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(gameID)

		// Get Auth Token.
		ctx := r.Context()
		token := ctx.Value("Token").(string)

		// Get the player by auth token.
		p := g.GetPlayerByToken(token)
		if p == nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check game is active.
		if g.Status == game.GameRunning {
			data := struct {
				Error string
			}{
				"The game is not active.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Calculate the new position.
		var newPos game.Point

		switch payload.Direction {
		case "north":
			newPos = p.Pos.North()
		case "south":
			newPos = p.Pos.South()
		case "west":
			newPos = p.Pos.West()
		case "east":
			newPos = p.Pos.East()
		}

		// Check if another player is already in the new position.
		for _, player := range g.Players {
			if player.ID != p.ID && newPos == *player.Pos {
				data := struct {
					Error string
				}{
					"Another player is already at this location.",
				}

				writeJSON(w, data, http.StatusConflict)
				return
			}
		}

		// Check if the player is trying to move outside the maze.
		if newPos.X < 0 || newPos.X > (g.Size-1) || newPos.Y < 0 || newPos.Y > (g.Size-1) {
			data := struct {
				Error string
			}{
				"Move position out of maze bounds.",
			}

			writeJSON(w, data, http.StatusConflict)
			return
		}

		// Check if the player is trying to move into a wall.
		if g.Maze[newPos.X][newPos.Y] == game.Wall {
			data := struct {
				Error string
			}{
				"Cannot move into a wall.",
			}

			writeJSON(w, data, http.StatusConflict)
			return
		}

		// Move player.
		g.MovePlayer(p, newPos)

		data := struct {
			Player *game.Player
		}{
			p,
		}

		writeJSON(w, data, 200)
	}
}

// playerMove queues a player move.
func (s *Server) playerMove() http.HandlerFunc {
	type Payload struct {
		Direction string `json:"direction"`
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

		// Get the game ID and search for the game.
		gameIDStr := chi.URLParam(r, "gameID")

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(gameID)

		// Get Auth Token.
		ctx := r.Context()
		token := ctx.Value("Token").(string)

		// Get the player by auth token.
		p := g.GetPlayerByToken(token)
		if p == nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check game is running.
		if g.Status != game.GameRunning {
			data := struct {
				Error string
			}{
				"The game is not running.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check if direction is valid.
		if !slices.Contains(game.Directions, payload.Direction) {
			data := struct {
				Error string
			}{
				"Invalid direction.",
			}

			writeJSON(w, data, http.StatusBadRequest)
			return
		}

		p.Lock()
		p.NextMove = payload.Direction
		p.Unlock()

		data := struct {
			Message string
		}{
			"Successfully queued next move.",
		}

		writeJSON(w, data, 200)
	}
}

// playerAbilityBomb uses the bomb ability.
func (s *Server) playerAbilityBomb() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the game ID and search for the game.
		gameIDStr := chi.URLParam(r, "gameID")

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(gameID)

		// Get Auth Token.
		ctx := r.Context()
		token := ctx.Value("Token").(string)

		// Get the player by auth token.
		p := g.GetPlayerByToken(token)
		if p == nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check game is active.
		if g.Status != game.GameRunning {
			data := struct {
				Error string
			}{
				"The game is not active.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check if player ability on cooldown.
		if !p.Abilities.BombAvailable {
			data := struct {
				Error string
			}{
				"Bomb ability on cooldown.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Calculate position.
		pos := &game.Point{
			X: p.Pos.X,
			Y: p.Pos.Y,
		}

		// Set ability on cooldown.
		p.Lock()
		p.Abilities.BombAvailable = false
		p.Unlock()

		// Make new object and action log.
		g.NewObject(game.Bomb, "", pos, p)
		g.NewAction(game.BombPlace, pos)

		data := struct {
			Message string
		}{
			"Successfully placed bomb.",
		}

		writeJSON(w, data, 200)
	}
}

// playerAbilityShoot uses the shoot ability.
func (s *Server) playerAbilityShoot() http.HandlerFunc {
	type Payload struct {
		Direction string `json:"direction"`
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

		// Get the game ID and search for the game.
		gameIDStr := chi.URLParam(r, "gameID")

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(gameID)

		// Get Auth Token.
		ctx := r.Context()
		token := ctx.Value("Token").(string)

		// Get the player by auth token.
		p := g.GetPlayerByToken(token)
		if p == nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check game is active.
		if g.Status != game.GameRunning {
			data := struct {
				Error string
			}{
				"The game is not active.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Check if player ability on cooldown.
		if !p.Abilities.ShootAvailable {
			data := struct {
				Error string
			}{
				"Shoot ability on cooldown.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Calculate position.
		pos := &game.Point{
			X: p.Pos.X,
			Y: p.Pos.Y,
		}

		// Set ability on cooldown.
		p.Lock()
		p.Abilities.ShootAvailable = false
		p.Unlock()

		// Make new object and action log.
		g.NewObject(game.Bullet, payload.Direction, pos, p)
		g.NewAction(game.Shoot, pos)

		data := struct {
			Message string
		}{
			"Successful shoot action.",
		}

		writeJSON(w, data, 200)
	}
}

// playerStatus gives the status of a player.
func (s *Server) playerStatus() http.HandlerFunc {
	type PlayerStatusResponse struct {
		Name         string
		ID           int
		Color        string
		Pos          *game.Point
		Team         game.ClaimType
		Maze         [][]game.MazeTileType `json:"maze"`
		Claims       [][]game.ClaimType    `json:"claims"`
		sync.RWMutex `json:"-"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Get the game ID and search for the game.
		gameIDStr := chi.URLParam(r, "gameID")

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		g := s.GetGameByID(gameID)

		// Get Auth Token.
		ctx := r.Context()
		token := ctx.Value("Token").(string)

		// Get the player by auth token.
		p := g.GetPlayerByToken(token)
		if p == nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusForbidden)
			return
		}

		// Player viewport size.
		size := 5

		// Find local maze.
		maze := make([][]game.MazeTileType, size)
		for i := range maze {
			maze[i] = make([]game.MazeTileType, size)
		}

		for x := (p.Pos.X - 2); x < (p.Pos.X + 2); x++ {
			for y := (p.Pos.Y - 2); y < (p.Pos.Y + 2); y++ {
				if x >= 0 && x < (g.Size-1) && y >= 0 && y < (g.Size-1) {
					maze[x][y] = g.Maze[x][y]
				}
			}
		}

		// Find local claims.
		claims := make([][]game.ClaimType, size)
		for i := range claims {
			claims[i] = make([]game.ClaimType, size)
		}

		for x := (p.Pos.X - 2); x < (p.Pos.X + 2); x++ {
			for y := (p.Pos.Y - 2); y < (p.Pos.Y + 2); y++ {
				if x >= 0 && x < (g.Size-1) && y >= 0 && y < (g.Size-1) {
					claims[x][y] = g.Claims[x][y]
				}
			}
		}

		resp := &PlayerStatusResponse{
			Name:   p.Name,
			ID:     p.ID,
			Pos:    p.Pos,
			Team:   p.Team,
			Maze:   maze,
			Claims: claims,
		}

		writeJSON(w, resp, 200)
	}
}
