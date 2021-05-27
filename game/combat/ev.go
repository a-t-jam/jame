package combat

// Event is a change to the combat world
type Event interface {
	run()
}

type Attack struct {
	attacker uint
	target   uint
}

func (a *Attack) run() {
	//
}
