package main

import (
	"testing"

	"github.com/fogleman/ease"
	"github.com/matryer/is"
)

func TestScrollEasing(t *testing.T) {
	is := is.New(t)
	g := &Game{
		position:     vec{0, 0},
		scrollFrames: 10,
	}
	g.scrollTo(vec{x: 10, y: 0})
	g.scrollFunc = ease.Linear

	// is.Equal(g.xoffset, 1.0)
	is.Equal(g.scrollOffset.x, 10.0)
	// is.Equal(g.xsteps, 10)
	// for i := 0; i < g.scrollSteps; i++ {
	g.Update()
	is.Equal(g.position.x, 1.0)
	g.Update()
	is.Equal(g.position.x, 2.0)
	g.Update()
	g.Update()
	is.Equal(g.position.x, 4.0)
	// }
}
