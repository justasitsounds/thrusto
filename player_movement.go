package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	tmath "github.com/justasitsounds/thrusto/math"
)

type keyboardMover struct {
	container *element
	drawer    *screenDrawer
}

func newKeyboardMover(container *element) *keyboardMover {
	return &keyboardMover{
		container: container,
		drawer:    container.getComponent(&screenDrawer{}).(*screenDrawer),
	}
}

func (km *keyboardMover) ondraw(screen *ebiten.Image) error {
	return nil
}

func (km *keyboardMover) onupdate() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		km.container.rotation -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		km.container.rotation += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		km.container.velocity.y += math.Sin(km.container.rotation) * thrust
		km.container.velocity.x += math.Cos(km.container.rotation) * thrust
	}
	km.container.velocity.y += gravity

	km.container.velocity.x *= (1 - friction)
	km.container.velocity.y *= (1 - friction)

	km.container.position.x += km.container.velocity.x
	km.container.position.y += km.container.velocity.y

	km.container.position.x = tmath.Clampf(km.container.position.x, 0, float64(screenwidth))
	km.container.position.y = tmath.Clampf(km.container.position.y, 0, float64(screenheight))

	return nil
}
