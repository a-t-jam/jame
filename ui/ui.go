// ui is for retained-mode rendering

package ui

import (
	"image"
	"log"

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
	Surface  *Surface
	BaseOpts ebiten.DrawImageOptions
}

func (n *Node) Draw(target *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(n.X, n.Y)

	switch n.Align {
	case AlignLeftUp:
	case AlignCenter:
		frame := n.Surface.CurrentFrame()

		w, h := frame.Size()
		x := -float64(w) / 2.0
		y := -float64(h) / 2.0

		opts.GeoM.Translate(x, y)
	}

	n.Surface.Draw(target, &opts)
}

// UvRect specifies sub region of a texture with normalized coordinates
type UvRect struct {
	x float64
	y float64
	w float64
	h float64
}

// Surface is an image surface (how to render)
// TODO: text surface
// TODO: refactor allocation-related code (new, append)
type Surface struct {
	Img            *ebiten.Image
	Uvs            []UvRect
	CurrentFrameIx int
}

func NewImageSurface(img *ebiten.Image) *Surface {
	s := new(Surface)

	s.Img = img
	s.Uvs = append(s.Uvs, UvRect{x: 0.0, y: 0.0, w: 1.0, h: 1.0})

	return s
}

func NewAnimSurface(img *ebiten.Image, n_x int, n_y int) *Surface {
	s := new(Surface)

	s.Img = img
	w := 1.0 / float64(n_x)
	h := 1.0 / float64(n_y)

	for i_y := 0; i_y < n_y; i_y++ {
		for i_x := 0; i_x < n_x; i_x++ {
			x := float64(i_x) * w
			y := float64(i_y) * h
			s.Uvs = append(s.Uvs, UvRect{x, y, w, h})
		}
	}

	return s
}

func (s *Surface) CurrentFrame() *ebiten.Image {
	return s.Frame(s.CurrentFrameIx).(*ebiten.Image)
}

func (s *Surface) Frame(frame int) image.Image {
	n_frames := len(s.Uvs)

	if frame > n_frames {
		log.Fatalln("wrong frame: ", frame)
	}

	uv := s.Uvs[frame]

	w, h := s.Img.Size()

	rect := image.Rectangle{}
	rect.Min.X = int(uv.x * float64(w))
	rect.Min.Y = int(uv.y * float64(h))
	rect.Max.X = rect.Min.X + int(uv.w*float64(w))
	rect.Max.Y = rect.Min.Y + int(uv.h*float64(h))

	return s.Img.SubImage(rect)
}

func (s *Surface) Draw(target *ebiten.Image, opts *ebiten.DrawImageOptions) {
	target.DrawImage(s.CurrentFrame(), opts)
}
