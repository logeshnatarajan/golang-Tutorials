package main

import "fmt"

func main() {

	a := 55
	fmt.Println(&a)
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 56
	fmt.Println(a)

}
