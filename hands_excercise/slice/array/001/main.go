package main

import "fmt"

func main() {

	var a [5]int
	for i, v := range a {
		fmt.Println(i, "--", v)

	}
	fmt.Printf("%T", a)

}
