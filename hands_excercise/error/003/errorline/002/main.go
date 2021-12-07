package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("kk.txt")
	if err != nil {
		fmt.Println(err)

	}
	log.SetOutput(f)
	f2, err := os.Open("nofile.txt")

	if err != nil {

		log.Println("err happened haha", err)
	}
	defer f2.Close()
	fmt.Println("to look error see in kk.text")

}
