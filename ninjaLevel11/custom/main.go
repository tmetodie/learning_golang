package main

import "fmt"

func main() {
	err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All is good.")
}

func foo() error {
	return fmt.Errorf("Foo: %w", bar())
}

func bar() error {
	return fmt.Errorf("Bar: Failed on bar()")
}
