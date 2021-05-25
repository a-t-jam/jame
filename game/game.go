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
var(
    bg *ebiten.Image
    playerSprite *ebiten.Image
)

// example structs
type Pos struct {
    X, Y int
}

type Entity struct {
    Pos
    Name string
}

type Item struct {
    Entity
}

type Actor struct {
    Entity
    Strength int
    Health int
}

type Player struct {
    Actor
    Inventory []*Item
}

func Init() {
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
        op := &ebiten.DrawImageOptions{}
        op1 := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((1280-790)/2, (720-480)/2)
	screen.DrawImage(bg, op)

        // sprite 
        op1.GeoM.Translate(50, 100)
        op1.GeoM.Scale(2, 2)
	screen.DrawImage(playerSprite, op1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

