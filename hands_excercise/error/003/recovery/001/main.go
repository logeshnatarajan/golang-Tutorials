package main

import "fmt"

func main() {

	f()
	fmt.Println("return normally form  f")

}
func f() {
	defer func() {

		if r := recover(); r != nil {

			fmt.Println("recovering r", r)

		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("returned from g")

}
func g(n int) {
	if n > 3 {
		fmt.Println("OMG panickingggggggggg")
		panic(fmt.Sprintf("%v", n))

	}
	defer fmt.Println("defer g ", n)
	fmt.Println("printing g ", n)
	g(n + 1)

}
