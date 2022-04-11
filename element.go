package main

import (
	"fmt"
	"math"
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

func (v vec) add(other vec) vec {
	return vec{v.x + other.x, v.y + other.y}
}

func (v vec) mul(other vec) vec {
	return vec{v.x * other.x, v.y * other.y}
}

func (v vec) scale(scalar float64) vec {
	return vec{v.x * scalar, v.y * scalar}
}

func (v vec) sub(other vec) vec {
	return vec{v.x - other.x, v.y - other.y}
}

func (v vec) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v vec) String() string {
	return fmt.Sprintf("{ x : %.2f, y : %.2f }", v.x, v.y)
}

type element struct {
	position      vec
	velocity      vec
	rotation      float64
	active        bool
	label         string
	components    []component
	eventHandlers map[string][]func()
}

func (elem *element) draw(screen *ebiten.Image) error {
	for _, comp := range elem.components {
		if err := comp.ondraw(screen); err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) raiseEvent(event string) {
	for _, handler := range elem.eventHandlers[event] {
		handler()
	}
}

func (elem *element) register(event string, handlers []func()) {
	if elem.eventHandlers == nil {
		elem.eventHandlers = make(map[string][]func())
	}
	for _, handler := range handlers {
		elem.eventHandlers[event] = append(elem.eventHandlers[event], handler)
	}
}

func (elem *element) String() string {
	return fmt.Sprintf("label:%s, position:%v", elem.label, elem.position)
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
