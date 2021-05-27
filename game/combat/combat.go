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
	// 0: player,
	// 1: enemy
	actors [2]scene.Combat
	cur    int
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

	if !(actors[0].Alive) {
		return Defeated
	}

	if !(actors[1].Alive) {
		return Win
	}

	return Continue
}

// Enter initializes the combat scene
func Enter(scene *scene.Scene, enemy scene.Combat) {
	state = State{}

	state.actors[0] = scene.Player
	state.actors[1] = enemy
}

func Update(scene *scene.Scene) error {
	for {
		switch state.status() {
		case Continue:
		case Defeated:
			fmt.Println("dead end")
			return nil
		case Win:
			fmt.Println("win")
			return nil
		default:
			log.Fatal("wrong status")
		}

		actor := &state.actors[state.cur]

		// skip dead actors
		if !actor.Alive {
			state.inc()
			continue
		}

		fmt.Println("Actor", state.cur, "takes turn")
		action := takeTurn(state.cur)

		if action == nil {
			state.inc()
			continue
		}

		// TODO: run action and play animation

		return nil
	}
}

func takeTurn(actorIx int) Event {
	actor := &state.actors[actorIx]
	if actor.IsFriend {
		// it's player
		// TODO: enter combat player input state
		return Attack{
			attacker: actorIx,
			target:   1,
		}
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
}
