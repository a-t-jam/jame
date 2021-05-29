package assets

import (
	"bytes"
	"embed"
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed winddorf
//go:embed duck pipoya sprites
//go:embed bg fonts
var Data embed.FS

var (
	BattleDuck *ebiten.Image
	Bg         *ebiten.Image
	Ocean1     *ebiten.Image
	Ocean2     *ebiten.Image
	PixelFont  font.Face
	DebugFont  font.Face
)

func init() {
	BattleDuck = LoadImg("duck/N-wait.png")
	Bg = LoadImg("winddorf/kyoto.jpg")
	Ocean1 = LoadImg("bg/ocean1.jpg")
	Ocean2 = LoadImg("bg/ocean2.jpg")
	PixelFont = LoadFont("fonts/8bitOperatorPlus8-Regular.ttf", 72, 24)
	DebugFont = LoadFont("fonts/8bitOperatorPlus8-Regular.ttf", 72, 14)
}

func DrawOcean1(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1.0/4.0, 1.0/4.0)
	screen.DrawImage(Ocean1, &opts)
}

func DrawOcean2(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1.0/3.0, 1.0/3.0)
	screen.DrawImage(Ocean2, &opts)
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

func LoadFont(path string, dpi float64, size float64) font.Face {
	fontBytes, err := Data.ReadFile("fonts/8bitOperatorPlus8-Regular.ttf")
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	font_face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return font_face
}
