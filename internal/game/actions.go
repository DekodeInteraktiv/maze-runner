package game

type ActionType uint16

const (
	Shoot ActionType = iota
	BombPlace
	BombExplode
	PlayerHit
)

type Action struct {
	ID   int        `json:"id"`
	Type ActionType `json:"type"`
	Pos  *Point     `json:"pos"`
}

// NewAction returns a new activity.
func (g *Game) NewAction(activityType ActionType, pos *Point) {
	g.Lock()
	defer g.Unlock()

	action := &Action{
		ID:   0,
		Type: activityType,
		Pos:  pos,
	}

	g.ActionLog = append(g.ActionLog, action)
}
