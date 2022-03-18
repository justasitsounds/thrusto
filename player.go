package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const thrust = 0.1
const friction = 0.01

func newPlayer() *element {
	const unit = 8
	np := &element{
		active:   true,
		position: vec{100, 100},
		rotation: -math.Pi / 2,
		label:    "player",
	}
	sd := newScreenDrawer(np, func() *ebiten.Image { return shipImage(unit) })
	np.addComponent(sd)
	np.addComponent(newKeyboardMover(np))
	np.addComponent(newKeyboardShooter(np, time.Millisecond*250))
	return np
}

/*
//deltoid ship
1--+--+--+
+--4--+--2
3--+--+--+
*/
func shipImage(unit float32) *ebiten.Image {
	var path vector.Path
	xf, yf := float32(unit), float32(unit)
	path.MoveTo(0, 0)
	path.LineTo(xf*3, yf)
	path.LineTo(0, yf*2)
	path.LineTo(xf, yf)
	path.LineTo(0, 0)
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
	im := ebiten.NewImage(int(unit)*3, int(unit)*2)
	im.DrawTriangles(vs, is, emptySubImage, op)
	return im
}

const bulletSpeed = 10
