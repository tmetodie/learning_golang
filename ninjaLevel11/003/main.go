package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Custom Error: %v", ce.info)
}

func main() {
	a := customErr{
		info: "Custom Error!",
	}
	log.Fatalln(a)
}
