package main

import "fmt"

func main() {
	a := foo()
	b, c := bar()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
func foo() int {

	return 43

}
func bar() (int, string) {
	return 44, "hello"

}
