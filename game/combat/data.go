package combat

import (
	_ "embed"

	_ "image/jpeg"
	_ "image/png"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/ui"
)

var (
	DefaultPlayer scene.Combat
	Enemy1        scene.Combat
)

func init() {
	DefaultPlayer = scene.Combat{
		Surface:  ui.NewAnimSurface(assets.BattleDuck, 5, 1),
		IsFriend: true,
		MaxHp:    100,
		Hp:       100,
		Atk:      50,
		Def:      50,
	}

	Enemy1 = scene.Combat{
		Surface:  ui.NewImageSurface(assets.LoadImg("sprites/sotrak_rewop.png")),
		IsFriend: false,
		MaxHp:    100,
		Hp:       100,
		Atk:      50,
		Def:      50,
	}
}
