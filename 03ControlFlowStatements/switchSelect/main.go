package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Println("x =", x)

	switch {
	case 0 <= x && x < 100:
		fmt.Println("x is less than 100")
	case 100 <= x && x < 200:
		fmt.Println("x is between 100 and 200")
	default:
		fmt.Println("x is greater than or equal to 200")
	}

	y := rand.Intn(10)
	z := rand.Intn(10)
	fmt.Println("y =", y, "z =", z)

	if y+z < 4 {
		fmt.Println("y + z is less than 4")
	} else if y+z >= 6 {
		fmt.Println("y + z is greater than 6")
	} else if y >= 4 && z <= 6 {
		fmt.Println("y is greater than or equal to 4 and z is less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not equal to 5")
	} else {
		fmt.Println("None of the above conditions met")
	}

	switch y + z {
	case 0, 1, 2, 3:
		fmt.Println("y + z is less than 4")
	case 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		fmt.Println("y + z is greater than 6")
	case 4, 5:
		if y >= 4 && z <= 6 {
			fmt.Println("y is greater than or equal to 4 and z is less than or equal to 6")
		} else if y != 5 {
			fmt.Println("y is not equal to 5")
		}
	default:
		fmt.Println("None of the above conditions met")
	}
}

func init() {
	// This init function runs before main
	// It can be used for setup tasks
	fmt.Println("Switch Statement Example")
}
