package main

import "fmt"

type person struct {
	name string
	ss   string
}

func (a *person) speak() {

	fmt.Println(a.name)
	fmt.Println(a.ss)

}

type human interface {
	speak()
}

func ssm(h human) {
	h.speak()
}

func main() {

	p1 := person{"podu", "aatam"}
	ssm(&p1)

}
