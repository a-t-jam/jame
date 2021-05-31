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
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed duck pipoya sprites se
//go:embed bg fonts
var Data embed.FS

var (
	Audio = audio.NewContext(48000)
)

var (
	PixelFont font.Face
	DebugFont font.Face
)

var (
	BattleDuck *ebiten.Image
	Bubble     *audio.Player
	TravelDuck *ebiten.Image
	Ocean1     *ebiten.Image
	Ocean2     *ebiten.Image
)

func init() {
	BattleDuck = LoadImg("duck/N-wait.png")
	Bubble = LoadWav("se/onjin/bubble_04.wav")
	TravelDuck = LoadImg("duck/N-walk.png")
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
	fontBytes, err := Data.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

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

func LoadWav(path string) *audio.Player {
	data, err := Data.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	decoded, err := wav.Decode(Audio, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	player, err := audio.NewPlayer(Audio, decoded)
	if err != nil {
		log.Fatal(err)
	}

	return player
}
