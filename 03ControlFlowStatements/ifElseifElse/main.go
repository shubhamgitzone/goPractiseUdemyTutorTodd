package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)
	fmt.Println("x =", x)

	if 0 >= x && x < 100 {
		fmt.Println("x is less than 100")
	} else if 100 <= x && x < 200 {
		fmt.Println("x is between 100 and 200")
	} else {
		fmt.Println("x is greater than or equal to 200")
	}
}
