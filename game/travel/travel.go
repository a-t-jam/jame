package travel

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/dialog"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/ui"
)

var (
	playerSprite *ebiten.Image
	playerNode   ui.Node
	playerPos    int
)

func init() {
	playerPos = 0
	playerNode = ui.Node{X: 1280.0 / 2.0, Y: 720.0 - 200.0, Align: ui.AlignCenter, Surface: Surface}
}

func Update(scene *scene.Scene) error {
	if playerPos == 0 {
		return dialog.Update(scene, dialog.Dialogs["moving_instruction"])
	}
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	assets.DrawOcean1(screen)

	playerNode.Draw(screen)
	dialog.Draw(scene, screen)

	debugDraw(scene, screen)
}

func debugDraw(scene *scene.Scene, screen *ebiten.Image) {
	message := fmt.Sprintf("FPS: %v", ebiten.CurrentFPS())
	text.Draw(screen, message, assets.PixelFont, 100.0, 300.0, color.White)
}
