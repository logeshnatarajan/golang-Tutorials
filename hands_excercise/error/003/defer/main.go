package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	fmt.Println("poduuuuuuuu")
	_, err := os.Open("nofile.txt")
	if err != nil {
		//log.Fatalln(err)
		log.Panic(err)
	}
	defer bar()
}
func foo() {

	fmt.Println("from foooooooo")

}
func bar() {
	fmt.Println("jghdfhbsehfaeiushfbgy")
}
