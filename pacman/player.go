package pacman

import (
	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/weverton-souza/go-pacman/images"
)

type player struct {
	images     [8]*ebiten.Image
	currentImg int
	curPos     pos
}

func newPlayer(y, x int) *player {
	p := &player{}
	p.loadImages()
	p.curPos = pos{y, x}
	return p
}

func (p *player) loadImages() {
	copy(p.images[:], loadImages(pacimages.PlayerImages[:]))
}

func (p *player) image() *ebiten.Image {
	return p.images[p.currentImg]
}

func (p *player) draw(screen *ebiten.Image) {
	x := float64(p.curPos.x * stageBlocSize)
	y := float64(p.curPos.y * stageBlocSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	err := screen.DrawImage(p.image(), op)
	if err != nil {
		return
	}
}
