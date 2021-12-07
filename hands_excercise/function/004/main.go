package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (r person) speak() {

	fmt.Println("name is :", r.first, r.last)
	fmt.Println("age : ", r.age)

}

func main() {

	p1 := person{"logesh", "kan", 21}
	fmt.Println(p1)
	p1.speak()

}
