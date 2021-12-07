package main

import "fmt"

func main() {

	defer foo()
	bar()

}

func foo() {

	fmt.Println("poduuu foo")

}
func bar() {

	fmt.Println("poduu first bar tha")

}
