package main

import (
	"fmt"

	cat "github.com/tmetodie/golang_cat"
)

type canine struct {
	name string
	age  int
}

func main() {
	monti := canine{
		name: "Monti",
		age:  cat.Years(3),
	}

	fmt.Println(monti)
}
