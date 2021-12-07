package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {
	u1 := user{"logesh", 21}
	u2 := user{"kan", 22}
	u3 := user{"podu", 23}
	u4 := user{"mmm", 234}

	users := []user{u1, u2, u3, u4}
	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

}
