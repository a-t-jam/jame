package combat

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/dialog"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	state State
)

type State struct {
	// 0: player,
	// 1: enemy
	actors   [2]scene.Combat
	nodes    []ui.Node
	cur      int
	guiState StateStack
	// running animation (play GUI of it!)
}

// Combat scene status
const (
	Continue = iota
	Defeated
	Win
)

func (c *State) inc() {
	c.cur += 1
	c.cur %= len(state.actors)
}

func (state *State) enemies() []scene.Combat {
	return state.actors[1:]
}

func (state *State) status() int {
	actors := &state.actors

	if !(actors[0].Alive()) {
		return Defeated
	}

	if !(actors[1].Alive()) {
		return Win
	}

	return Continue
}

// handleStatus returns true if the combat is finished
func (state *State) handleStatus(scene *scene.Scene) bool {
	switch state.status() {
	case Continue:
		return false
	case Defeated:
		fmt.Println("dead end")
		return true
	case Win:
		fmt.Println("win")
		return true
	default:
		log.Fatal("wrong status")
		return true
	}
}

// Enter initializes the combat scene
func Enter(scene *scene.Scene, enemy scene.Combat) {
	state = State{}
	state.guiState.Push(Tick)

	state.actors[0] = scene.Player
	state.actors[1] = enemy

	// TODO: easier use?
	state.nodes = append(state.nodes, ui.Node{
		X:       1280.0 / 2.0,
		Y:       720.0 - 200.0,
		Align:   ui.AlignCenter,
		Surface: state.actors[0].Surface,
	})

	state.nodes = append(state.nodes, ui.Node{
		X:       1280.0 / 2.0,
		Y:       200.0,
		Align:   ui.AlignCenter,
		Surface: state.actors[1].Surface,
	})
}

func Update(scene *scene.Scene) error {
	// see `state.go`
	switch state.guiState.Top() {
	case Tick:
		updateTick(scene)
	case Anim:
		updateAnim(scene)
	case PlayerInput:
		updatePlayerInput(scene)
	case Dialog:
		updateDialog(scene)
	default:
		log.Fatalln("wrong combat state")
	}

	return nil
}

// takeTurn returns an `Event` that an actor (`Combat`) invokes
func takeTurn(actorIx int) Event {
	actor := &state.actors[actorIx]

	if actor.IsFriend {
		// it's player. let the user select their action (see `state.go`)
		return nil
	} else {
		// it's enemy
		return Attack{
			attacker: actorIx,
			target:   0,
		}
	}
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	assets.DrawOcean1(screen)

	updateAnims(scene, screen)

	for _, node := range state.nodes {
		node.Draw(screen)
	}

	dialog.Draw(scene, screen)

	drawDebug(scene, screen)
}

func updateAnims(scene_ *scene.Scene, screen *ebiten.Image) {
	// the duck animation
	s := state.nodes[0].Surface

	elapsed := time.Since(scene.StartTime)
	n := elapsed.Milliseconds() / (1000 * 8 / 60)

	n_frames := len(s.Uvs)
	n_pingpong := (n_frames)*2 - 1

	frame := int(n) % n_pingpong
	if frame >= n_frames {
		frame -= n_frames
	}

	s.CurrentFrameIx = frame
}

func drawDebug(scene *scene.Scene, screen *ebiten.Image) {
	message := fmt.Sprintf("FPS: %v\n", ebiten.CurrentFPS())

	message += "State stack:"
	for _, s := range state.guiState.states {
		message += fmt.Sprintf(" %d", int(s))
	}

	message += "\n"
	message += fmt.Sprintf("player: %#v", state.actors[0])

	message += "\n"
	message += fmt.Sprintf("enemy: %#v", state.actors[1])

	text.Draw(screen, message, assets.DebugFont, 40, 340, color.White)
}
