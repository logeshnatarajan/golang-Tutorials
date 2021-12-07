package main

import (
	"fmt"
	"sort"
)

func main() {
	var xi = []int{5, 6, 7, 8, 2, 4, 6, 3, 7, 9, 5, 3}
	fmt.Println("before sorting ", xi)
	bo := sort.IntsAreSorted(xi)
	fmt.Println(bo)
	sort.Ints(xi)
	fmt.Println("after sorting", xi)
	bo = sort.IntsAreSorted(xi)
	fmt.Println(bo)
}
