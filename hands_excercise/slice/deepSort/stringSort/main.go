package main

import (
	"fmt"
	"sort"
)

func main() {
	var xs = []string{"logesh", "arjun", "harshini", "gundan", "ayroo"}
	fmt.Println("before sorted ", xs)
	bo := sort.StringsAreSorted(xs)
	fmt.Println(bo)
	sort.Strings(xs)
	fmt.Println("after sorting ", xs)
	bo = sort.StringsAreSorted(xs)
	fmt.Println(bo)
}
