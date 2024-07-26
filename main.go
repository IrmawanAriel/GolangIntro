package main

import (
	"fmt"
)

func main() {
	segitiga(50)
}

func segitiga(x int) {

	for j := x; j >= 0; j-- {

		for k := 0; k <= x-j; k++ {
			fmt.Print(" ")
		}
		for i := (2 * j) - 1; i >= 1; i-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
