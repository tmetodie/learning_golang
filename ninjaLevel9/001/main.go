package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("First:", i)
		}
		wg.Done()
	}()

	go func(x int) {
		for i := 0; i <= x; i++ {
			fmt.Println("Second:", i)
		}
		wg.Done()
	}(5)

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
