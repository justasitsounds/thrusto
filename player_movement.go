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
	km.container.velocity.y += gravity

	km.container.velocity.x *= (1 - friction)
	km.container.velocity.y *= (1 - friction)

	km.container.position.x += km.container.velocity.x
	km.container.position.y += km.container.velocity.y

	screenPos := screenPosition(km.container.position, game.position)
	log.Printf("ship %v | game %v | center %v | screen %v\n", km.container.position, game.position, center, screenPos)

	if center.sub(screenPos).length() > float64(screenheight)/4 {
		game.scrollTo(km.container.position.sub(center), km.container.velocity.length())
	}

	return nil
}

func screenPosition(shipPos vec, gamePos vec) vec {
	return vec{
		x: shipPos.x - gamePos.x,
		y: shipPos.y - gamePos.y,
	}
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
