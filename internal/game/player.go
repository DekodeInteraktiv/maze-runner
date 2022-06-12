package game

import (
	"sync"
	"time"
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

// ShootCooldown manages the shoot cooldown.
func (p *Player) ShootCooldown() {
	p.Lock()
	p.Abilities.ShootAvailable = false
	p.Unlock()

	ticker := time.NewTicker(1 * time.Second)
	end := time.Now().Add(5 * time.Second)

	go func(p *Player) {
		for {
			select {
			case <-ticker.C:
				if time.Now().After(end) {
					ticker.Stop()

					p.Lock()
					p.Abilities.ShootAvailable = true
					p.Unlock()
				}
			}
		}
	}(p)
}

// BombCooldown manages the bomb cooldown.
func (p *Player) BombCooldown() {
	p.Lock()
	p.Abilities.BombAvailable = false
	p.Unlock()

	ticker := time.NewTicker(1 * time.Second)
	end := time.Now().Add(10 * time.Second)

	go func(p *Player) {
		for {
			select {
			case <-ticker.C:
				if time.Now().After(end) {
					ticker.Stop()

					p.Lock()
					p.Abilities.BombAvailable = true
					p.Unlock()
				}
			}
		}
	}(p)
}
