package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

//Player holds player state
type Player struct {
	x, y      float64
	rotation  float64
	vx        float64
	vy        float64
	dead      bool
	shipImage *ebiten.Image
}

func NewPlayer() *Player {
	const unit = 8

	var path vector.Path
	xf, yf := float32(unit), float32(unit*2)
	path.MoveTo(xf, yf)
	path.LineTo(xf-unit, yf+unit)
	path.LineTo(xf, yf-2*unit)
	path.LineTo(xf+unit, yf+unit)
	path.LineTo(xf, yf)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}
	im := ebiten.NewImage(unit*2, unit*3)
	im.DrawTriangles(vs, is, emptySubImage, op)
	om := ebiten.NewImage(unit*2, unit*3)
	om.Bounds()

	return &Player{
		x:  float64(screenwidth) / 2,
		y:  float64(screenheight) / 2,
		vx: 0,
		vy: 0,

		shipImage: im,
	}
}

const gravity = 0.02
const thrust = 0.1

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update
	return g.player.update()
}

func (p *Player) update() error {

	if p.y > float64(screenheight) || p.y < 0 {
		p.dead = true
		p.y = float64(screenheight) / 2
		return nil
	}
	if p.x > float64(screenwidth) || p.x < 0 {
		p.dead = true
		p.x = float64(screenwidth) / 2
		return nil

	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.vy += math.Sin(p.rotation-math.Pi/2) * thrust
		p.vx += math.Cos(p.rotation-math.Pi/2) * thrust
		log.Printf("vy: %f, vx: %f", p.vy, p.vx)
	}
	p.vy += gravity

	p.x += p.vx
	p.y += p.vy

	return nil

}

func (p *Player) draw(screen *ebiten.Image) {
	// Draw a ship.
	op := &ebiten.DrawImageOptions{}
	//move the origin to center of the image
	op.GeoM.Translate(-float64(p.shipImage.Bounds().Dx())/2, -float64(p.shipImage.Bounds().Dy())/2)
	//apply rotation
	op.GeoM.Rotate(p.rotation)
	//place ship
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.shipImage, op)
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.White)
	g.player.draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{
		player: NewPlayer(),
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
