package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type bullet struct {
	x, y     float64
	rotation float64
	active   bool
	image    *ebiten.Image
}

var magazine []*bullet

func initMagazine(bulletCount int) {
	for i := 0; i < bulletCount; i++ {
		bul := newBullet()
		magazine = append(magazine, bul)
	}
}

func bulletFromMagazine() (*bullet, bool) {
	for _, b := range magazine {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}

func bulletImage() *ebiten.Image {
	im := ebiten.NewImage(4, 4)
	im.Fill(color.Black)
	return im
}

func newBullet() *bullet {
	return &bullet{
		image:  bulletImage(),
		active: false,
	}
}

func (b *bullet) update() {
	var vx, vy float64
	vx = math.Cos(b.rotation) * bulletSpeed
	vy = math.Sin(b.rotation) * bulletSpeed
	b.x += vx
	b.y += vy
	if b.x < 0 || b.x > float64(screenwidth) || b.y < 0 || b.y > float64(screenheight) {
		b.active = false
	}
}

func (b *bullet) draw(screen *ebiten.Image) {
	if !b.active {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, op)
}
