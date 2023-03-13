package cli

import (
	"fmt"
	"os"
	"strconv"
)

func GetMaxValue() {
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err == nil {
			iArgs = append(iArgs, val)
		}
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}

	fmt.Println("Max value", max)
}
