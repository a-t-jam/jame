package scene

import(
	"bytes"
	"image"
	"log"

	_ "image/png"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/a-t-jam/jame/assets"
)

// LoadImg loads an ebiten image from the `assets` directory
func LoadImg(path string) *ebiten.Image {
	// byte data
	imgByte, err := assets.Data.ReadFile(path)
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

