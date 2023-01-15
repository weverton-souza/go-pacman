package pacman

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
)

func isWall(e elem) bool {
	return w0 <= e && e <= w24
}

func loadImage(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	handleError(err)
	ebImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handleError(err)
	return ebImg
}

func loadImages(images [][]byte) []*ebiten.Image {
	var res []*ebiten.Image
	size := len(images)
	for i := 0; i < size; i++ {
		res = append(res, loadImage(images[i]))
	}
	return res
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func canMove(m [][]elem, p pos) bool {
	return !isWall(m[p.y][p.x])
}

func addPosDirection(d input, p pos) pos {
	r := pos{p.y, p.x}

	switch d {
	case up:
		r.y--
	case right:
		r.x++
	case down:
		r.y++
	case left:
		r.x--
	}

	if r.x < 0 {
		r.x = 0
	}
	if r.y < 0 {
		r.y = 0
	}

	return r
}
