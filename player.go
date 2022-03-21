package main

import (
	"math"
	"math/rand"
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
	seq := &sequence{
		images: []*ebiten.Image{
			shipWithFlame(unit, 1),
			shipWithFlame(unit, 3),
			shipWithFlame(unit, 5),
			shipWithFlame(unit, 2),
		},
		loop:       true,
		sampleRate: 30,
	}
	sequences := map[string]*sequence{
		"idle": {
			images: []*ebiten.Image{shipImage(unit)},
			loop:   false,
		},
		"burn": seq,
	}
	an := newAnimator(np, sequences, "burn")
	np.addComponent(an)
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

/*
+--+--2--+--+
+--+--+--+--+
1--+--+--+--3
+--+--+--+--+
+--+--4--+--+
*/
func shipWithFlame(unit float32, seed int64) *ebiten.Image {
	var path vector.Path
	xf, yf := float32(unit/2), float32(unit/2)

	rand.Seed(seed * 52413)
	initx := rand.Intn(int(unit))

	path.MoveTo(float32(initx), yf*2)
	path.LineTo(xf, yf)
	path.LineTo(xf*2, yf*2)
	path.LineTo(xf, yf*3)
	path.LineTo(float32(initx), yf*2)
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xff / float32(0xff)
		vs[i].ColorG = 0x99 / float32(0xff)
		vs[i].ColorB = 0x33 / float32(0xff)
	}

	im := shipImage(unit)
	im.DrawTriangles(vs, is, emptySubImage, op)
	return im
}
