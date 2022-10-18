package main

import (
	"fmt"
	"image"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type animator struct {
	container        *element
	sequences        map[string]*sequence
	currentSequence  string
	lastFrameChanged time.Time
	finished         bool
	transformation   *ebiten.GeoM
	width, height    float64
	origin           vec
	static           bool
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string, static bool) *animator {
	return &animator{
		container:        container,
		sequences:        sequences,
		currentSequence:  defaultSequence,
		lastFrameChanged: time.Now(),
		static:           static,
	}
}

func (an *animator) onupdate() error {
	sequence := an.sequences[an.currentSequence]

	frameInterval := float64(time.Second) / sequence.sampleRate

	if time.Since(an.lastFrameChanged) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFrameChanged = time.Now()
	}
	return nil
}

func (an *animator) ondraw(screen *ebiten.Image) error {
	img := an.sequences[an.currentSequence].image()
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	origin := vec{-float64(width) / 2, -float64(height) / 2} //default origin center of image
	op := &ebiten.DrawImageOptions{}
	//move the origin
	op.GeoM.Translate(origin.x, origin.y)
	//apply rotation
	op.GeoM.Rotate(an.container.rotation)
	//move to container x,y
	if an.static {
		op.GeoM.Translate(an.container.position.x, an.container.position.y)
	} else {
		op.GeoM.Translate(an.container.position.x-game.position.x, an.container.position.y-game.position.y)
	}

	if an.transformation != nil {
		an.transformation.Concat(op.GeoM)
		screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: *an.transformation})
	} else {
		screen.DrawImage(img, op)
	}
	return nil
}

type screenDrawer struct {
	container *element
	img       *ebiten.Image

	width, height  float64
	origin         vec
	transformation *ebiten.GeoM
	static         bool
}

type sequence struct {
	images     []*ebiten.Image
	frame      int
	sampleRate float64
	loop       bool
}

func (s *sequence) image() *ebiten.Image {
	return s.images[s.frame]
}

func (s *sequence) nextFrame() bool {
	if s.frame == len(s.images)-1 {
		if s.loop {
			s.frame = 0
		} else {
			return true
		}
	} else {
		s.frame++
	}
	return false
}

func newScreenDrawer(container *element, imgfunc func() *ebiten.Image) *screenDrawer {
	img := imgfunc()
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	origin := vec{-float64(width) / 2, -float64(height) / 2} //default origin center of image

	return &screenDrawer{
		container: container,
		img:       img,
		width:     float64(width),
		height:    float64(height),
		origin:    origin,
	}
}

func (s *screenDrawer) withOrigin(origin vec) *screenDrawer {
	s.origin = origin
	return s
}

func (s *screenDrawer) transform(transformation *ebiten.GeoM) error {
	s.transformation = transformation
	return nil
}

func (s *screenDrawer) setWidth(newwidth float64) error {
	return nil
}

func (s *screenDrawer) ondraw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	//move the origin
	op.GeoM.Translate(s.origin.x, s.origin.y)
	//apply rotation
	op.GeoM.Rotate(s.container.rotation)
	//move to container x,y
	if s.static {
		op.GeoM.Translate(s.container.position.x, s.container.position.y)
	} else {
		op.GeoM.Translate(s.container.position.x-game.position.x, s.container.position.y-game.position.y)
	}

	if s.transformation != nil {
		s.transformation.Concat(op.GeoM)
		screen.DrawImage(s.img, &ebiten.DrawImageOptions{GeoM: *s.transformation})
	} else {
		screen.DrawImage(s.img, op)
	}
	return nil
}

func (s *screenDrawer) onupdate() error {
	return nil
}

func newTileImage(mapDef []int, rowlength int) *ebiten.Image {
	tileSrc, _ := imageFS.Open("assets/images/cave_tiles.png")
	defer tileSrc.Close()

	e, _, err := ebitenutil.NewImageFromReader(tileSrc)
	if err != nil {
		panic(fmt.Sprintf("couldn't load tile src:%v", err))
	}

	tileSize := 40
	rowCount := len(mapDef) / rowlength

	width := tileSize * rowlength
	height := tileSize * rowCount

	img := ebiten.NewImage(width, height)
	for i, t := range mapDef {
		op := &ebiten.DrawImageOptions{}
		tilex := float64((i % rowlength) * tileSize)
		tiley := float64((i / rowlength) * tileSize)
		op.GeoM.Translate(tilex, tiley)
		fmt.Printf("tilex:%v, tiley:%v\n", tilex, tiley)

		sx := t
		sy := 0
		fmt.Printf("subImage(x:%v, y:%v) placing at :%v,%v\n", sx, sy, tilex, tiley)
		fmt.Printf("RECT(%v,%v,%v,%v)\n", sx, sy, sx+tileSize, sy+tileSize)
		img.DrawImage(e.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	}
	return img
}
