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
		ID:        objectID.new(),
		Type:      objectType,
		Direction: direction,
		Pos:       pos,
		Owner:     p,
	}

	ticker := time.NewTicker(350 * time.Millisecond)
	end := time.Now().Add(5 * time.Second)

	// Process as bomb.
	if objectType == Bomb {
		go func(g *Game, o *Object, end time.Time) {
			for {
				select {
				case <-g.Active:
					ticker.Stop()
					return
				case <-ticker.C:
					if time.Now().After(end) {
						ticker.Stop()

						g.NewAction(BombExplode, object.Pos)

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
