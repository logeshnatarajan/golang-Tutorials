package main

import "fmt"

func main() {
	fmt.Println(sequence(5)(6)(7))
}
func sequence(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return (x * x) + (y * y) + (z * z)
		}
	}
}
