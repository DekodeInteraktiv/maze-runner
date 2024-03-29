package game

import "time"

type ObjectType uint16

const (
	Bullet ObjectType = iota
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
		ID:        objectID.new(),
		Type:      objectType,
		Direction: direction,
		Pos:       pos,
		Owner:     p,
	}

	// Process as bomb.
	if objectType == Bomb {
		ticker := time.NewTicker(350 * time.Millisecond)
		explode := time.Now().Add(5 * time.Second)

		go func(g *Game, o *Object, explode time.Time) {
			for {
				select {
				case <-g.Active:
					ticker.Stop()
					return
				case <-ticker.C:
					if time.Now().After(explode) {
						ticker.Stop()
						// TODO: Check game is active.

						g.NewAction(BombExplode, o.Pos)

						g.Lock()
						// Bomb explodes and paints all tiles in 1 tile range (9 total).
						for x := (pos.X - 1); x <= (pos.X + 1); x++ {
							for y := (pos.Y - 1); y <= (pos.Y + 1); y++ {
								if x >= 0 && x <= (g.Size-1) && y >= 0 && y <= (g.Size-1) {
									o.Owner.RLock()
									g.Claims[x][y] = o.Owner.Team
									o.Owner.RUnlock()
								}
							}
						}

						// Remove bomb object.
						g.RemoveObject(o.ID)

						g.Unlock()
					}
				}
			}
		}(g, object, explode)
	}

	// Process as bullet.
	if objectType == Bullet {
		ticker := time.NewTicker(250 * time.Millisecond)

		// Calculate the starting position.
		var newPos Point

		switch direction {
		case "north":
			newPos = pos.North()
		case "south":
			newPos = pos.South()
		case "west":
			newPos = pos.West()
		case "east":
			newPos = pos.East()
		}

		// Set new position.
		object.Pos.X = newPos.X
		object.Pos.Y = newPos.Y

		go func(g *Game, o *Object) {
			for {
				select {
				case <-g.Active:
					ticker.Stop()
					return
				case <-ticker.C:
					// Bullet moves
					// Calculate the new position.
					var newPos Point

					switch o.Direction {
					case "north":
						newPos = o.Pos.North()
					case "south":
						newPos = o.Pos.South()
					case "west":
						newPos = o.Pos.West()
					case "east":
						newPos = o.Pos.East()
					}

					// Check if the object is trying to move outside the maze.
					g.Lock()
					if newPos.X < 0 || newPos.X > (g.Size-1) || newPos.Y < 0 || newPos.Y > (g.Size-1) {
						ticker.Stop()
						g.RemoveObject(o.ID)
						g.Unlock()
						return
					}
					g.Unlock()

					// Check if the object is trying to move into a wall.
					g.Lock()
					if g.Maze[newPos.X][newPos.Y] == Wall {
						ticker.Stop()
						g.RemoveObject(o.ID)

						// Claim wall tile.
						o.Owner.RLock()
						g.Claims[newPos.X][newPos.Y] = o.Owner.Team
						o.Owner.RUnlock()
						g.Unlock()
						return
					}
					g.Unlock()

					// Check if the object hits a player.
					for _, player := range g.Players {
						player.RLock()
						if player.ID != o.Owner.ID && newPos == *player.Pos {
							player.RUnlock()

							ticker.Stop()
							g.NewAction(PlayerHit, &newPos)
							g.PlayerHit(player)

							g.Lock()
							g.RemoveObject(o.ID)
							g.Unlock()

							return
						}
						player.RUnlock()
					}

					// Move to new position.
					o.Pos.X = newPos.X
					o.Pos.Y = newPos.Y
				}
			}
		}(g, object)
	}

	g.Objects = append(g.Objects, object)
}

// RemoveObject removes an object from the game.
func (g *Game) RemoveObject(id int) {
	for i, o := range g.Objects {
		if o.ID == id {
			g.Objects = append(g.Objects[:i], g.Objects[i+1:]...)
			return
		}
	}
}
