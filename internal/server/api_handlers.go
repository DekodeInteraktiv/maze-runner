package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/PeterBooker/maze-game-server/internal/game"
	"github.com/go-chi/chi"
)

// error holds data about an error
type error struct {
	Err string `json:"error"`
}

// gameCreate sets up a new game.
func (s *Server) gameCreate() http.HandlerFunc {
	type CreateGameResponse struct {
		ID           int            `json:"id"`
		Password     string         `json:"password"`
		Token        string         `json:"token"`
		Active       bool           `json:"active"`
		Timer        uint           `json:"timer"`
		Players      []*game.Player `json:"players"`
		Maze         [][]uint8      `json:"maze"`
		sync.RWMutex `json:"-"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		g := s.CreateGame()

		g.RLock()
		defer g.RUnlock()
		newGame := (*CreateGameResponse)(g)

		writeJSON(w, newGame, 200)
	}
}

// gameInfo gets info for a specific game.
func (s *Server) gameInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "gameID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, 404)
		}

		g := s.GetGameByID(id)

		g.RLock()
		defer g.RUnlock()

		writeJSON(w, g, 200)
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

			writeJSON(w, data, 404)
		}

		g := s.GetGameByID(id)

		g.Start()

		g.RLock()
		data := struct {
			Message string
		}{
			fmt.Sprintf("Game (ID: %d) is starting in 5 seconds.", g.ID),
		}
		g.RUnlock()

		writeJSON(w, data, 200)
	}
}

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

		// Register the player in the game.
		p := g.RegisterPlayer(payload.Name, payload.Color)

		data := struct {
			Player *game.Player
		}{
			p,
		}

		writeJSON(w, data, 200)
	}
}

// playerMove moves a player.
func (s *Server) playerMove() http.HandlerFunc {
	type Payload struct {
		Direction string `json:"direction"`
		Spaces    uint   `json:"spaces"`
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

		// Register the player in the game.
		p, err := g.GetPlayerByToken(token)
		if err != nil {
			data := struct {
				Error string
			}{
				"Authentication token does not match any player registered for this game.",
			}

			writeJSON(w, data, http.StatusNotFound)
			return
		}

		data := struct {
			Player *game.Player
		}{
			p,
		}

		writeJSON(w, data, 200)
	}
}

/*
{
    "id": 2,
    "password": "distracted",
    "token": "7db66e97a5b64aaba34cd1fcdc79c0f5",
    "active": false,
    "timer": 0,
    "players": null
}

{
    "Player": {
        "Name": "Cath",
        "ID": 1,
        "Sprite": null,
        "Color": "#eeeeee",
        "Pos": {
            "X": 0,
            "Y": 0
        },
        "Token": "58eefe57ee7f424f92c2cf69dd76e8d7"
    }
}

*/

// game token: 6cb936ddd08d4c53aed39fb8e1b940f2
// player token:

// register ...
/*func (s *Server) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		ctx := r.Context()
		token, _ := ctx.Value("Token").(string)

		fmt.Println("Token: " + token)

		if name == "" {
			data := error{
				Err: "Name is required",
			}

			writeJSON(w, data, 404)
			return
		}

		s.Game.RegisterPlayer(name)

		data := struct {
			Message string
			ID      string
		}{
			"Hello " + name + ", you are successfully registered.",
			"193473464793",
		}

		writeJSON(w, data, 200)
	}
}*/

// playerMove ...
/*func (s *Server) playerMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "direction")

		if name == "" {
			data := error{
				Err: "Name is required",
			}

			writeJSON(w, data, 404)
			return
		}

		s.Game.Lock()

		ctx := r.Context()
		token, _ := ctx.Value("Token").(string)

		p, err := s.Game.GetPlayerByToken(token)
		if err != nil {

		}

		s.Game.Unlock()

		data := struct {
			Message string
			ID      string
		}{
			"Hello " + p.Name + ", you have successfully moved.",
			"193473464793",
		}

		writeJSON(w, data, 200)
	}
}*/

// getPlayer ...
/*func (s *Server) getPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token, _ := ctx.Value("Token").(string)

		fmt.Println("Token: " + token)

		p, err := s.Game.GetPlayerByToken(token)
		if err != nil {

		}

		data := struct {
			Player *game.Player
		}{
			p,
		}

		writeJSON(w, data, 200)
	}
}*/

// apiExample ...
func (s *Server) apiExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		if name == "" {
			data := error{
				Err: "Name is required",
			}

			writeJSON(w, data, 404)
			return
		}

		data := struct {
			Message string
			ID      string
		}{
			"Hello " + name + ", you are successfully registered.",
			"193473464793",
		}

		writeJSON(w, data, 200)
	}
}

// apiExample ...
/*func (s *Server) imageExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		img := s.Game.GetImage()

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			log.Panicf("Failed to encode image: %v\n", err)
		}

		w.Write(buf.Bytes())
	}
}*/

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panicf("Failed to encode JSON: %v\n", err)
	}
}
