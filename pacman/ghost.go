package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/weverton-souza/go-pacman/pacman/handler"
)

type ghost struct {
	kind       elem
	currentImg int
	curPos     pos
}

func newGhost(y, x int, k elem) *ghost {
	return &ghost{
		kind:   k,
		curPos: pos{y, x},
	}
}

func (g *ghost) image(imgs []*ebiten.Image) *ebiten.Image {
	return imgs[g.currentImg]
}

// draw the ghost
func (g *ghost) draw(screen *ebiten.Image, imgs []*ebiten.Image) {
	x := float64(g.curPos.x * stageBlocSize)
	y := float64(g.curPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	err := screen.DrawImage(g.image(imgs), op)
	handler.HandleError(handler.RUNTIME, err)
}
