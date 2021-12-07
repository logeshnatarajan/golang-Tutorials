package main

import "fmt"

type person struct {
	age int
}

func changeme(a *person) {
	*a = person{21}
	//a.age = 24
}
func main() {

	p1 := person{

		age: 22,
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)
}
