package main

import (
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/fogleman/ease"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
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
		log.Fatal("couldn't create font face")
	}

	emptyImage.Fill(color.White)

	sampleRate := 44100
	audioContext = audio.NewContext(sampleRate)

}

func main() {
	game = &Game{
		position:     vec{0, 0},
		scrollFrames: 15,
		scrollFunc:   ease.Linear,
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
