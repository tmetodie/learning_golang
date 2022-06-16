package main

import "fmt"

func main() {
	fmt.Println("5 + 8  =", mySum(5, 8))
	fmt.Println("3 + 6  =", mySum(3, 6))
	fmt.Println("12 + 4 =", mySum(12, 4))
}

func mySum(a ...int) int {
	var sum int
	for _, i := range a {
		sum += i
	}
	return sum
}
