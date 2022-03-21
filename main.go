package main

import (
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const gravity = 0.05

var (
	//go:embed assets/fonts
	fonts embed.FS

	gravityRegular font.Face

	elements      []*element
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	screenheight  = 240
	screenwidth   = 320
	screenScale   = 2
)

func init() {
	fontBytes, err := fonts.ReadFile("assets/fonts/GravityRegular5.ttf")
	if err != nil {
		log.Fatal("can't open embedded font")
	}
	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal("couldn't parse font")
	}
	gravityRegular, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    8,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("couldn't create face")
	}

	emptyImage.Fill(color.White)
}

// Game implements ebiten.Game interface.
type Game struct {
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update
	for _, e := range elements {
		e.update()
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.White)
	for _, e := range elements {
		if e.active {
			e.draw(screen)
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenwidth, screenheight
}

func main() {
	game := &Game{}
	//point up
	np := newPlayer()
	elements = append(elements, np)
	elements = append(elements, initMagazine(4)...)
	elements = append(elements, newFuelBar())
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenwidth*screenScale, screenheight*screenScale)
	ebiten.SetWindowTitle("THRUSTO")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
