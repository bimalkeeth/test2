package main

import "fmt"

type interfaceX interface {
	MethodA()
	AddChild(interfaceX)
}
type Composite struct {
	children []interfaceX
}

func (c *Composite) MethodA() {
	if len(c.children) == 0 {
		fmt.Println("I ama leaf")
		return
	}
	fmt.Println("I am a composite")
	for _, child := range c.children {
		child.MethodA()
	}

}

func (c *Composite) AddChild(child interfaceX) {
	c.children = append(c.children, child)
}

func main() {
	var parent interfaceX
	parent = &Composite{}
	parent.MethodA()

	var child Composite
	var child2 Composite
	parent.AddChild(&child)
	parent.AddChild(&child2)
	var child3 Composite
	child.AddChild(&child3)

	parent.MethodA()

}
