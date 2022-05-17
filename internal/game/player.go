package game

import (
	"sync"
)

type Player struct {
	Name   string
	ID     int
	Sprite []byte
	Color  string
	Pos    *Point
	Token  string
	sync.RWMutex
}

// Move moves a player.
func (p *Player) Move(direction string) *Point {
	p.Lock()
	defer p.Unlock()

	return p.Pos
}

type Point struct {
	X, Y int
}

func (p *Point) MoveNorth() {
	p.Y++
}

func (p *Point) MoveSouth() {
	p.Y--
}

func (p *Point) MoveWest() {
	p.X--
}

func (p *Point) MoveEast() {
	p.X++
}
