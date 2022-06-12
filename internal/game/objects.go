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

	ticker := time.NewTicker(350 * time.Millisecond)
	explode := time.Now().Add(5 * time.Second)
	end := time.Now().Add(10 * time.Second)

	// Process as bomb.
	if objectType == Bomb {
		go func(g *Game, o *Object, end time.Time) {
			for {
				select {
				case <-g.Active:
					ticker.Stop()
					return
				case <-ticker.C:
					if time.Now().After(explode) {
						// TODO: Check game is active.

						g.NewAction(BombExplode, o.Pos)

						g.Lock()
						// Bomb explodes and paints all tiles in 1 tile range (9 total).
						for x := (pos.X - 1); x <= (pos.X + 1); x++ {
							for y := (pos.Y - 1); y <= (pos.Y + 1); y++ {
								if x >= 0 && x < (g.Size-1) && y >= 0 && y < (g.Size-1) {
									o.Owner.RLock()
									g.Claims[x][y] = o.Owner.Team
									o.Owner.RUnlock()
								}
							}
						}
						g.Unlock()

						// Remove bomb object.
						g.RemoveObject(o.ID)
					}

					if time.Now().After(end) {
						ticker.Stop()

						o.Owner.Lock()
						o.Owner.Abilities.BombAvailable = true
						o.Owner.Unlock()
					}
				}
			}
		}(g, object, end)
	}

	// Process as bullet.
	if objectType == Bullet {
		go func(g *Game, o *Object) {
			for {
				select {
				case <-g.Active:
					ticker.Stop()
					return
				case <-ticker.C:
					if time.Now().After(end) {
						ticker.Stop()

						// Bullet moves
						// Calculate the new position.
						var newPos Point

						switch o.Direction {
						case "north":
							newPos = p.Pos.North()
						case "south":
							newPos = p.Pos.South()
						case "west":
							newPos = p.Pos.West()
						case "east":
							newPos = p.Pos.East()
						}

						// Check if the player is trying to move outside the maze.
						g.RLock()
						if newPos.X < 0 || newPos.X > (g.Size-1) || newPos.Y < 0 || newPos.Y > (g.Size-1) {
							ticker.Stop()
							// TODO: Delete object.
							return
						}
						g.RUnlock()

						// Check if the player is trying to move into a wall.
						g.RLock()
						if g.Maze[newPos.X][newPos.Y] == Wall {
							ticker.Stop()
							// TODO: Delete object.
							return
						}
						g.RUnlock()

						// Check if another player is already in the new position.
						for _, player := range g.Players {
							player.RLock()
							if player.ID != p.ID && newPos == *player.Pos {
								ticker.Stop()
								g.NewAction(PlayerHit, pos)
								// TODO: Delete object.

								player.RUnlock()
								return
							}
							player.RUnlock()
						}

						// Move to new position.
						object.Pos.X = newPos.X
						object.Pos.Y = newPos.Y

						object.Owner.Lock()
						object.Owner.Abilities.ShootAvailable = true
						object.Owner.Unlock()
					}
				}
			}
		}(g, object)
	}

	g.Objects = append(g.Objects, object)
}

func (g *Game) RemoveObject(id int) {
	g.Lock()
	defer g.Unlock()

	for i, o := range g.Objects {
		if o.ID == id {
			g.Objects = append(g.Objects[:i], g.Objects[i+1:]...)
			return
		}
	}
}
