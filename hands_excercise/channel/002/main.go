package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")

}
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 15; i++ {
			c <- i

		}
		// closing is must while ranging the channel
		close(c)
	}()

	return c

}
func receive(c <-chan int) {
	// ranging channel keep listening the channel untill its closed
	for v := range c {
		fmt.Println(v)
	}
}
