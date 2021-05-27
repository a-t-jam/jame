package combat

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
