package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Combat struct {
	// GUI
	Img *ebiten.Image
	// status
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
	DialogState
)

type Scene struct {
	State  GameState
	Len    uint
	Pos    uint
	Player Combat
}
