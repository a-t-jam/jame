package game

import(
	"bytes"
	_ "embed"
	"image"
	"log"

	_ "image/png"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/a-t-jam/jame/assets"
)

// bg is a temporary background image.
var bg *ebiten.Image

func Init() {
	// byte data
	imgByte, err := assets.Winddorf.ReadFile("winddorf/kyoto.jpg")
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
	bg = ebitenImg
}

// GameState is the global game state
type GameState = int

// Game states
const (
	TravelState = iota
	CombatState
)

type Game struct{
	state GameState
}

func New() Game {
	return Game {
		state: TravelState,
		}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(bg, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

