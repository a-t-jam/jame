package combat

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	AttackTex  *ebiten.Image
	SwingSound *audio.Player
)

func init() {
	AttackTex = assets.LoadImg("pipoya/effects/008.png")
	SwingSound = assets.LoadWav("se/match/swing.wav")
}

func NewAttackSurface() *ui.Surface {
	return ui.NewAnimSurface(AttackTex, 8, 1)
}
