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

	km.container.position.x += km.container.velocity.x //- game.position.x
	km.container.position.y += km.container.velocity.y //- game.position.y

	log.Printf("ship pos %v\n", km.container.position)
	if km.container.position.x < float64(screenwidth)/4 {
		game.scrollX(float64(screenwidth)/2 - km.container.position.x)
		// game.scrollX(-float64(screenwidth) / 4)
		// game.position.x += float64(screenwidth) / 4
	}
	if km.container.position.x > 3*(float64(screenwidth)/4) {
		// game.scrollX(float64(screenwidth) / 4)
		// game.position.x -= float64(screenwidth) / 4
		game.scrollX(km.container.position.x - float64(screenwidth)/2)
	}
	if km.container.position.y < float64(screenheight)/4 {
		game.scrollY(-float64(screenheight)/2 - km.container.position.y)
		// game.posit-ion.y += float64(screenhight) / 4
	}
	if km.container.position.y > 3*(float64(screenheight)/4) {
		game.scrollY(km.container.position.y - float64(screenheight)/2)
		// game.positon.y -= float64(screenheight) / 4

	}

	//keep ship on screen - would be useful for scrolling bounds?
	// km.container.position.x = tmath.Clampf(km.container.position.x, 0, float64(screenwidth))
	// km.container.position.y = tmath.Clampf(km.container.position.y, 0, float64(screenheight))

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
