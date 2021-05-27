package assets

import (
	"bytes"
	"embed"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed winddorf
//go:embed sprites
//go:embed fonts
var Data embed.FS

var (
	Bg *ebiten.Image
)

func DrawBg(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((1280-790)/2, (720-480)/2)
	screen.DrawImage(Bg, op)
}

func init() {
	Bg = LoadImg("winddorf/kyoto.jpg")
}

// LoadImg loads an ebiten image from the `assets` directory
func LoadImg(path string) *ebiten.Image {
	// byte data
	imgByte, err := Data.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	// std image
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	// ebiten image
	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg
}
