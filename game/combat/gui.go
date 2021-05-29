package combat

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	AttackTex *ebiten.Image
)

func init() {
	AttackTex = assets.LoadImg("pipoya/effects/006.png")
}
