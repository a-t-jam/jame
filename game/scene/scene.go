package scene

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
)

var (
	PlayerPos int
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

func DrawPlayerPos(screen *ebiten.Image) {
	// draw the distance
	mes := fmt.Sprintf("%d steps away from the surface", 15-PlayerPos)
	face := assets.PixelFont
	x := 50
	y := 50

	text.Draw(screen, mes, face, x+4, y+4, color.Gray16{Y: 32})
	text.Draw(screen, mes, face, x, y, color.White)

}
