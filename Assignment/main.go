package main

import (
	"fmt"
)

func main () {
	ints := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, i := range ints {
		if i % 2 != 0 {
			fmt.Printf("%d Odd\n", i)
		} else {
			fmt.Printf("%d Even\n", i)
		}
	}
}