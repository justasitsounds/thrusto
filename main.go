package main

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const gravity = 0.05

var (
	//go:embed assets/fonts
	fonts embed.FS
	//go:embed assets/audio
	audioFS embed.FS

	gravityRegular font.Face
	audioContext   *audio.Context
	elements       []*element

	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	screenheight  = 240
	screenwidth   = 320
	screenScale   = 2
	game          *Game
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

	sampleRate := 44100
	audioContext = audio.NewContext(sampleRate)

}

// Game implements ebiten.Game interface.
type Game struct {
	position vec
	xsteps   int
	xoffset  float64
	ysteps   int
	yoffset  float64
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update
	if g.xsteps > 0 {
		g.position.x += g.xoffset
		g.xsteps--
	}

	if g.ysteps > 0 {
		g.position.y += g.yoffset
		g.ysteps--
	}

	// g.position.x += 1

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
	ebitenutil.DebugPrint(screen, fmt.Sprintf("GamePosition: %v", g.position))
	ebitenutil.DrawRect(screen, float64(screenwidth)/4, float64(screenheight)/4, float64(screenwidth)/2, float64(screenheight)/2, color.RGBA{0xff, 0x00, 0x00, 0x80})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenwidth, screenheight
}

func (g *Game) scrollX(xoffset float64) {
	var steps = 240
	g.xoffset = xoffset / float64(steps)
	g.xsteps = steps
}

func (g *Game) scrollY(yoffset float64) {
	var steps = 240
	g.yoffset = yoffset / float64(steps)
	g.ysteps = steps
}

func main() {
	game = &Game{
		position: vec{0, 0},
	}

	cave := newCave()

	np := newPlayer(vec{float64(screenwidth) / 2, float64(screenheight) / 2})

	elements = append(elements, cave)
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
