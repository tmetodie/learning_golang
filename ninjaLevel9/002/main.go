package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("My name is", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Tisho",
	}
	p1.speak()
	//saySomething(p1)
}
