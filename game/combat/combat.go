package combat

import(
	"fmt"

	_ "image/png"
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/a-t-jam/jame/game/scene"
)

var(
    bg *ebiten.Image
)

// FIXME: duplicate image loading
func init() {
	bg = scene.LoadImg("winddorf/kyoto.jpg")
}

type CombatState struct {
	actors [10]scene.Combat
}

var (
	resume <-chan bool
	state CombatState
)

// genTurnState returns a coroutine implemented as a channel that holds turn-based game state.
//
// Returns if we should continue the combat scene.
func genTurnState(state *CombatState) chan bool {
	yield := make (chan bool);

	go func() {
		defer close(yield)
		for i := 1; i < 10; i++ {
			yield <- true;
		}
	}()

	return yield
}

// Enter initializes the combat scene
func Enter(scene *scene.Scene) {
	// TODO: overwrite `state` with `scene.ducks`
	state = CombatState {}
	resume = genTurnState(&state)
}

func Update(scene *scene.Scene) error {
	running := <-resume
	if !running {
		// TODO: it's over. pop the combat scene
	}
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	//
}
