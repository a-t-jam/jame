package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a-t-jam/jame/game"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hello, World!")

	game := game.New()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
