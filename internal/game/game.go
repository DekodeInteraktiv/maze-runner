package game

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	noise "github.com/ojrac/opensimplex-go"
)

const (
	GameOpen     = "open"
	GameRunning  = "running"
	GameFinished = "finished"
)

var (
	gameID     = incrementer{id: 1}
	playerID   = incrementer{id: 1}
	actionID   = incrementer{id: 1}
	objectID   = incrementer{id: 1}
	Directions = []string{"north", "south", "west", "east"}
)

type Game struct {
	ID           int              `json:"id"`
	Password     string           `json:"password"`
	Token        string           `json:"-"`
	Status       string           `json:"status"`
	Timer        uint             `json:"timer"`
	TimeLimit    uint             `json:"time_limit"`
	Active       chan bool        `json:"-"`
	Players      []*Player        `json:"players"`
	Size         int              `json:"size"`
	Maze         [][]MazeTileType `json:"maze"`
	Claims       [][]ClaimType    `json:"claims"`
	Objects      []*Object        `json:"objects"`
	ActionLog    []*Action        `json:"-"`
	sync.RWMutex `json:"-"`
}

// TODO: Maze as [][]uint8 does not work as []uint8 is []byte and gets base64 encoded.
// Change to Maze *Maze with a Maze struct, with its own JSON encode method perhaps?

// New returns a new game instance.
func New(size int, distribution float64, timelimit uint) *Game {
	simplex := noise.New(rand.Int63())

	grid := newGrid(size)

	claims := newClaims(size)

	objects := make([]*Object, 0, 50)

	actionLog := make([]*Action, 0, 50)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			var tile MazeTileType
			tile = Floor
			v := simplex.Eval2(float64(x), float64(y))
			if v < distribution {
				tile = Wall
			}
			grid[x][y] = tile
			claims[x][y] = Unclaimed
		}
	}

	max := size - 1

	grid[0][0] = Floor
	grid[1][0] = Floor
	grid[0][1] = Floor
	grid[1][1] = Floor

	grid[max][max] = Floor
	grid[max-1][max] = Floor
	grid[max][max-1] = Floor
	grid[max-1][max-1] = Floor

	grid[0][max] = Floor
	grid[1][max] = Floor
	grid[0][max-1] = Floor
	grid[1][max-1] = Floor

	grid[max][0] = Floor
	grid[max][1] = Floor
	grid[max-1][0] = Floor
	grid[max-1][1] = Floor

	return &Game{
		ID:        gameID.new(),
		Password:  generatePassword(),
		Token:     strings.Replace(uuid.New().String(), "-", "", -1),
		TimeLimit: timelimit,
		Active:    make(chan bool),
		Status:    GameOpen,
		Size:      size,
		Maze:      grid,
		Claims:    claims,
		Objects:   objects,
		ActionLog: actionLog,
	}
}

// Reset resets the game.
func (g *Game) Reset() {
	g.Lock()
	defer g.Unlock()

	g.Claims = newClaims(g.Size)

	g.Objects = make([]*Object, 0, 50)

	g.ActionLog = make([]*Action, 0, 50)

	g.Active = make(chan bool)
	g.Timer = 0
	g.Status = GameOpen

	for _, p := range g.Players {
		p.Pos = g.PlayerStartPosition(p.Team)
	}
}

// Start starts the game.
func (g *Game) Start() {
	// Start the game after 5 seconds.
	duration := time.Duration(5) * time.Second
	time.AfterFunc(duration, g.runGame)
	time.AfterFunc(duration, g.runMovement)
}

// Stop stops the game.
func (g *Game) Stop() {
	g.Lock()
	defer g.Unlock()

	<-g.Active
	g.Status = GameFinished
}

