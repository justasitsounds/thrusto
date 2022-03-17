package main

import "github.com/hajimehoshi/ebiten/v2"

type screenDrawer struct {
	container *element
	img       *ebiten.Image

	width, height float64
}

func newScreenDrawer(container *element, imgfunc func() *ebiten.Image) *screenDrawer {
	img := imgfunc()
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y

	return &screenDrawer{
		container,
		img,
		float64(width),
		float64(height),
	}
}

func (s *screenDrawer) ondraw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	//move the origin to center of the image
	op.GeoM.Translate(-s.width/2, -s.height/2)
	//apply rotation
	op.GeoM.Rotate(s.container.rotation)
	//move to container x,y
	op.GeoM.Translate(s.container.position.x, s.container.position.y)

	screen.DrawImage(s.img, op)
	return nil
}

func (s *screenDrawer) onupdate() error {
	return nil
}
