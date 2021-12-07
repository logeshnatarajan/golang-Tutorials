package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("no_file")
	if err != nil {
		//fmt.Println("OMG... err haha", err)
		log.Println("OMG... err haha", err)
		//log.Fatalln(err)
		//panic(err)
	}

}
