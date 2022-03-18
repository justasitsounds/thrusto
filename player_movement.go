package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	tmath "github.com/justasitsounds/thrusto/math"
)

type keyboardMover struct {
	container *element
}

func newKeyboardMover(container *element) *keyboardMover {
	return &keyboardMover{
		container: container,
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
		availableimpulse := burnFuel(thrust)
		km.container.velocity.y += math.Sin(km.container.rotation) * availableimpulse
		km.container.velocity.x += math.Cos(km.container.rotation) * availableimpulse
	}
	km.container.velocity.y += gravity

	km.container.velocity.x *= (1 - friction)
	km.container.velocity.y *= (1 - friction)

	km.container.position.x += km.container.velocity.x
	km.container.position.y += km.container.velocity.y

	//keep ship on screen - would be useful for scrolling bounds?
	km.container.position.x = tmath.Clampf(km.container.position.x, 0, float64(screenwidth))
	km.container.position.y = tmath.Clampf(km.container.position.y, 0, float64(screenheight))

	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func (ks *keyboardShooter) ondraw(screen *ebiten.Image) error {
	return nil
}

func (ks *keyboardShooter) onupdate() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if time.Since(ks.lastShot) > ks.cooldown {
			ks.shoot(ks.container.position.x, ks.container.position.y, ks.container.rotation)
			ks.lastShot = time.Now()
		}
	}
	return nil
}

func (ks *keyboardShooter) shoot(x, y, rotation float64) {
	if b, ok := bulletFromMagazine(); ok {
		b.position.x = x
		b.position.y = y
		b.rotation = rotation
		b.active = true
	}
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
		lastShot:  time.Now(),
	}
}
