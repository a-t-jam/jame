// ui is for retained-mode rendering

package ui

import (
	"github.com/a-t-jam/jame/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	sprites *[]*ebiten.Image
}

type Node struct {
	X float64
	Y float64
	Surface
}

func (n *Node) draw(target *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(n.X, n.Y)
	n.Surface.draw(target, &opts)
}

// Surface is an image surface (how to render)
// TODO: enable text
type Surface struct {
	Img *ebiten.Image
}

func (s *Surface) draw(target *ebiten.Image, opts *ebiten.DrawImageOptions) {
	target.DrawImage(s.Img, opts)
}
