package game

import (
	"log"

	_ "embed"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a-t-jam/jame/game/combat"
	"github.com/a-t-jam/jame/game/dialog"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/game/travel"
)

// GameState is the global game state
type GameState = int

// Game states
const (
	TravelState = iota
	CombatState
	DialogState
)

type Game struct {
	State GameState
	Scene scene.Scene
}

func New() Game {
	return Game{
		State: TravelState,
		Scene: scene.Scene{
			Len:    10,
			Pos:    0,
			Player: combat.DefaultPlayer,
		},
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.State = TravelState
	}

	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.State = CombatState
		combat.Enter(&g.Scene, combat.Enemy1)
	}

	switch g.State {
	case TravelState:
		return travel.Update(&g.Scene)
	case CombatState:
		return combat.Update(&g.Scene)
	case DialogState:
		return combat.Update(&g.Scene)
	default:
		log.Fatal("wrong state")
		return nil
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case TravelState:
		travel.Draw(&g.Scene, screen)
	case CombatState:
		combat.Draw(&g.Scene, screen)
	case DialogState:
		dialog.Draw(&g.Scene, screen)
	default:
		log.Fatal("wrong state")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
