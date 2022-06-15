package main

import (
	"fmt"

	"github.com/tmetodie/learning_golang/tree/main/ninjaLevel12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	monti := canine{
		name: "Monti",
		age:  dog.Years(3),
	}

	fmt.Println(monti)
}
