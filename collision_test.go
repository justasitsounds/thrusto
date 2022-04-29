package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestCollision(t *testing.T) {
	is := is.New(t)
	c := circle{
		center: vec{x: 10, y: 10},
		radius: 3.0,
	}

	/*
		cave := []vec{
			vec{x: 0, y: 0},
			vec{x: 20, y: 0},
			vec{x: 20, y: 20},
			vec{x: 4, y: 30},
			vec{x: 4, y: 60},
			vec{x: 30, y: 70},
			vec{x: 56, y: 60},
			vec{x: 56, y: 30},
			vec{x: 40, y: 20},
			vec{x: 40, y: 0},
			vec{x: 60, y: 0},
			vec{x: 60, y: 80},
			vec{x: 0, y: 80},
		}

		interiorVertices := cave[2 : len(cave)-2]

		is.True(!c.collidesWith(interiorVertices))
	*/
	is.True(!c.collidesWith([]vec{{x: 0, y: 8}, {x: 20, y: 20}}))
}
