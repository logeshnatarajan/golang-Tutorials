package main

import "fmt"

type customerr struct {
	info  string
	info2 string
}

func (dd customerr) Error() string {

	return fmt.Sprintf("here is the error %v", dd.info)

}
func main() {

	c1 := customerr{
		info:  "poduuuuu",
		info2: "poduuu22222",
	}
	c2 := customerr{
		info:  "hhhhhh",
		info2: "hhhhhh2222",
	}
	foo(c1)
	fmt.Println(c1)
	fmt.Println(c1.info2)
	fmt.Println("ithu c2", c2)

}
func foo(e error) {

	fmt.Println("foo ran", e)
}
