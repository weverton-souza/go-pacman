package pacman

import (
	"bytes"
	"github.com/weverton-souza/go-pacman/pacman/handler"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/weverton-souza/go-pacman/images"
)

type scene struct {
	matrix      [][]elem
	wallSurface *ebiten.Image
	images      map[elem]*ebiten.Image
	stage       *stage
}

func newScene(st *stage) (s *scene) {
	s = &scene{}
	s.stage = st
	if s.stage == nil {
		s.stage = defaultStage
	}
	s.images = make(map[elem]*ebiten.Image)
	s.loadImage()
	s.createStage()
	s.buildWallSurface()
	return
}

func (s *scene) createStage() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])
	s.matrix = make([][]elem, h)

	for i := 0; i < h; i++ {
		s.matrix[i] = make([]elem, w)

		for j := 0; j < w; j++ {
			c := s.stage.matrix[i][j] - '0'
			if c <= 9 {
				s.matrix[i][j] = elem(c)
			} else {
				s.matrix[i][j] = elem(s.stage.matrix[i][j] - 'a' + 10)
			}
		}
	}
}

func (s *scene) screenWidth() (w int) {
	w = len(s.stage.matrix[0]) * stageBlocSize
	return
}

func (s *scene) screenHeight() (h int) {
	h = len(s.stage.matrix) * stageBlocSize
	return
}

func (s *scene) buildWallSurface() {
	h := len(s.stage.matrix)
	w := len(s.stage.matrix[0])

	sizeW := ((w*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH := ((h*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	s.wallSurface, _ = ebiten.NewImage(sizeW, sizeH, ebiten.FilterDefault)

	for i := 0; i < sizeH/backgroundImageSize; i++ {
		y := float64(i * backgroundImageSize)

		for j := 0; j < sizeW/backgroundImageSize; j++ {
			op := &ebiten.DrawImageOptions{}
			x := float64(j * backgroundImageSize)
			op.GeoM.Translate(x, y)

			err := s.wallSurface.DrawImage(s.images[backgroundElem], op)
			handler.HandleError(handler.RUNTIME, err)
		}
	}

	for i := 0; i < h; i++ {
		y := float64(i * stageBlocSize)
		for j := 0; j < w; j++ {
			if isWall(s.matrix[i][j]) {
				op := &ebiten.DrawImageOptions{}
				x := float64(j * stageBlocSize)
				op.GeoM.Translate(x, y)

				err := s.wallSurface.DrawImage(s.images[s.matrix[i][j]], op)
				handler.HandleError(handler.RUNTIME, err)
			}
		}
	}
}

func (s *scene) loadImage() {
	for i := w0; i <= w24; i++ {
		img, _, err := image.Decode(bytes.NewReader(pacimages.WallImages[i]))
		handler.HandleError(handler.RUNTIME, err)

		s.images[i], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		handler.HandleError(handler.RUNTIME, err)
	}

	img, _, err := image.Decode(bytes.NewReader(pacimages.Background_png))
	handler.HandleError(handler.RUNTIME, err)

	s.images[backgroundElem], err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	handler.HandleError(handler.RUNTIME, err)
}

func (s *scene) update(screen *ebiten.Image) error {
	if err1, err2 := screen.Clear(), screen.DrawImage(s.wallSurface, nil); !ebiten.IsDrawingSkipped() && (err1 != nil || err2 != nil) {

		handler.HandleError(handler.RUNTIME, err1, err2)
		return nil
	}

	return nil
}
