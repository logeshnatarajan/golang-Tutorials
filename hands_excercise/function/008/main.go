package main

import "fmt"

func main() {

	ff := first(4)
	result := ff(4)
	fmt.Println(result)
}

func first(y int) func(int) int {
	return func(x int) int {
		return sum(x, y)
	}
}
func sum(x, y int) int {
	return x + y
}
