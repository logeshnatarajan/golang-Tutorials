package main

import "fmt"

func main() {

	a := []int{2, 3, 4, 5, 7}
	fmt.Println(a)
	foo(a)
	fmt.Println(a)
}
func foo(b []int) {
	b[0], b[1] = b[1], b[0]
}
