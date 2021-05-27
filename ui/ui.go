// ui is for retained-mode rendering

package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Align = int

const (
	AlignLeftUp = iota
	AlignCenter
)

type SpriteSheet struct {
	sprites *[]*ebiten.Image
}

type Node struct {
	X float64
	Y float64
	Align
	Surface
	// z order + align?
}

func (n *Node) Draw(target *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(n.X, n.Y)

	switch n.Align {
	case AlignLeftUp:
	case AlignCenter:
		w, h := n.Surface.Img.Size()
		x := -float64(w) / 2.0
		y := -float64(h) / 2.0
		opts.GeoM.Translate(x, y)
	}

	n.Surface.Draw(target, &opts)
}

// Surface is an image surface (how to render)
// TODO: enable text
type Surface struct {
	Img *ebiten.Image
}

func (s *Surface) Draw(target *ebiten.Image, opts *ebiten.DrawImageOptions) {
	target.DrawImage(s.Img, opts)
}
