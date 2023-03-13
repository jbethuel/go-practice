package componsition

import "fmt"

type contact struct {
	firstName string
	lastName  string
	address
}

type address struct {
	city string
	zip  string
}

func info(c contact) {
	fmt.Printf("%s lives in %s\n", c.firstName, c.address.city)
}

func Compose() {
	address := address{
		zip:  "8700",
		city: "Malaybalay City",
	}
	person := contact{
		firstName: "Bethuel",
		lastName:  "Dela Cruz",
		address:   address,
	}

	info(person)
}
