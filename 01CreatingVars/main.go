package main

import "fmt"

var message string = "Hello, Go!"

const pi float64 = 3.14

func main() {

	fmt.Println(message)
	fmt.Println("Value of pi constant:", pi)

	count := 42
	fmt.Println("Count variable inside function:", count)
}
