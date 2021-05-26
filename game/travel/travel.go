package travel

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

var (
	bg           *ebiten.Image
	playerSprite *ebiten.Image
)

func init() {
	// byte data
	imgByte, err := assets.Data.ReadFile("winddorf/kyoto.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	// std image
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	bg = ebitenImg

	// testing sprites
	imgByte, err = assets.Data.ReadFile("sprites/amg1_rt2.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, _, err = image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	ebitenImg = ebiten.NewImageFromImage(img)
	playerSprite = ebitenImg
}

func Update(scene *scene.Scene) error {
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	op := &ebiten.DrawImageOptions{}
	op1 := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((1280-790)/2, (720-480)/2)
	screen.DrawImage(bg, op)

	// sprite
	op1.GeoM.Translate(50, 100)
	op1.GeoM.Scale(2, 2)
	screen.DrawImage(playerSprite, op1)
}
