package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	js := `[{"Name":"logesh","Age":21},{"Name":"kan","Age":22},{"Name":"podu","Age":23},{"Name":"mmm","Age":234}]`
	var users []user

	err := json.Unmarshal([]byte(js), &users)
	if err != nil {
		fmt.Println(users)

	}
	fmt.Println(users)
}
