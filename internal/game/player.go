package game

import (
	"sync"
)

type Player struct {
	Name      string    `json:"name"`
	ID        int       `json:"id"`
	Styles    string    `json:"styles"`
	Pos       *Point    `json:"pos"`
	Token     string    `json:"token"`
	Team      ClaimType `json:"team"`
	Abilities Abilities `json:"abilities"`
	NextMove  string    `json:"-"`
	sync.RWMutex
}

type Abilities struct {
	BombAvailable  bool `json:"bomb_available"`
	ShootAvailable bool `json:"shoot_available"`
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

func (p *Point) North() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p *Point) South() Point {
	return Point{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p *Point) West() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p *Point) East() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y,
	}
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
