package combat

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/dialog"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/ui"
	"github.com/hajimehoshi/ebiten/v2"
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
		X:     1280.0 / 2.0,
		Y:     720.0 - 200.0,
		Align: ui.AlignCenter,
		Surface: ui.Surface{
			Img: state.actors[0].Img,
		},
	})

	state.nodes = append(state.nodes, ui.Node{
		X:     1280.0 / 2.0,
		Y:     200.0,
		Align: ui.AlignCenter,
		Surface: ui.Surface{
			Img: state.actors[1].Img,
		},
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

	for _, node := range state.nodes {
		node.Draw(screen)
	}

	dialog.Draw(scene, screen)
}

func drawCentered(screen *ebiten.Image, img *ebiten.Image, x float64, y float64) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)

	w, h := img.Size()
	opts.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

	screen.DrawImage(img, &opts)
}
