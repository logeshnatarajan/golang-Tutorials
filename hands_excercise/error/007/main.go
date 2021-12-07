package main

import (
	"fmt"
	"log"
)

type sqer struct {
	lat  string
	long string
	err  error
}

func (se sqer) Error() string {

	return fmt.Sprintf("THE TYPE INTERFACE ERROR IS %v%v%v", se.lat, se.long, se.err)

}
func main() {

	_, err := sqrt(-10)
	if err != nil {

		log.Fatalln(err)
	}

}
func sqrt(f float64) (float64, error) {

	if f < 0 {

		//c1 := errors.New("new la tha pannirukan")
		c1 := fmt.Errorf("la use panniruka value vanthu %v", f)
		return 45, sqer{"lattttt", "longgggggg", c1}

	}
	return 44, nil
}
