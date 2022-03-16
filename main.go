package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	screenheight  = 480
	screenwidth   = 640
)

func init() {
	emptyImage.Fill(color.White)
}

// Game implements ebiten.Game interface.
type Game struct {
	player *Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update
	g.player.update()
	for _, bullet := range magazine {
		bullet.update()
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.White)
	g.player.draw(screen)
	for _, bullet := range magazine {
		bullet.draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenwidth, screenheight
}

func main() {
	game := &Game{
		player: NewPlayer(),
	}
	initMagazine(4)
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenwidth, screenheight)
	ebiten.SetWindowTitle("THRUSTO")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
