package main

import "fmt"

func main() {
	a := []string{"james", "bond", "shaken and not stirred"}
	b := []string{"miss", "moneypenny", "hellooo james"}
	x := [][]string{a, b}
	fmt.Println(x)
	for i, v := range x {

		fmt.Println("person #", i)
		for j, val := range v {

			fmt.Println("\t\t", j, "--", val)

		}

	}

}
