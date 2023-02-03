package exercise1

import (
	"fmt"
)

func Numbers() {
	a := 10
	b := 2

	if a < b {
		fmt.Printf("%.2f", float64(a)/float64(b))
	} else if b < a {
		fmt.Println(a << b)
	} else {
		if a%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	}

}
