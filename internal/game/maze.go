package game

import (
	"math/rand"
	"sync"

	noise "github.com/ojrac/opensimplex-go"
)

type Maze struct {
	Grid         [][]MazeTileType `json:"grid"`
	Claims       [][]ClaimType    `json:"claims"`
	sync.RWMutex `json:"-"`
}

type MazeTileType uint16

const (
	Floor MazeTileType = iota
	Wall
	Portal
)

type ClaimType uint16

const (
	Unclaimed ClaimType = iota
	Red
	Blue
	Green
	Yellow
)

func NewMaze(size int) *Maze {
	simplex := noise.New(rand.Int63())

	grid := make([][]MazeTileType, size)
	for i := range grid {
		grid[i] = make([]MazeTileType, size)
	}

	claims := make([][]ClaimType, size)
	for i := range claims {
		claims[i] = make([]ClaimType, size)
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
			claims[x][y] = Unclaimed
		}
	}

	return &Maze{
		Grid:   grid,
		Claims: claims,
	}
}
