package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var magazine []*element

func initMagazine(bulletCount int) []*element {
	for i := 0; i < bulletCount; i++ {
		bul := newBullet()
		magazine = append(magazine, bul)
	}
	return magazine
}

func bulletFromMagazine() (*element, bool) {
	for _, b := range magazine {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}

func bulletImage() *ebiten.Image {
	im := ebiten.NewImage(3, 3)
	im.Fill(color.Black)
	return im
}

func newBullet() *element {
	el := &element{}
	el.addComponent(newScreenDrawer(el, bulletImage))
	el.addComponent(newBulletMover(el))

	return el
}

type bulletMover struct {
	container *element
	velocity  vec
}

func newBulletMover(container *element) *bulletMover {
	return &bulletMover{
		container: container,
		velocity:  vec{0, 0},
	}
}
func (bm *bulletMover) onupdate() error {
	// var vx, vy float64
	bm.velocity.x = math.Cos(bm.container.rotation) * bulletSpeed
	bm.velocity.y = math.Sin(bm.container.rotation) * bulletSpeed
	bm.container.position.x += bm.velocity.x
	bm.container.position.y += bm.velocity.y
	if bm.container.position.x < 0 || bm.container.position.x > float64(screenwidth) || bm.container.position.y < 0 || bm.container.position.y > float64(screenheight) {
		bm.container.active = false
	}
	return nil
}

func (bm *bulletMover) ondraw(screen *ebiten.Image) error {
	return nil
}
