package main

import (
	"fmt"
	"sort"
)

func main() {
	var xi = []string{"abc", "dkjfb", "qdrhg", "oiewru", "psorh", "sdryg", "qdcdv", "asdf", "mfnb"}
	index := sort.SearchStrings(xi, "psorh")
	fmt.Println(xi)
	fmt.Println("psorh in position of : ", index)
	sort.Strings(xi)
	//oru velai antha no slice la illana athu enga fit agumo antha place aa sollu
	index = sort.SearchStrings(xi, "psorh")
	fmt.Println("psorh is in the position of : ", index)
	fmt.Println(xi)
}
