package game

import (
	"fmt"
	"image"
	"sync"
	"time"

	"gitlab.com/zaba505/maze"
)

var seed = time.Now().UnixNano()

type Game struct {
	Name    string
	Maze    *maze.Maze
	Players []*Player
	sync.RWMutex
}

// New returns a new game instance.
func New() *Game {
	m := maze.Generate(20, 20, seed)

	//fmt.Println(m.MarshalText())

	return &Game{
		Name: "Maze Game",
		Maze: m,
	}
}

func (g *Game) GetImage() *image.Gray {
	return maze.Gray(g.Maze)
}

func (g *Game) Start() {

}

// RegisterPlayer registers a player.
func (g *Game) RegisterPlayer(name string) *Player {
	g.Lock()

	p := Player{
		ID:    "1",
		Name:  name,
		Token: "AAAA",
	}
	g.Players = append(g.Players, &p)

	g.Unlock()

	return &p
}

// RegisterPlayer registers a player.
func (g *Game) GetPlayerByToken(token string) (*Player, error) {
	g.Lock()
	defer g.Unlock()

	for _, p := range g.Players {
		fmt.Println(p)
		if p.Token == token {
			return p, nil
		}
	}

	return nil, nil
}
