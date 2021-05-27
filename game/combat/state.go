package combat

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/game/scene"
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

func updateTick(scene *scene.Scene) {
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
			state.inc()
			continue
		}

		action.run()
		// TODO: play animation

		state.inc()
		return
	}
}

func updateAnim(scene *scene.Scene) {
	// play or wait for event animation
}

func updatePlayerInput(scene *scene.Scene) {
	// decide action (`Event`) of the player
}
