package win

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

func Enter(scene_ *scene.Scene) {
	assets.WinSound.Play()
}

func Update(scene_ *scene.Scene) error {
	return nil
}

func Draw(scene_ *scene.Scene, screen *ebiten.Image) {
	assets.DrawOcean2(screen)

	mes := "Duck enter the surface!\nThank you for playing <3"
	face := assets.PixelFont
	x := 1280 / 2
	y := 720 / 2

	bounds := text.BoundString(face, mes)
	x -= (bounds.Max.X - bounds.Min.X) / 2.0
	y -= (bounds.Max.Y - bounds.Min.Y) / 2.0

	text.Draw(screen, mes, assets.PixelFont, x+4, y+4, color.Gray16{Y: 32})
	text.Draw(screen, mes, face, x, y, color.White)
}
