package main

import (
	"fmt"
	"sort"
)

func main() {
	var xi = []int{9, 8, 7, 6, 5, 4, 3, 2}
	i := sort.Search(len(xi), func(i int) bool {
		return xi[i] <= 5
	})
	fmt.Print(i)
	i = sort.SearchInts(xi, 5)
	fmt.Print("\t", i)
	sort.Ints(xi)
	i = sort.SearchInts(xi, 5)
	fmt.Print("\t", i)
	//same goes for string float64

}
