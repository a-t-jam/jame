package combat

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type CombatState struct {
	actors [10]scene.Combat
	cur    int
}

var (
	state CombatState
)

func takeTurn(actor *scene.Combat) Event {
	return nil
}

// Enter initializes the combat scene
func Enter(scene *scene.Scene) {
	state = CombatState{}
	// TODO: overwrite `state` with `scene.ducks`
}

func Update(scene *scene.Scene) error {
	fmt.Println(state.cur)
	actor := &state.actors[state.cur]

	action := takeTurn(actor)

	if action == nil {
		state.cur += 1
		state.cur %= len(state.actors)
		return nil
	}

	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	assets.DrawBg(screen)
}
