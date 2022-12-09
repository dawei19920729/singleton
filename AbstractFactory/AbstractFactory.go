package AbstractFactory

import "fmt"

type PcFactory interface {
	createMouse() *Mouse
	createKeybo() *Keybo
}
type Mouse interface {
	SayHi()
}
type Keybo interface {
	SayHello()
}
type DelFactory struct {
}

func (df *DelFactory) CreateMouse() Mouse {
	fmt.Println("this is Delmouse create")
	return &DelMouse{}
}
func (df *DelFactory) CreateKeybo() Keybo {
	fmt.Println("this is DelKeybo create")
	return &DelKeybo{}
}

type DelMouse struct {
}
type DelKeybo struct {
}
type HPFactory struct {
}

func (hf *HPFactory) CreateMouse() *HpMouse {
	fmt.Println("this is HpMouse create")

	return &HpMouse{}
}
func (hf *HPFactory) CreateKeybo() *HpKeybo {
	fmt.Println("this is HpKeybo create")
	return &HpKeybo{}
}

type HpMouse struct {
}
type HpKeybo struct {
}

func (dm *DelMouse) SayHi() {
	fmt.Println("this is DelMouse sayHi")
}
func (dm *DelKeybo) SayHello() {
	fmt.Println("this is DelKeybo sayHello")
}
func (hm *HpMouse) SayHi() {
	fmt.Println("this is HpMouse sayHi")
}
func (hm *HpKeybo) SayHello() {
	fmt.Println("this is HpKeybo sayHello")
}
func test() {

}
