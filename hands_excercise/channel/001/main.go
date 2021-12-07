package main

import (
	"fmt"
)

var mm = map[string]int{"string": 1}

func main() {

	ch := make(chan int)
	go func() {

		ch <- 42
		ch <- 44

	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	mapp := make(chan *map[string]int)
	go func() {
		mapp <- &mm
	}()
	tem := <-mapp
	fmt.Println(tem)

}
