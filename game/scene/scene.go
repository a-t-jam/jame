package scene

// example structs
type Item struct {
    //
}

type Actor struct {
	Combat
}

type Combat struct {
	//
}

type Scene struct {
	Len uint
	Pos uint
	Inventory []*Item
	Ducks []*Actor
}
