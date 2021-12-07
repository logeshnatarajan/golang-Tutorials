package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var in int64
	gr := 100
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&in, 1)
			fmt.Println(atomic.LoadInt64(&in))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(in)

	fmt.Println("all done")

}
