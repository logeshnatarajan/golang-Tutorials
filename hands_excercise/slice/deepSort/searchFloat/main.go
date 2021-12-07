package main

import (
	"fmt"
	"sort"
)

func main() {
	var xi = []float64{53.876, 643.4356, 634.45, 64.87, 6.23456, 3.2345, 753.435, 655.890, 28.5678}
	index := sort.SearchFloat64s(xi, 64.87)
	fmt.Println(xi)
	fmt.Println("64.87 in position of : ", index)
	sort.Float64s(xi)
	//oru velai antha no slice la illana athu enga fit agumo antha place aa sollu
	index = sort.SearchFloat64s(xi, 64.87)
	fmt.Println("64.87 is in the position of : ", index)
	fmt.Println(xi)
}
