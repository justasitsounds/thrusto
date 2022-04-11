package main

import (
	"testing"

	"github.com/fogleman/ease"
	"github.com/matryer/is"
)

func TestScrollEasing(t *testing.T) {
	is := is.NewRelaxed(t)
	g := &Game{
		position:     vec{0, 0},
		scrollFrames: 4,
	}
	g.scrollTo(vec{x: 4, y: 0}, 3.0)
	g.scrollFunc = ease.Linear

	is.Equal(g.scrollOffset.x, 4.0)

	g.Update()
	is.Equal(g.position.x, 1.0)
	g.Update()
	is.Equal(g.position.x, 2.0)
	g.Update()
	g.Update()
	is.Equal(g.position.x, 4.0)

	g.scrollTo(vec{x: 8, y: 0}, 3.0)

	is.Equal(g.scrollOffset.x, 4.0)
	g.Update()
	is.Equal(g.position.x, 5.0)
	g.Update()
	g.Update()
	g.Update()
	is.Equal(g.position.x, 8.0)
}
