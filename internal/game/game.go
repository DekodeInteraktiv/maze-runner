package game

import (
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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
	ID           int                       `json:"id"`
	Password     string                    `json:"password"`
	Token        string                    `json:"-"`
	Active       bool                      `json:"active"`
	Timer        uint                      `json:"timer"`
	Players      []*Player                 `json:"players"`
	Maze         [][]uint8                 `json:"-"`
	Maze2        map[uint8]map[uint8]uint8 `json:"-"`
	sync.RWMutex `json:"-"`
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

	m2 := make(map[uint8]map[uint8]uint8, 10)

	m2[1] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[2] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[3] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[4] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[5] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[6] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[7] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[8] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[9] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	m2[10] = map[uint8]uint8{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}

	m2[1][7] = 1

	return &Game{
		ID:       gameID.new(),
		Password: generatePassword(),
		Maze:     m,
		Maze2:    m2,
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

	p := Player{
		ID:    playerID.new(),
		Name:  name,
		Color: color,
		Token: strings.Replace(uuid.New().String(), "-", "", -1),
		Pos:   &Point{X: 0, Y: 0},
	}

	g.Players = append(g.Players, &p)

	return &p
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
