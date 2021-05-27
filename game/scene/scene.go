package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Combat struct {
	// GUI
	Img *ebiten.Image
	// status
	Alive    bool
	IsFriend bool
	// states
	MaxHp uint
	Hp    uint
	Atk   uint
	Def   uint
}

type Scene struct {
	Len    uint
	Pos    uint
	Player Combat
}
