package main

import (
	"fmt"
	"time"
)

// This is an example of race condition
// 2 goroutines tries to read&write sharedInt and there is no access control.

var sharedInt int = 0
var unusedValue int = 0

func runSimpleReader() {
	for {
		var val int = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
			fmt.Println("used", unusedValue)
		}
	}
}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
		fmt.Println(sharedInt)
	}
}

func main() {
	fmt.Println("start")
	go runSimpleReader()
	go runSimpleWriter()
	fmt.Println(sharedInt)
	time.Sleep(2 * time.Second)
}
