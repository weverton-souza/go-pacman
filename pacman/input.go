package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func keyPressed() input {
	if inpututil.KeyPressDuration(ebiten.KeyUp) > 0 || inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		return up
	}

	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		return left
	}

	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 || inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		return right
	}

	if inpututil.KeyPressDuration(ebiten.KeyDown) > 0 || inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		return down
	}

	return 0
}
