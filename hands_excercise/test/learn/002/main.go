package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println(sum)

}
func foo(n ...int) int {

	total := 0
	for _, v := range n {
		total += v

	}
	return total

}
