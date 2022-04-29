package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type collision struct {
	obj       *resolv.Object
	container *element
}

func (c *collision) ondraw(screen *ebiten.Image) error {
	return nil
}

func (c *collision) onupdate() error {
	c.obj.X = c.container.position.x
	c.obj.Y = c.container.position.y
	return nil
}
