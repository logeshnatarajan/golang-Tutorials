package main

import "fmt"

func main() {
	for i := 65; i <= 67; i++ {

		fmt.Println("\n##", i)
		for j := 0; j < 3; j++ {

			fmt.Println("\n#", j)
			fmt.Printf("%#U\t", i)

		}

	}

}
