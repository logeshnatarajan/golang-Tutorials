package main

import "fmt"

func main() {

	m := map[string][]string{

		"james bond":       {"shaken not stirred", "martinis", "women"},
		"miss moneypenney": {"james bond", "computer science", "literature"},
		"dr yes":           {"being evil", "ice cream", "sunset"},
	}
	fmt.Println(m)
	for k, v := range m {

		fmt.Println(k)
		fmt.Println("\t fav thinges : ")
		for i, val := range v {

			fmt.Println("\t\t\t", i, val)

		}

	}

}
