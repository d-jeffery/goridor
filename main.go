package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goridor/goridor"
	"log"
)

func main() {
	game, err := goridor.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(goridor.ScreenWidth, goridor.ScreenHeight)
	ebiten.SetWindowTitle("GORIDOR")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
