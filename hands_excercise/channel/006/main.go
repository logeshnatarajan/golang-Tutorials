package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			for j := 0; j < 10; j++ {
				c <- j
			}
			mu.Unlock()
		}()

	}

	for a := 0; a < 100; a++ {
		fmt.Println(a, "---", <-c)
	}
}
