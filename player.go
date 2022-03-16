package main

import (
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	tmath "github.com/justasitsounds/thrusto/math"
)

//Player struct for the player
type Player struct {
	x, y          float64
	height, width float64
	rotation      float64
	vx            float64
	vy            float64
	dead          bool
	shipImage     *ebiten.Image
	lastShot      time.Time

	fuel int64
}

//NewPlayer creates a new player
func NewPlayer() *Player {
	const unit = 8

	im := shipImage(unit)

	return &Player{
		x:      float64(screenwidth) / 2,
		y:      float64(screenheight) / 2,
		vx:     0,
		vy:     0,
		height: float64(im.Bounds().Dy()),
		width:  float64(im.Bounds().Dx()),

		rotation:  -math.Pi / 2,
		shipImage: im,
	}
}

func shipImage(unit float32) *ebiten.Image {
	var path vector.Path
	xf, yf := float32(unit), float32(unit)
	path.MoveTo(xf, yf)
	path.LineTo(xf+unit, yf-unit)
	path.LineTo(xf, yf+2*unit)
	path.LineTo(xf-unit, yf-unit)
	path.LineTo(xf, yf)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x11 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}
	im := ebiten.NewImage(int(unit)*2, int(unit)*3)
	im.DrawTriangles(vs, is, emptySubImage, op)
	return im
}

const gravity = 0.05
const thrust = 0.1
const friction = 0.01

func (p *Player) update() error {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.vy += math.Sin(p.rotation) * thrust
		p.vx += math.Cos(p.rotation) * thrust
	}
	p.vy += gravity
	p.vx *= (1 - friction)
	p.vy *= (1 - friction)
	p.x += p.vx
	p.y += p.vy

	p.x = tmath.Clampf(p.x, 0, float64(screenwidth))
	p.y = tmath.Clampf(p.y, 0, float64(screenheight))

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(p.lastShot) > time.Millisecond*200 {
			p.shoot(p.x, p.y, p.rotation)
			p.lastShot = time.Now()
		}
	}

	return nil

}

func (p *Player) draw(screen *ebiten.Image) {
	// Draw a ship.
	op := &ebiten.DrawImageOptions{}
	//move the origin to center of the image
	op.GeoM.Translate(-float64(p.shipImage.Bounds().Dx())/2, -float64(p.shipImage.Bounds().Dy())/2)
	//apply rotation
	op.GeoM.Rotate(p.rotation - math.Pi/2)
	//place ship
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.shipImage, op)
}

const bulletSpeed = 10

func (p *Player) shoot(x, y, angle float64) {
	log.Println("shooting")
	if b, ok := bulletFromMagazine(); ok {
		b.x = x
		b.y = y
		b.rotation = angle
		b.active = true
	} else {
		log.Println("no bullets")
	}
}
