package main

import (
	"errors"
	"fmt"
	"log"
)

var infoerr = errors.New("no. must be greater then 6")

func main() {
	fmt.Printf("%T\n", infoerr)
	var a float64
	fmt.Scanf("%f", a)

	_, err := sum(a)
	if err != nil {
		log.Fatalln(err)

	}

}
func sum(n float64) (float64, error) {
	if n < 6 {
		return 66.6, infoerr
	}
	return n + 6, nil
}
