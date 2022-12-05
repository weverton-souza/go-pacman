package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/weverton-souza/go-pacman/pacman"
	"github.com/weverton-souza/go-pacman/pacman/handler"
	_ "image/png"
)

func main() {
	game := pacman.NewGame()

	err := ebiten.Run(game.Update, game.ScreenWidth(), game.ScreenHeight(), 1.5, "Pacman")
	handler.HandleError(handler.RUNTIME, err)
}
