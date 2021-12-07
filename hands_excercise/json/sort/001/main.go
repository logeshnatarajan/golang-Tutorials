package main

import (
	"fmt"
	"sort"
)

func main() {

	xs := []string{"gjhyrfhdf", "ewfv", "gyefwhusw", "wqiu", "afhgh", "zfkh", "bdsfkhb"}
	xi := []int{34, 6, 74544, 67, 437, 9, 643, 56, 7, 3, 1, 32, 4, 66432, 1325, 6}
	fmt.Println("before sort : ", xs)
	sort.Strings(xs)
	fmt.Println("after sort : ", xs)
	fmt.Println("-------")
	fmt.Println("before sort : ", xi)
	sort.Ints(xi)
	fmt.Println("after sort : ", xi)

}
