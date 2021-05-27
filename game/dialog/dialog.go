package dialog

import (
	_ "embed"
	"log"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

var (
	pixelFont       font.Face
	selectedDialogs []string
	displayDialog   string
	i               int
)

func init() {
	selectedDialogs = dialogs["entry"]
	i = 0
	displayDialog = selectedDialogs[i]

	fontBytes, err := assets.Data.ReadFile("fonts/8bitOperatorPlus8-Regular.ttf")
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	pixelFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Update(scene *scene.Scene) error {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) && i < len(selectedDialogs)-1 {
		i = i + 1
	}
	displayDialog = selectedDialogs[i]
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	text.Draw(screen, displayDialog, pixelFont, 40, 40, color.White)
}
