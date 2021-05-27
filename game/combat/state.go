package combat

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/game/dialog"
)

import (
	"log"
)

type GuiState int

// Combat GUI states
const (
	Tick = iota
	Anim
	PlayerInput
        Dialog
)

type StateStack struct {
	states []GuiState
}

func (s *StateStack) Clear() {
	s.states = []GuiState{}
}

func (s *StateStack) IsEmpty() bool {
	return len(s.states) == 0
}

func (s *StateStack) Top() GuiState {
	last := len(s.states) - 1
	return s.states[last]
}

func (s *StateStack) Push(state GuiState) {
	s.states = append(s.states, state)
}

func (s *StateStack) Pop() GuiState {
	if s.IsEmpty() {
		log.Fatal("error: stack is empty")
		return 0
	}

	last := len(s.states) - 1
	pop := s.states[last]
	s.states = s.states[:last]

	return pop
}

// ----------------------------------------

var (
	PlayerEvent *Event
)

func runAction(action Event) {
	// `action.run` will create animation and mutate animation queue
	action.run()
	// and we'll play the queued animation in `Anim` state
	state.guiState.Push(Anim)
}

func updateTick(scene *scene.Scene) {
	// when we selected player event in `updatePlayerInput`
	if PlayerEvent != nil {
		runAction(*PlayerEvent)
		PlayerEvent = nil
		state.inc()
	}

	for {
		if state.handleStatus() {
			return
		}

		actorIx := state.cur
		actor := &state.actors[actorIx]

		// skip dead actors
		if !actor.Alive {
			state.inc()
			continue
		}

		fmt.Println("Actor", actorIx, "takes turn")
		action := takeTurn(actorIx)

		if action == nil {
			// it's player action. let's enter `PlayerInput` state
			state.guiState.Push(PlayerInput)
			return
		} else {
			runAction(action)
			state.inc()
			return
		}
	}
}

func updateAnim(scene *scene.Scene) {
	// play or wait for event animation

	// we've finished playing animation. go back to the tick-the-game state
	state.guiState.Pop()
}

func updatePlayerInput(scene *scene.Scene) {
	var ev Event
	//print("get player input")

	// TODO: use GUI
	// FIXME: for now I'm doing a space button to test, must be done with button GUI
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		print("Player attacks")
		ev = &Attack{
			attacker: 0, // 0 == player
			target:   1,
		}
	}
	if ev != nil {
		PlayerEvent = &ev
		// go back to the tick state
		state.guiState.Pop()
                state.guiState.Push(Dialog)
	}
}

func updateDialog(scene *scene.Scene) {
    print("hi from dialog")
    dl := dialog.Update(scene, dialog.Dialogs["player_attack"])
    if(dl == nil) {
        state.guiState.Pop()
        state.guiState.Push(Tick)
    }
}
