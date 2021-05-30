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
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) && i < len(selectedDialogs)-1 {
		i = i + 1
	}
	selectedDialogs = dialogInput
	displayDialog = selectedDialogs[i]
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	text.Draw(screen, displayDialog, assets.PixelFont, 450, 50, color.White)
}
