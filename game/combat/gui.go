package combat

import (
	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
)

func NewAttackSurface() *ui.Surface {
	s := ui.NewAnimSurface(assets.AttackTex, 8, 1)
	s.Scale = [2]float64{2.0, 2.0}
	return s
}
