package pacman

import (
	"bytes"
	"github.com/weverton-souza/go-pacman/pacman/handler"
	"image"

	"github.com/hajimehoshi/ebiten"
	pacimages "github.com/weverton-souza/go-pacman/images"
)

type scene struct {
	matrix        [][]elem
	wallSurface   *ebiten.Image
	images        map[elem]*ebiten.Image
	stage         *stage
	dotManager    *dotManager
	bigDotManager *bigDotManager
	player        *player
	ghostManager  *ghostManager
	textManager   *textManager
}

func newScene(st *stage) (s *scene) {
	s = &scene{}
	s.stage = st
	if s.stage == nil {
		s.stage = defaultStage
	}
	s.images = make(map[elem]*ebiten.Image)
	s.dotManager = newDotManager()
	s.bigDotManager = newBigDotManager()
	s.ghostManager = newGhostManager()
	s.textManager = newTextManager(len(s.stage.matrix[0])*stageBlocSize, len(s.stage.matrix)*stageBlocSize)
	s.loadImages()
	s.createStage()
	s.buildWallSurface()
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

			switch s.matrix[i][j] {
			case dotElem:
				s.dotManager.add(i, j)
			case bigDotElem:
				s.bigDotManager.add(i, j)
			case playerElem:
				s.player = newPlayer(i, j)
			case blinkyElem:
				s.ghostManager.addGhost(i, j, blinkyElem)
			case inkyElem:
				s.ghostManager.addGhost(i, j, inkyElem)
			case pinkyElem:
				s.ghostManager.addGhost(i, j, pinkyElem)
			case clydeElem:
				s.ghostManager.addGhost(i, j, clydeElem)
			}
		}
	}
}

func (s *scene) screenWidth() (w int) {
	w = len(s.stage.matrix[0]) * stageBlocSize
	return
}

func (s *scene) screenHeight() (h int) {
	h = ((len(s.stage.matrix)*stageBlocSize)/backgroundImageSize + 2) * backgroundImageSize
	return
}

func (s *scene) buildWallSurface() {
	height := len(s.stage.matrix)
	width := len(s.stage.matrix[0])

	sizeW := ((width*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH := ((height*stageBlocSize)/backgroundImageSize + 1) * backgroundImageSize
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

	for i := 0; i < height; i++ {
		y := float64(i * stageBlocSize)
		for j := 0; j < width; j++ {
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

/*
* When IsDrawingSkipped is true, the rendered result is not adopted.
*
 */
func (s *scene) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err1 := screen.Clear()
	if err1 != nil {
		return err1
	}

	err2 := screen.DrawImage(s.wallSurface, nil)
	if err2 != nil {
		return err2
	}

	s.dotManager.draw(screen)
	s.bigDotManager.draw(screen)
	s.player.draw(screen)
	s.ghostManager.draw(screen)
	s.textManager.draw(screen, 0, 1, s.player.images[1])

	return nil
}

func (s *scene) loadImages() {
	for i := w0; i <= w24; i++ {
		s.images[i] = loadImage(pacimages.WallImages[i])
	}
	s.images[backgroundElem] = loadImage(pacimages.Background_png)
}
