package main

import (
	"errors"
	"log"
)

func main() {

	_, err := sum(2)
	if err != nil {
		log.Fatalln(err)

	}

}
func sum(n float64) (float64, error) {
	if n < 6 {
		return 66.6, errors.New("no. must be greater then 6")
	}
	return n + 6, nil
}
