package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Printf("Type a number to check if it is odd or even:\n")
	fmt.Scanln(&input)
	input = math.Mod(input, 2)

	var result string
	if input == 0 {
		result = "even"
	} else {
		result = "odd"
	}

	fmt.Printf("User input is %v\n", result)
}
