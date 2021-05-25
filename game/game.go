package game

import(
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

