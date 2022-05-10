package game

import (
	"image"
	"sync"
	"time"

	"gitlab.com/zaba505/maze"
)

var id = incrementer{
	id: 1,
}

type Game struct {
	ID            int        `json:"id"`
	Maze          *maze.Maze `json:"-"`
	Players       []*Player  `json:"players"`
	*sync.RWMutex `json:"-"`
}

// New returns a new game instance.
func New() *Game {
	m := maze.Generate(20, 20, time.Now().UnixNano())

	return &Game{
		ID:   id.new(),
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
		if p.Token == token {
			return p, nil
		}
	}

	return nil, nil
}

type incrementer struct {
	sync.Mutex
	id int
}

func (a *incrementer) new() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
