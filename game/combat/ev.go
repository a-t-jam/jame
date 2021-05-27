package combat

import (
	"fmt"
)

// Event is a change to the combat world
type Event interface {
	run()
}

type Attack struct {
	attacker int
	target   int
}

func (a Attack) run() {
	fmt.Println("attack action:", a)
}
