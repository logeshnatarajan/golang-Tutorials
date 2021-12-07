package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type customerr struct {
	err error
}

func (a customerr) Error() string {

	return fmt.Sprint(fmt.Errorf("its a error "))
}

type foo struct {
	First string
	Last  string
}

func main() {
	p1 := customerr{

		err: fmt.Errorf("wrong"),
	}

	f1 := foo{"logesh", "eugreu"}
	fmt.Println(f1)
	bs, errr := json.Marshal(f1)
	if errr != nil {
		errr := fmt.Errorf("the error is %v", p1.err)
		log.Fatal(errr)

	}
	fmt.Println(string(bs))

}
