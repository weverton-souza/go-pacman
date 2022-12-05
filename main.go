package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/weverton-souza/go-pacman/pacman"
	_ "image/png"
	"log"
)

func main() {
	game := pacman.NewGame()

	if err := ebiten.Run(game.Update, game.ScreenWidth(), game.ScreenHeight(), 1.5, "Pacman"); err != nil {
		log.Fatal(err)
	}
}
