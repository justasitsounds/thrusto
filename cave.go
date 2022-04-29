package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

/*
cave 1
M -146.86552 -114.25683
L 146.86552 -114.25683
L 146.86552 114.25682000000002
L 23.00303000000001 114.25682000000002
L 13.90293000000001 80.38422000000003
L 70.52578000000001 28.81698000000003
L 105.40950000000001 -18.705769999999973
L 83.16481000000002 -75.32861999999997
L -37.15874999999998 -79.37310999999997
L -122.59858999999999 -36.90596999999997
L -73.55915999999999 36.40040000000003
L -68.50354999999999 80.38422000000003
L -37.15875999999999 76.84529000000003
L -26.036409999999986 99.59554000000003
L -26.036409999999986 114.25681000000003
L -146.86552999999998 114.25681000000003
*/

func caveImage(unit float32) *ebiten.Image {
	var path vector.Path
	xf, yf := float32(unit), float32(unit)
	path.MoveTo(0, 0)
	path.LineTo(100*xf, 0*yf)
	path.LineTo(100*xf, 40*yf)
	path.LineTo(80*xf, 40*yf)
	path.LineTo(70*xf, 12*yf)
	path.LineTo(40*xf, 12*yf)
	path.LineTo(20*xf, 50*yf)
	path.LineTo(40*xf, 88*yf)
	path.LineTo(70*xf, 88*yf)
	path.LineTo(80*xf, 60*yf)
	path.LineTo(100*xf, 60*yf)
	path.LineTo(100*xf, 100*yf)
	path.LineTo(0, 100*yf)
	path.LineTo(0, 0)
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0x33 / float32(0xff)
		vs[i].ColorG = 0xcc / float32(0xff)
		vs[i].ColorB = 0x99 / float32(0xff)
	}
	im := ebiten.NewImage(100*int(xf), 100*int(xf))
	im.DrawTriangles(vs, is, emptySubImage, op)
	return im
}

func newCave() *element {
	cavImg := caveImage(5)
	cave := &element{
		active:   true,
		position: vec{float64(screenwidth) / 2, float64(screenheight) / 2},
		rotation: -math.Pi / 2,
		label:    "cave",
		width:    cavImg.Bounds().Dx(),
		height:   cavImg.Bounds().Dy(),
	}
	cave.addComponent(newScreenDrawer(cave, func() *ebiten.Image { return cavImg }))

	caveObj := resolv.NewObject(cave.position.x, cave.position.y, float64(cave.width), float64(cave.height), cave.label)
	caveObj.SetShape(resolv.NewConvexPolygon(
		0, 0,
		100*5, 0,
		100*5, 5*40,
		80*5, 40*5,
		70*5, 12*5,
		40*5, 12*5,
		20*5, 50*5,
		40*5, 88*5,
		70*5, 88*5,
		80*5, 60*5,
		100*5, 60*5,
		100*5, 100*5,
		0, 100*5,
	))
	caveCollision := &collision{
		obj:       caveObj,
		container: cave,
	}

	cave.addComponent(caveCollision)
	cave.collisions = append(cave.collisions, caveCollision)
	return cave
}
