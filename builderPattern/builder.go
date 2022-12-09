package main

import "fmt"

type Bus struct {
	body  string
	wheel string
}

func (b *Bus) SetBody(body string) {
	b.body = body
}
func (b *Bus) SetWheel(wheel string) {
	b.wheel = wheel
}

type Builder interface {
	CreateBody(body string)
	CreateWheel(wheel string)
}
type BusBuilder struct {
	bus *Bus
}

func (Bbuilder *BusBuilder) CreateBody(body string) {
	Bbuilder.bus.SetBody(body)
}
func (Bbuilder *BusBuilder) CreateWheel(wheel string) {
	Bbuilder.bus.SetWheel(wheel)
}

type Director1 struct {
	builder Builder
}

func (d *Director1) Produce(body, wheel string) {
	if d == nil {
		return
	}
	d.builder.CreateWheel(wheel)
	d.builder.CreateBody(body)
}
func main() {
	bus := Bus{}
	busBuilder := &BusBuilder{&bus}
	director := Director1{busBuilder}
	director.Produce("bigBusBody", "bigBusWheel")
	fmt.Println(bus)
	fmt.Println(busBuilder.bus)
}
