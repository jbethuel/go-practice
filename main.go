package main

import (
	"fmt"
	arrayMapAndStruct "go-practice/array_map_and_struct"
	Input "go-practice/input"
	Loop "go-practice/loop"
	Strings "go-practice/string"
)

func main() {
	fmt.Println("=== strings ===")
	Strings.SimpleStrings()
	fmt.Println("=== inputs ===")
	Input.SimpleInput()
	fmt.Println("=== loops ===")
	Loop.SimpleLoops()
	fmt.Println("=== array, map and struct ===")
	arrayMapAndStruct.SimpleArray()
	arrayMapAndStruct.Maps()
	arrayMapAndStruct.Structs()
}
