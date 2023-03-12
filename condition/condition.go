package condition

import "fmt"

func IfAndElse() {

	const number = 5

	if number < 5 {
		fmt.Println("Number is lesser than 5")
	} else {
		fmt.Println("Number is greater than 5")
	}
}

func Switch() {
	const city string = "malaybalay"

	switch city {

	case "malaybalay":
		{

			fmt.Println(("It is malaybalay"))
		}

	case "cdo":
		{

			fmt.Println(("It is cdo"))
		}

	default:
		{
			fmt.Println("None")
		}

	}
}
