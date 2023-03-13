package pointers

type Person struct {
	firstName string
	lastName  string
	tags      []string
}

func addTag(personArg *Person, tag string) {
	personArg.tags = append(personArg.tags, tag)
}

func Get() Person {
	var person Person
	person.firstName = "Bethuel"
	person.lastName = "Dela Cruz"
	person.tags = make([]string, 0)

	addTag(&person, "Cool")
	addTag(&person, "Calm")

	return person
}
