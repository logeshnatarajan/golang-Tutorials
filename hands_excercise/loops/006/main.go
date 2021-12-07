package main

import "fmt"

func main() {

	switch {

	case false:
		fmt.Println("varathu")
	case true:
		fmt.Println("ithu varu")
		fallthrough
	case false:
		fmt.Println("ithu false analu varu")

	}

}
