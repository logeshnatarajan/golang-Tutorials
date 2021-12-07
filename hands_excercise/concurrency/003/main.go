package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	in := 0
	gr := 100
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := in
			runtime.Gosched()
			v++
			in = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(in)

	fmt.Println("all done")

}
