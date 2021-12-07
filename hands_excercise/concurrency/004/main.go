package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex // &sync.Mutex{}
	in := 0
	gr := 100
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			mx.Lock()
			v := in
			//runtime.Gosched()
			v++
			in = v
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(in)

	fmt.Println("all done")

}
