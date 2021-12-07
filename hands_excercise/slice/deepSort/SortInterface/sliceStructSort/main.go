package main

import (
	"fmt"
	"sort"
)

func main() {
	mobile := []struct {
		Brand string
		Price int
	}{
		{"Nokia", 700},
		{"Samsung", 505},
		{"Apple", 924},
		{"Sony", 655},
	}
	result := sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile[i].Price < mobile[j].Price })
	fmt.Println("Found price sorted:", result) // false

	mobile = []struct {
		Brand string
		Price int
	}{
		{"Nokia", 700},
		{"Samsung", 805},
		{"Apple", 924},
		{"Sony", 955},
	}
	result = sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile[i].Price < mobile[j].Price })
	fmt.Println("Found price sorted:", result) // true

	mobile = []struct {
		Brand string
		Price int
	}{
		{"iPhone", 900},
		{"MI", 805},
		{"OPPO", 724},
		{"Sony", 655},
	}
	result = sort.SliceIsSorted(mobile, func(i, j int) bool { return mobile[i].Brand < mobile[j].Brand })
	fmt.Println("Found brand sorted:", result) // false
}
