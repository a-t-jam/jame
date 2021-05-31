package scene

import (
	"time"

	_ "embed"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
)

var (
	StartTime       time.Time
	WalkDuckSurface *ui.Surface
)

func init() {
	StartTime = time.Now()

	WalkDuckSurface = ui.NewAnimSurface(assets.TravelDuck, 4, 1)
	WalkDuckSurface.Scale = [2]float64{2.0, 2.0}
}

type Combat struct {
	// GUI
	Surface *ui.Surface
	// meta
	IsFriend bool
	// states
	MaxHp int
	Hp    int
	Atk   int
	Def   int
}

func (c *Combat) Alive() bool {
	return c.Hp > 0
}

// GameState is the global game state
type GameState = int

// Game states
const (
	TravelState = iota
	CombatState
	WinState
)

type Scene struct {
	State  GameState
	Len    uint
	Pos    uint
	Player Combat
}
