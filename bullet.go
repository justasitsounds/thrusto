package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var magazine []*element

const bulletSpeed = 10

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
	im.Fill(color.White)
	return im
}

func newBullet() *element {
	el := &element{
		label: "bullet",
	}
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
	bm.velocity.x = math.Cos(bm.container.rotation) * bulletSpeed
	bm.velocity.y = math.Sin(bm.container.rotation) * bulletSpeed
	bm.container.position = bm.container.position.add(bm.velocity)

	screenPos := game.screenPosition(bm.container.position)

	bm.container.active = center.sub(screenPos).length() < float64(screenwidth) // if the bullet is too far fom the center of the screen make it inactive

	return nil
}

func (bm *bulletMover) ondraw(screen *ebiten.Image) error {
	return nil
}
