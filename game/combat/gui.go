package combat

import (
	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	AttackTex  *ebiten.Image = assets.LoadImg("pipoya/attack.png")
	SwingSound *audio.Player = assets.LoadWav("se/match/swing.wav")
)

func NewAttackSurface() *ui.Surface {
	s := ui.NewAnimSurface(AttackTex, 8, 1)
	s.Scale = [2]float64{2.0, 2.0}
	return s
}
