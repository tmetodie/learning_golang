package main

import "fmt"

func main() {
	foo(bar)
}

func foo(f func() int) {
	fmt.Println(f())
}

func bar() int {
	return 42
}
