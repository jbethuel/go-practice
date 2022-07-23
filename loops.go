package main

import (
	"fmt"
)

func main() {
	var iteration int = 10

	for {
		fmt.Printf("iteration #%v\n", iteration)
		iteration = iteration - 1

		if iteration == 0 {
			break
		}
	}
}
