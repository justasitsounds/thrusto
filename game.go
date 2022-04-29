package main

import (
	"image/color"
	"log"
	"math"

	"github.com/fogleman/ease"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"

	jmath "github.com/justasitsounds/thrusto/math"
)

// Game implements ebiten.Game interface.
type Game struct {
	position     vec
	scrollOffset vec
	scrollFrom   vec
	scrollFrames int
	scrollFunc   ease.Function
	scrollSteps  int
	scrolling    bool
	space        *resolv.Space
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if g.scrollSteps < g.scrollFrames { // we scrolling
		inc := g.scrollOffset.scale(g.scrollFunc(float64(g.scrollSteps) / float64(g.scrollFrames))) // divide distance by number of frames to animate over - using easing function to shape
		g.position = g.scrollFrom.add(inc)
		g.scrollSteps++
	}

	for _, e := range elements {
		if e.active {
			e.update()
		}
	}
	checkCollisions()
	return nil
}

func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if intersection := c1.obj.Shape.Intersection(c1.container.velocity.x, c1.container.velocity.y, c2.obj.Shape); intersection != nil {
						log.Println("Collision!")
						log.Println(intersection.Center)
					}
				}
			}
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x33, 0x33, 0x66, 0xff})
	for _, e := range elements {
		if e.active {
			e.draw(screen)
		}
	}
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("ship position: %v", g.position))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenwidth, screenheight
}

func (g *Game) scrollTo(target vec, speed float64) { // speed - not velocity - how technically correct is this?
	/* longest animation is 30 frames, shortest is 1 frame
	the faster the speed, the faster the animation */
	g.scrollFrames = jmath.Clamp(30-int(math.Pow(speed, 2)), 0, 30) + 1
	g.scrollFrom = g.position
	g.scrollOffset = target.sub(g.position)
	g.scrollSteps = 1
}

func (g *Game) screenPosition(elementPosition vec) vec {
	return elementPosition.sub(g.position)
}
