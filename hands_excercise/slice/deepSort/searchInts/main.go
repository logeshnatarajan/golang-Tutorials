package main

import (
	"fmt"
	"sort"
)

// The SearchInts function searches the position of x in a sorted slice of int and
// returns the index as specified by Search.
// This function works if slice is in sort order only.
// If found x in intSlice then it returns index position of intSlice otherwise
// it returns index position where x fits in sorted slice.
// The following example shows the usage of SearchInts() function:

func main() {
	var xi = []int{53, 643, 634, 64, 6, 3, 753, 655, 28}
	index := sort.SearchInts(xi, 64)
	fmt.Println(xi)
	fmt.Println("64 in position of : ", index)
	sort.Ints(xi)
	//oru velai antha no slice la illana athu enga fit agumo antha place aa sollu
	index = sort.SearchInts(xi, 64)
	fmt.Println("64 is in the position of : ", index)
	fmt.Println(xi)
}
