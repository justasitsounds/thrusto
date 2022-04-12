package main

import (
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type keyboardMover struct {
	container *element
	animator  *animator
}

var center = vec{x: float64(screenwidth) / 2, y: float64(screenheight) / 2}

func newKeyboardMover(container *element) *keyboardMover {
	km := &keyboardMover{
		container: container,
		animator:  container.getComponent(&animator{}).(*animator),
	}

	return km
}

func (km *keyboardMover) ondraw(screen *ebiten.Image) error {
	return nil
}

func (km *keyboardMover) onupdate() error {

	km.container.velocity.y += gravity

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		km.container.rotation -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		km.container.rotation += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		availableimpulse := burnFuel(thrust)
		if availableimpulse > 0 {
			km.container.raiseEvent("burn")
			km.container.velocity.y += math.Sin(km.container.rotation) * availableimpulse
			km.container.velocity.x += math.Cos(km.container.rotation) * availableimpulse
		} else {
			km.container.raiseEvent("idle")
		}
	} else {
		km.container.raiseEvent("idle")
	}

	km.container.velocity = km.container.velocity.scale(1 - friction)

	km.container.position = km.container.position.add(km.container.velocity)

	screenPos := game.screenPosition(km.container.position)
	log.Printf("ship %v | game %v | center %v | screen %v\n", km.container.position, game.position, center, screenPos)

	if center.sub(screenPos).length() > float64(screenheight)/4 { // if the ship is too far from the center of the screen
		game.scrollTo(km.container.position.sub(center), km.container.velocity.length())
	}

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
			ks.shoot(ks.container.position, ks.container.rotation)
			ks.lastShot = time.Now()
		}
	}
	return nil
}

func (ks *keyboardShooter) shoot(shipPos vec, rotation float64) {
	if b, ok := bulletFromMagazine(); ok {
		b.position = shipPos
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
