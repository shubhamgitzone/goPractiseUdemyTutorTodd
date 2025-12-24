package main

import "fmt"

var message string = "Hello, Go!"

const pi float64 = 3.14

func main() {

	fmt.Println(message)
	fmt.Printf("Value of pi constant: %v and type of pi: %T\n", pi, pi)

	count := 42
	fmt.Printf("Count variable inside function: %v and type of count : %T\n", count, count)
}
