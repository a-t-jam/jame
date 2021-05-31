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
	PlayerIq  int = 420
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
	mes := fmt.Sprintf("%d steps away from the surface", 15-PlayerPos)
	face := assets.PixelFont
	x := 1280 / 2
	y := 720 / 2

	bounds := text.BoundString(face, mes)
	x -= (bounds.Max.X - bounds.Min.X) / 2.0
	y -= (bounds.Max.Y - bounds.Min.Y) / 2.0

	text.Draw(screen, mes, face, x+4, y+4, color.Gray16{Y: 32})
	text.Draw(screen, mes, face, x, y, color.White)

}

func DrawPlayerIq(screen *ebiten.Image) {
	mes := fmt.Sprintf("IQ %d", PlayerIq)
	face := assets.PixelFont
	x := 1280 / 2
	y := 620

	bounds := text.BoundString(face, mes)
	x -= (bounds.Max.X - bounds.Min.X) / 2.0
	y -= (bounds.Max.Y - bounds.Min.Y) / 2.0

	text.Draw(screen, mes, face, x+2, y+2, color.Gray16{Y: 32})
	text.Draw(screen, mes, face, x, y, color.White)

}
