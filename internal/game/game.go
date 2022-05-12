package game

import (
	"sync"
)

var id = incrementer{
	id: 1,
}

type Game struct {
	ID            int       `json:"id"`
	Password      string    `json:"password"`
	Maze          [][]uint8 `json:"maze"`
	Players       []*Player `json:"players"`
	*sync.RWMutex `json:"-"`
}

// New returns a new game instance.
func New() *Game {
	//m := maze.Generate(20, 20, time.Now().UnixNano())

	//m := make([][]uint8, 20)
	//for i := range m {
	//m[i] = make([]uint8, 20)
	//}

	m := make([][]uint8, 10)
	m[0] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[1] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[2] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[3] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[4] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[5] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[6] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[7] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[8] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	m[9] = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return &Game{
		ID:       id.new(),
		Password: generatePassword(),
		Maze:     m,
	}
}

func (g *Game) Start() {

}

// RegisterPlayer registers a player.
func (g *Game) RegisterPlayer(name string) *Player {
	g.Lock()
	defer g.Unlock()

	p := Player{
		ID:    "1",
		Name:  name,
		Token: "AAAA",
	}

	g.Players = append(g.Players, &p)

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
