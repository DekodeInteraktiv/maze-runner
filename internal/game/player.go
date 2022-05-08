package game

import (
	"image/color"
	"sync"
)

type Player struct {
	Name   string
	ID     string
	Sprite []byte
	Color  color.RGBA
	Pos    *Point
	Token  string
	sync.RWMutex
}

type Point struct {
	X, Y int
}

// Move moves a player.
func (p *Player) Move(direction string) *Point {
	p.Lock()
	defer p.Unlock()

	return p.Pos
}
