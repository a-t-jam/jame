package scene

type Item struct {
	//
}

type Actor struct {
	Combat
}

type Combat struct {
	Alive    bool
	IsFriend bool
	// energy to take turn
	Energy     uint
	EnergyGain uint
	// states
	MaxHp uint
	Hp    uint
	Atk   uint
	Def   uint
}

type Scene struct {
	Len       uint
	Pos       uint
	Inventory []*Item
	Ducks     []*Actor
}
