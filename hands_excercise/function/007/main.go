package main

import "fmt"

func main() {
	fu := bar()
	a := foo(fu)
	fmt.Println(a)

}

func foo(f func() int) int {
	i := f()
	i++
	return i
}
func bar() func() int {
	return func() int {
		return 42
	}
}
