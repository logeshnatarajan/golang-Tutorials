package main

import "fmt"

func main() {

	m := map[string][]string{

		"james bond":       {"shaken not stirred", "martinis", "women"},
		"miss moneypenney": {"james bond", "computer science", "literature"},
		"dr yes":           {"being evil", "ice cream", "sunset"},
	}
	m["logesh"] = []string{"money", "money", "money"}
	fmt.Println(m)
	for k, v := range m {

		fmt.Println(k)
		fmt.Println("\t fav thinges : ")
		for i, val := range v {

			fmt.Println("\t\t\t", i, val)

		}

	}
	delete(m, "dr yes")
	fmt.Println(m)
}
