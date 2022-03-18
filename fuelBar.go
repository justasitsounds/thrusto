package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var fuellevel float64 = 10

func newFuelBar() *element {
	el := &element{
		position: vec{25, 8},
		active:   true,
		label:    "fuelbar",
	}
	sd := newScreenDrawer(el, fuelBarImage).withOrigin(vec{0, 0})
	el.addComponent(sd)
	el.addComponent(newFuelBarMover(el, fuellevel))

	return el
}

func fuelBarImage() *ebiten.Image {
	im := ebiten.NewImage(screenwidth-50, 4)
	im.Fill(color.RGBA{0, 128, 0, 255})
	return im
}

func burnFuel(amount float64) float64 {
	if fuellevel > amount {
		fuellevel -= amount
		return amount
	}
	return 0
}

type fuelBarMover struct {
	container *element
	drawer    *screenDrawer
	maxfuel   float64
}

func newFuelBarMover(container *element, maxfuel float64) *fuelBarMover {
	return &fuelBarMover{
		container: container,
		drawer:    container.getComponent(&screenDrawer{}).(*screenDrawer),
		maxfuel:   maxfuel,
	}

}

func (fm *fuelBarMover) onupdate() error {
	t := &ebiten.GeoM{}
	t.Scale(fuellevel/fm.maxfuel, 1)
	fm.drawer.transform(t)
	return nil
}

func (fm *fuelBarMover) ondraw(screen *ebiten.Image) error {
	return nil
}
