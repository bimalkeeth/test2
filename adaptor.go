package main

import "fmt"

type Adaptee struct {
}

func (a *Adaptee) ExistingMethod() {
	fmt.Println("Using existing method")
}

type Adaptor struct {
	adaptee *Adaptee
}

func NewAdaptor() *Adaptor {
	return &Adaptor{new(Adaptee)}
}

func (a *Adaptor) ExistingMethod() {
	fmt.Println("Doing some work")
	a.adaptee.ExistingMethod()
}

func main() {

	adaptor := NewAdaptor()
	adaptor.ExistingMethod()

}
