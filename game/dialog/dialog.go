package dialog

import (
	_ "embed"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

var (
	selectedDialogs []string
	displayDialog   string
	i               int
)

func init() {
	i = 0
}

func Update(scene *scene.Scene, dialogInput []string) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && i < len(selectedDialogs)-1 {
		i = i + 1
	}
	selectedDialogs = dialogInput
	displayDialog = selectedDialogs[i]
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	mes := displayDialog
	face := assets.LargePixelFont
	x := 1280 / 2
	y := 200

	bounds := text.BoundString(face, mes)
	x -= (bounds.Max.X - bounds.Min.X) / 2.0
	y -= (bounds.Max.Y - bounds.Min.Y) / 2.0

	text.Draw(screen, mes, face, x+4, y+4, color.Gray16{Y: 32})
	text.Draw(screen, mes, face, x, y, color.White)
}
