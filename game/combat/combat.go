package combat

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	state State
)

type State struct {
	// 0 ~ 2: ducks. 3 ~: ememies
	actors [8]scene.Combat
	cur    int
}

const (
	Continue = iota
	DucksDied
	EnemiseDied
)

func (c *State) inc() {
	c.cur += 1
	c.cur %= len(state.actors)
}

func (state *State) friends() []scene.Combat {
	return state.actors[0:3]
}

func (state *State) enemies() []scene.Combat {
	return state.actors[3:]
}

func (state *State) status() int {
	actors := &state.actors

	if !(actors[0].Alive || actors[1].Alive || actors[2].Alive) {
		return DucksDied
	}

	if !(actors[3].Alive || actors[4].Alive || actors[5].Alive || actors[6].Alive || actors[7].Alive) {
		return EnemiseDied
	}

	return Continue
}

// Enter initializes the combat scene
func Enter(scene *scene.Scene) {
	state = State{}
	// TODO: overwrite `state` with `scene.ducks`
}

func Update(scene *scene.Scene) error {
	for {
		switch state.status() {
		case Continue:
		case DucksDied:
			fmt.Println("dead end")
			return nil
		case EnemiseDied:
			fmt.Println("win")
			return nil
		default:
			log.Fatal("wrong status")
		}

		actor := &state.actors[state.cur]

		if !actor.Alive {
			state.inc()
			continue
		}

		fmt.Println("Actor", state.cur, "takes turn")
		action := takeTurn(actor)

		if action == nil {
			state.inc()
			continue
		}

		// TODO: run action and play animation

		return nil
	}
}

func takeTurn(actor *scene.Combat) Event {
	return nil
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	assets.DrawOcean1(screen)
}
