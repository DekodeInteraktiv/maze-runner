package game

import "time"

type ObjectType uint16

const (
	None ObjectType = iota
	Bullet
	Bomb
)

type Object struct {
	ID        int                              `json:"id"`
	Type      ObjectType                       `json:"type"`
	Direction string                           `json:"direction"`
	Pos       *Point                           `json:"pos"`
	Owner     *Player                          `json:"owner"`
	Func      func(game *Game, object *Object) `json:"-"`
}

// NewObject returns a new object.
func (g *Game) NewObject(objectType ObjectType, direction string, pos *Point, p *Player) {
	g.Lock()
	defer g.Unlock()

	object := &Object{
		ID:        0,
		Type:      objectType,
		Direction: direction,
		Pos:       pos,
		Owner:     p,
	}

	ticker := time.NewTicker(1 * time.Second)
	end := time.Now().Add(5 * time.Second)

	go func(g *Game, o *Object, end time.Time) {
		for {
			select {
			case <-g.Active:
				return
			case <-ticker.C:
				if time.Now().After(end) {
					ticker.Stop()

					g.Lock()
					// Bomb explodes and paints all tiles in 1 tile range (9 total).
					for x := (pos.X - 2); x < (pos.X + 2); x++ {
						for y := (pos.Y - 2); y < (pos.Y + 2); y++ {
							if x >= 0 && x < (g.Size-1) && y >= 0 && y < (g.Size-1) {
								g.Claims[x][y] = object.Owner.Team
							}
						}
					}
					g.Unlock()

					object.Owner.Lock()
					object.Owner.Abilities.BombAvailable = true
					object.Owner.Unlock()
				}
			}
		}
	}(g, object, end)

	g.Objects = append(g.Objects, object)
}
