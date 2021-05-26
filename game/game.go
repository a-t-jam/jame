package game

import(
	_ "embed"

	_ "image/png"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/game/travel"
	"github.com/a-t-jam/jame/game/combat"
)

// GameState is the global game state
type GameState = int

// Game states
const (
	TravelState = iota
	CombatState
)

type Game struct{
	State GameState
	Scene scene.Scene
}

func New() Game {
	return Game {
		State: TravelState,
		Scene: scene.Scene {
			Len: 10,
			Pos: 0,
			Inventory: nil,
			Ducks: nil,
		},
	}
}

func (g *Game) Update() error {
	if g.State == TravelState {
		return travel.Update(&g.Scene);
	} else if g.State == CombatState {
		return combat.Update(&g.Scene);
	}

	return nil; // ?
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == TravelState {
		travel.Draw(&g.Scene, screen);
	} else if g.State == CombatState {
		combat.Draw(&g.Scene, screen);
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
