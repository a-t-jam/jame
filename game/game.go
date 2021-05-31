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

type Game struct {
	Scene scene.Scene
}

func New() Game {
	return Game{
		Scene: scene.Scene{
			State:  scene.TravelState,
			Len:    10,
			Pos:    0,
			Player: combat.DefaultPlayer,
		},
	}
}

func (g *Game) Update() error {
	switch g.Scene.State {
	case scene.TravelState:
		return travel.Update(&g.Scene)
	case scene.CombatState:
		return combat.Update(&g.Scene)
	// case scene.DialogState:
	//     // FIXME:
	//     return dialog.Update(&g.Scene)
	default:
		log.Fatal("wrong state")
		return nil
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Scene.State {
	case scene.TravelState:
		travel.Draw(&g.Scene, screen)
	case scene.CombatState:
		combat.Draw(&g.Scene, screen)
	case scene.DialogState:
		// FIXME:
		dialog.Draw(&g.Scene, screen)
	default:
		log.Fatal("wrong state")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
