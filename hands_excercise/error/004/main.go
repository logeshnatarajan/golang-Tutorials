package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

var infoerr = errors.New("IT HAS ERROR WHILE CONVERTING TO JSON")

//var infoerr = fmt.Errorf("IT HAS ERROR WHILE CONVERTING TO JSON")
type person struct {
	First  string
	Last   string
	Saying []string
}

func main() {

	p1 := person{

		First: "logesh",
		Last:  "mmm",

		Saying: []string{"poduuuu", "aatam", "poduuu!!!"},
	}
	bs, err := json.Marshal(p1)
	if err != nil {

		log.Println(infoerr)
		return
		//log.Fatalln(infoerr)  -- ithuku return thevai illa
	}
	fmt.Println(string(bs))

}
