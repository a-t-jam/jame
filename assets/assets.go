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
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed duck pipoya sprites se
//go:embed bg fonts
var Data embed.FS

var (
	Audio = audio.NewContext(48000)
)

var (
	PixelFont      font.Face = LoadFont("fonts/8bitOperatorPlus8-Regular.ttf", 72, 24)
	LargePixelFont font.Face = LoadFont("fonts/8bitOperatorPlus8-Regular.ttf", 72, 48)
	// DebugFont font.Face = LoadFont("fonts/8bitOperatorPlus8-Regular.ttf", 72, 14)
)

var (
	BattleDuck *ebiten.Image = LoadImg("duck/N-wait.png")
	TravelDuck *ebiten.Image = LoadImg("duck/N-walk.png")
	Bubble     *audio.Player = LoadWav("se/onjin/bubble_04.wav")
	AttackTex  *ebiten.Image = LoadImg("pipoya/attack.png")
)
var (
	DeathSound *audio.Player = LoadWav("se/match/death.wav")
	SwingSound *audio.Player = LoadWav("se/match/swing.wav")
	WinSound   *audio.Player = LoadMp3("se/onjin/win.mp3")
)

var (
	Ocean1 *ebiten.Image = LoadImg("bg/ocean1.jpg")
	Ocean2 *ebiten.Image = LoadImg("bg/ocean2.jpg")
)

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

func LoadMp3(path string) *audio.Player {
	data, err := Data.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	decoded, err := mp3.Decode(Audio, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	player, err := audio.NewPlayer(Audio, decoded)
	if err != nil {
		log.Fatal(err)
	}

	return player
}
