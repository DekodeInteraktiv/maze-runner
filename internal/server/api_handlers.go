package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PeterBooker/maze-game-server/internal/game"
	"github.com/go-chi/chi"
)

// error holds data about an error
type error struct {
	Err string `json:"error"`
}

// gameCreate sets up a new game.
func (s *Server) gameCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g := game.New()

		data := struct {
			Game *game.Game
		}{
			g,
		}

		writeJSON(w, data, 200)
	}
}

// gameInfo gets info for a specific game.
func (s *Server) gameInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		fmt.Println("ID: " + idStr)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			data := struct {
				Error string
			}{
				"Invalid Game ID",
			}

			writeJSON(w, data, 404)
		}

		fmt.Println("Getting Game...")

		g := s.GetGameByID(id)

		//g.Lock()
		//defer g.Unlock()

		data := struct {
			Game *game.Game
		}{
			g,
		}

		writeJSON(w, data, 200)
	}
}

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
