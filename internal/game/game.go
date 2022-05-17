package game

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	noise "github.com/ojrac/opensimplex-go"
)

var (
	gameID = incrementer{
		id: 1,
	}
	playerID = incrementer{
		id: 1,
	}
)

type Game struct {
	ID           int              `json:"id"`
	Password     string           `json:"password"`
	Token        string           `json:"-"`
	Active       bool             `json:"active"`
	Timer        uint             `json:"timer"`
	Players      []*Player        `json:"players"`
	Maze         [][]MazeTileType `json:"maze"`
	sync.RWMutex `json:"-"`
}

// TODO: Maze as [][]uint8 does not work as []uint8 is []byte and gets base64 encoded.
// Change to Maze *Maze with a Maze struct, with its own JSON encode method perhaps?

// New returns a new game instance.
func New(size int) *Game {
	//maze := NewMaze(size)

	simplex := noise.New(rand.Int63())

	grid := make([][]MazeTileType, size)
	for i := range grid {
		grid[i] = make([]MazeTileType, size)
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			var tile MazeTileType
			tile = Floor
			v := simplex.Eval2(float64(x), float64(y))
			if v < -0.35 {
				tile = Wall
			}
			grid[x][y] = tile
		}
	}

	return &Game{
		ID:       gameID.new(),
		Password: generatePassword(),
		Token:    strings.Replace(uuid.New().String(), "-", "", -1),
		Maze:     grid,
		Active:   false,
	}
}

// Start triggers
func (g *Game) Start() {
	// Start the game after 5 seconds.
	duration := time.Duration(5) * time.Second
	time.AfterFunc(duration, g.setActive)
	time.AfterFunc(duration, g.runGame)
}

// start starts the game.
func (g *Game) setActive() {
	g.Lock()
	g.Active = true
	g.Unlock()
}

// runGame runs the game timer.
func (g *Game) runGame() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func(g *Game) {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				g.Lock()
				g.Timer++
				g.Unlock()
				// If 5 minutes (300 seconds) has passed, end the game.
				if g.Timer >= 30 {
					ticker.Stop()

					g.Lock()
					g.Active = false
					g.Unlock()

					done <- true
				}
			}
		}
	}(g)
}

// RegisterPlayer registers a player.
func (g *Game) RegisterPlayer(name, color string) *Player {
	g.Lock()
	defer g.Unlock()

	var pos *Point

	switch len(g.Players) {
	case 0:
		pos = &Point{X: 0, Y: 0}
	case 1:
		pos = &Point{X: 49, Y: 49}
	case 2:
		pos = &Point{X: 0, Y: 49}
	case 3:
		pos = &Point{X: 49, Y: 0}
	}

	p := &Player{
		ID:    playerID.new(),
		Name:  name,
		Color: color,
		Token: strings.Replace(uuid.New().String(), "-", "", -1),
		Pos:   pos,
	}

	g.Players = append(g.Players, p)

	return p
}

// GetPlayerByToken finds a player by their token.
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
