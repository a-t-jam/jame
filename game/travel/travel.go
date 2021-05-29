package travel

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

var (
	playerSprite *ebiten.Image
)

func init() {
	playerSprite = assets.LoadImg("sprites/amg1_rt2.png")
}

func Update(scene *scene.Scene) error {
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	assets.DrawOcean1(screen)

	// sprite
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 100)
	op.GeoM.Scale(2, 2)
	screen.DrawImage(playerSprite, &op)

	debugDraw(scene, screen)
}

func debugDraw(scene *scene.Scene, screen *ebiten.Image) {
	message := fmt.Sprintf("FPS: %v", ebiten.CurrentFPS())
	text.Draw(screen, message, assets.PixelFont, 100.0, 300.0, color.White)
}
