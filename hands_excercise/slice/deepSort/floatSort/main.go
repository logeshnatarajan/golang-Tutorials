package main

import (
	"fmt"
	"sort"
)

func main() {
	var xf = []float64{11.000, 234.6676, 245.688, 324675.67754, 345.76436, 32.7}
	fmt.Println("before sorted", xf)
	bo := sort.Float64sAreSorted(xf)
	fmt.Println(bo)
	sort.Float64s(xf)
	fmt.Println("after sorted", xf)
	bo = sort.Float64sAreSorted(xf)
	fmt.Println(bo)
}
