package combat

import (
	_ "embed"

	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/scene"
)

var (
	PlayerImage   *ebiten.Image
	DefaultPlayer scene.Combat
	//
	Enemy1Image *ebiten.Image
	Enemy1      scene.Combat
)

func init() {
	// FIXME:
	PlayerImage = assets.LoadImg("sprites/amg1_rt2.png")

	DefaultPlayer = scene.Combat{
		Img:      PlayerImage,
		Alive:    true,
		IsFriend: true,
		MaxHp:    100,
		Hp:       100,
		Atk:      50,
		Def:      50,
	}

	// FIXME:
	Enemy1Image = assets.LoadImg("sprites/amg1_rt2.png")

	Enemy1 = scene.Combat{
		Img:      Enemy1Image,
		Alive:    true,
		IsFriend: false,
		MaxHp:    100,
		Hp:       100,
		Atk:      50,
		Def:      50,
	}
}
