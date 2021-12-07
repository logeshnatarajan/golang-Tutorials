package main

import "fmt"

func main() {

	a := func() int {

		return 42

	}()
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := foo()
	fmt.Println(b())
	fmt.Printf("%T", b)

}
func foo() func() int {
	return func() int {

		return 48
	}
}
