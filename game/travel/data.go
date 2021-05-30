package travel

import (
	_ "embed"

	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/ui"
)

var (
	Surface *ui.Surface
)

func init() {
	surface := ui.NewAnimSurface(assets.TravelDuck, 4, 1)
	surface.Scale = [2]float64{2.0, 2.0}
	Surface = surface
}
