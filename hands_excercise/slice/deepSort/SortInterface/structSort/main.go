package main

import (
	"fmt"
	"sort"
)

type Mobile struct {
	Brand string
	Price int
}
type ByPrice []Mobile

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

// ByBrand implements sort.Interface for []Mobile based on
// the Brand field.
type ByBrand []Mobile

func (a ByBrand) Len() int           { return len(a) }
func (a ByBrand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBrand) Less(i, j int) bool { return a[i].Brand > a[j].Brand }

func main() {
	var ss = []Mobile{{Brand: "logesh", Price: 800}, {"samsung", 900}, {"oneplus", 600}, {"aaa", 01}}
	fmt.Println("before sorting ", ss)
	sort.Sort(ByPrice(ss))
	fmt.Println("based on price", ss)
	sort.Sort(ByBrand(ss))
	fmt.Println("based on brand", ss)
}
