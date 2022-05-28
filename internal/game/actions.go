package game

type ActionType uint16

const (
	GameStart ActionType = iota
	GameEnd
	Shoot
	BombPlace
	BombExplode
	PlayerHit
)

type Action struct {
	ID   int        `json:"id"`
	Type ActionType `json:"type"`
	Pos  *Point     `json:"pos,omitempty"`
}

// NewAction returns a new activity.
func (g *Game) NewAction(activityType ActionType, pos *Point) {
	g.Lock()
	defer g.Unlock()

	action := &Action{
		ID:   actionID.new(),
		Type: activityType,
		Pos:  pos,
	}

	g.ActionLog = append(g.ActionLog, action)
}
