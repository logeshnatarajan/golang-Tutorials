package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("donuuu......")
	fmt.Println("goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("after goroutines:", runtime.NumGoroutine())

}
func foo() {

	fmt.Println("from fooo ....")
	wg.Done()

}
func bar() {
	fmt.Println("from bar........")
	wg.Done()

}