// runGame runs the game timer.
func (g *Game) runGame() {
	g.Lock()
	g.Status = GameRunning
	g.Unlock()

	g.NewAction(GameStart, nil)

	ticker := time.NewTicker(1 * time.Second)

	go func(g *Game) {
		for {
			select {
			case <-g.Active:
				return
			case <-ticker.C:
				g.Lock()
				g.Timer++
				g.Unlock()

				if g.Timer >= (g.TimeLimit - 15) {
					g.NewAction(GameEndWarning, nil)
				}

				if g.Timer >= g.TimeLimit {
					ticker.Stop()

					g.Lock()
					g.Status = GameFinished
					g.Unlock()

					g.NewAction(GameEnd, nil)

					g.Active <- true
				}
			}
		}
	}(g)
}

// runMovement runs the player movement action.
func (g *Game) runMovement() {
	ticker := time.NewTicker(350 * time.Millisecond)

	go func(g *Game) {
		for {
			select {
			case <-g.Active:
				ticker.Stop()
				return
			case <-ticker.C:
				for _, p := range g.Players {
					if p != nil {
						// Skip if no move queued.
						if p.NextMove == "" {
							continue
						}

						p.Lock()
						direction := p.NextMove
						p.NextMove = ""

						// Calculate the new position.
						var newPos Point

						switch direction {
						case "north":
							newPos = p.Pos.North()
						case "south":
							newPos = p.Pos.South()
						case "west":
							newPos = p.Pos.West()
						case "east":
							newPos = p.Pos.East()
						}
						p.Unlock()

						g.RLock()
						// Check if the player is trying to move outside the maze.
						if newPos.X < 0 || newPos.X > (g.Size-1) || newPos.Y < 0 || newPos.Y > (g.Size-1) {
							g.RUnlock()
							continue
						}

						// Check if the player is trying to move into a wall.
						if g.Maze[newPos.X][newPos.Y] == Wall {
							g.RUnlock()
							continue
						}
						g.RUnlock()

						// Move player.
						g.MovePlayer(p, newPos)
					}
				}
			}
		}
	}(g)
}

// RegisterPlayer registers a player.
func (g *Game) RegisterPlayer(name, styles string) *Player {
	g.Lock()
	defer g.Unlock()

	var pos *Point
	var team ClaimType

	switch len(g.Players) {
	case 0:
		team = Red
		pos = g.PlayerStartPosition(team)
	case 1:
		team = Blue
		pos = g.PlayerStartPosition(team)
	case 2:
		team = Green
		pos = g.PlayerStartPosition(team)
	case 3:
		team = Yellow
		pos = g.PlayerStartPosition(team)
	}

	p := &Player{
		ID:       playerID.new(),
		Name:     name,
		Styles:   styles,
		Team:     team,
		Token:    strings.Replace(uuid.New().String(), "-", "", -1),
		Pos:      pos,
		NextMove: "",
	}

	p.Abilities.BombAvailable = true
	p.Abilities.ShootAvailable = true

	g.Players = append(g.Players, p)

	return p
}

func (g *Game) PlayerStartPosition(team ClaimType) *Point {
	g.RLock()
	defer g.RUnlock()

	switch team {
	case Red:
		return &Point{X: 0, Y: 0}
	case Blue:
		return &Point{X: g.Size - 1, Y: g.Size - 1}
	case Green:
		return &Point{X: 0, Y: g.Size - 1}
	case Yellow:
		return &Point{X: g.Size - 1, Y: 0}
	}

	return nil
}

// GetPlayerByToken finds a player by their token.
func (g *Game) GetPlayerByToken(token string) *Player {
	g.Lock()
	defer g.Unlock()

	for _, p := range g.Players {
		if p.Token == token {
			return p
		}
	}

	return nil
}

func (g *Game) MovePlayer(p *Player, newPos Point) {
	g.Lock()
	defer g.Unlock()

	// Move to new position.
	p.Pos.X = newPos.X
	p.Pos.Y = newPos.Y

	// Claim tile.
	g.Claims[p.Pos.X][p.Pos.Y] = p.Team
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
