package main

import "fmt"

type (
	first  = func(int) int
	second = func(int) first
)

func main() {
	fmt.Println(sequence(6)(7)(8))
}
func sequence(x int) second {
	return func(y int) first {
		return func(z int) int {
			return (x * x) + (y * y) + (z * z)
		}
	}
}
