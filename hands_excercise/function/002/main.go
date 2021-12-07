package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	a := foo(xi...)

	fmt.Println("from foo ", a)

	b := bar(xi)
	fmt.Println("from bar:", b)

}
func foo(v ...int) int {
	total := 0
	for _, val := range v {

		total += val

	}
	return total

}
func bar(i []int) int {

	total := 0
	for _, val := range i {

		total += val

	}
	return total

}
