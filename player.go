package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const thrust = 0.1
const friction = 0.01

func newPlayer(startpos vec) *element {
	const unit = 8
	player := &element{
		active:   true,
		position: startpos,
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
	idleImage := shipImage(unit)
	sequences := map[string]*sequence{
		"idle": {
			images: []*ebiten.Image{idleImage},
			loop:   false,
		},
		"burn": seq,
	}
	an := newAnimator(player, sequences, "idle", false)
	player.addComponent(an)

	player.width = idleImage.Bounds().Dx()
	player.height = idleImage.Bounds().Dy()

	player.on("burn", func() { an.currentSequence = "burn" })
	player.on("idle", func() { an.currentSequence = "idle" })

	player.addComponent(newKeyboardMover(player))

	thrustSound := newSound("assets/audio/thrust_loop.ogg", true)
	player.on("burn", thrustSound.play)
	player.on("idle", thrustSound.stop)

	shotsound := newSound("assets/audio/shot.ogg", false)
	player.on("shoot", func() {
		shotsound.player.Rewind()
		shotsound.play()
	})

	player.addComponent(newKeyboardShooter(player, time.Millisecond*250))

	shipObj := resolv.NewObject(player.position.x, player.position.y, float64(player.width), float64(player.height), player.label)
	shipObj.SetShape(resolv.NewRectangle(0, 0, float64(player.width), float64(player.height)))
	shipCollision := &collision{
		obj:       shipObj,
		container: player,
	}

	player.addComponent(shipCollision)

	player.collisions = append(player.collisions, shipCollision)

	return player
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
		vs[i].ColorR = 1
		vs[i].ColorG = 0.8 - (rand.Float32() * 4 / 10) //0.6 +-0.2
		vs[i].ColorB = 0.3 - (rand.Float32() * 2 / 10) //0.2+-0.1
	}

	im := shipImage(unit)
	im.DrawTriangles(vs, is, emptySubImage, op)
	return im
}
