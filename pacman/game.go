package pacman

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	scene *scene
}

func NewGame() *Game {
	g := &Game{}
	g.scene = newScene(nil)
	return g
}

func (g *Game) ScreenWidth() int {
	return g.scene.screenWidth()
}

func (g *Game) ScreenHeight() int {
	return g.scene.screenHeight()
}

func (g *Game) Update(screen *ebiten.Image) error {
	return g.scene.update(screen)
}
