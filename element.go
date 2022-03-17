package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type component interface {
	ondraw(screen *ebiten.Image) error
	onupdate() error
}

type vec struct {
	x, y float64
}

type element struct {
	position   vec
	velocity   vec
	rotation   float64
	active     bool
	components []component
}

func (elem *element) draw(screen *ebiten.Image) error {
	log.Println("drawing element")
	for _, comp := range elem.components {

		if err := comp.ondraw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		if err := comp.onupdate(); err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) addComponent(comp component) {
	for _, existing := range elem.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(comp)))
		}
	}
	elem.components = append(elem.components, comp)
}

func (elem *element) getComponent(query component) component {
	typ := reflect.TypeOf(query)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf("no component of type %v found", typ))
}
