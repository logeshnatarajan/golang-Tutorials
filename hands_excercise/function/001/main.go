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

	return 42

}
func bar() (int, string) {
	return 43, "hello"

}
