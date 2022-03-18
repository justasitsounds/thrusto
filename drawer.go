package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type screenDrawer struct {
	container *element
	img       *ebiten.Image

	width, height  float64
	origin         vec
	transformation *ebiten.GeoM
}

func newScreenDrawer(container *element, imgfunc func() *ebiten.Image) *screenDrawer {
	img := imgfunc()
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	origin := vec{-float64(width) / 2, -float64(height) / 2} //default origin center of image

	log.Printf("newScreenDrawer: container:%s, width:%d, height:%d, origin: %v", container, width, height, origin)
	return &screenDrawer{
		container: container,
		img:       img,
		width:     float64(width),
		height:    float64(height),
		origin:    origin,
	}
}

func (s *screenDrawer) withOrigin(origin vec) *screenDrawer {
	s.origin = origin
	log.Printf("newScreenDrawer: container:%s, width:%f, height:%f, origin: %v", s.container, s.width, s.height, origin)
	return s
}

func (s *screenDrawer) transform(transformation *ebiten.GeoM) error {
	s.transformation = transformation
	return nil
}

func (s *screenDrawer) setWidth(newwidth float64) error {
	return nil
}

func (s *screenDrawer) ondraw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	//move the origin
	op.GeoM.Translate(s.origin.x, s.origin.y)
	//apply rotation
	op.GeoM.Rotate(s.container.rotation)
	//move to container x,y
	op.GeoM.Translate(s.container.position.x, s.container.position.y)

	if s.transformation != nil {
		s.transformation.Concat(op.GeoM)
		screen.DrawImage(s.img, &ebiten.DrawImageOptions{GeoM: *s.transformation})
	} else {
		screen.DrawImage(s.img, op)
	}
	return nil
}

func (s *screenDrawer) onupdate() error {
	return nil
}
