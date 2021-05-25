package game

import(
	"bytes"
	"image"
	_ "image/png"
	_ "image/jpeg"
	"io/ioutil"
	"log"

	"github.com/rakyll/statik/fs"
	_ "github.com/a-t-jam/jame/statik"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// bg is a temporary background image.
var bg *ebiten.Image

func Init() {
	// load image byte data
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	imgFile, err := statikFS.Open("/winddorf/kyoto.jpg")
	if err != nil {
		log.Fatal(err)
	}

	imgBytes, err := ioutil.ReadAll(imgFile)
    if err != nil {
        panic(err)
    }

	// create std (?) image
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal(err)
	}

	// create ebiten image
	ebitenImage := ebiten.NewImageFromImage(img)
	bg = ebitenImage
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

