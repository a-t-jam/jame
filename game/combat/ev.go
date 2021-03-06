// ev (event) is for separating GUI from the internal state progress

package combat

import (
	"fmt"

	"github.com/a-t-jam/jame/game/scene"
)

// Event is a change to the combat world
type Event interface {
	run()
	anim() Anim
}

// TODO: use generational index

// Anim is always a melee attack animation
type Anim struct {
	actor  int
	target int
}

type Attack struct {
	attacker int
	target   int
}

func (a Attack) run() {
	fmt.Println("attack: ", a)

	attacker := &cState.actors[a.attacker]
	target := &cState.actors[a.target]

	target.Hp -= attacker.Atk

	// FIXME: this hack
	if a.target == 0 {
		scene.PlayerIq = target.Hp
	}
}

func (a Attack) anim() Anim {
	return Anim{
		actor:  a.attacker,
		target: a.target,
	}
}
