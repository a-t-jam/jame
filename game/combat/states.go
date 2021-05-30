package combat

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/ui"
)

type GuiState int

// Combat GUI states
const (
	Tick = iota
	AnimState
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
	fmt.Printf("    after push: %#v\n", s.states)
}

func (s *StateStack) Pop() GuiState {
	if s.IsEmpty() {
		log.Fatal("error: stack is empty")
		return 0
	}

	last := len(s.states) - 1
	pop := s.states[last]
	s.states = s.states[:last]

	fmt.Printf("    after pop: %#v\n", s.states)

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
	pushAnim(action.anim())
}

func updateTick(s *scene.Scene) {
	// when we selected player event in `updatePlayerInput`
	if PlayerEvent != nil {
		runAction(*PlayerEvent)
		PlayerEvent = nil
		cState.inc()
		return
	}

	for {
		if cState.handleStatus(s) {
			s.State = scene.TravelState
			return
		}

		actorIx := cState.cur
		actor := &cState.actors[actorIx]

		// skip dead actors
		if !actor.Alive() {
			cState.inc()
			continue
		}

		fmt.Println("Actor", actorIx, "takes turn")
		action := takeTurn(actorIx)

		if action == nil {
			// it's player action. let's enter `PlayerInput` state
			cState.guiState.Push(PlayerInput)
			return
		} else {
			runAction(action)
			cState.inc()
			return
		}
	}
}

type AnimRunState struct {
	anims []Anim
	// true while we're playing the first element of `anims`
	playing bool
	start   time.Time
}

// pushAnim pushes a given animation to the animaton queue and lets you enter `AnimState`
func pushAnim(anim Anim) {
	if cState.guiState.Top() != AnimState {
		cState.guiState.Push(AnimState)
	}

	aState.anims = append(aState.anims, anim)
	aState.playing = false
}

var (
	aState AnimRunState
)

func updateAnim(scene *scene.Scene) {
	if len(aState.anims) == 0 {
		cState.guiState.Pop()
		aState.playing = false
		return
	}

	anim := &aState.anims[0]

	// new animation: create new node
	if !aState.playing {
		aState.playing = true
		aState.start = time.Now()

		node := ui.Node{Align: ui.AlignCenter, Surface: NewAttackSurface()}

		if anim.target == 0 {
			node.X = 1280.0 / 2.0
			node.Y = 720 - 200
		} else {
			node.X = 1280.0 / 2.0
			node.Y = 200
		}

		cState.nodes = append(cState.nodes, node)
	}

	// play fixed 8 frame animation
	n_frames := 8
	n_wait_frames := 12

	ms := time.Since(aState.start).Milliseconds()
	frame := int(ms / (1000 / 60))

	if frame == 0 {
		SwingSound.Rewind()
		SwingSound.Play()
	}

	if frame < n_frames {
		node := &cState.nodes[2]
		node.Surface.CurrentFrameIx = frame
	}

	// at end of the animation
	if frame >= n_frames {
		// remove the node for the animation
		cState.nodes = cState.nodes[0:2]
	}

	// on finish: go to next animation
	if frame >= n_frames+n_wait_frames {
		// dequeue the animation description
		aState.anims = aState.anims[1:]
		aState.playing = false
	}
}

func updatePlayerInput(scene *scene.Scene) {
	var ev Event
	//print("get player input")

	// TODO: use GUI
	// FIXME: for now I'm doing a space button to test, must be done with button GUI
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		println("Player attacks")
		ev = &Attack{
			attacker: 0, // 0 == player
			target:   1,
		}
	}

	if ev != nil {
		PlayerEvent = &ev
		// go back to the tick state
		cState.guiState.Pop()
	}
}

//func updateDialog(scene *scene.Scene) {
//	println("hi from dialog")
//	dl := dialog.Update(scene, dialog.Dialogs["player_attack"])
//	if dl == nil {
//		cState.guiState.Pop()
//	}
//}
