package arrayMapAndStruct

import "fmt"

func SimpleArray() {
	var books []string

	books = append(books, "First")
	books = append(books, "Second")
	books = append(books, "Third")

	fmt.Println(books)
}

func Maps() {
	// accepts only 1 data type
	var user = make(map[string]string, 0)
	user["id"] = "123"
	user["email"] = "jbethueldc@gmail.com"
	user["name"] = "Bethuel Dela Cruz"

	fmt.Println(user)
}

func Structs() {
	// accepts multiple data types
	type User struct {
		id    string
		email string
		name  string
		age   uint
	}

	var users = make([]User, 0)
	var user = User{
		id:    "123",
		email: "jbethueldc@gmail.com",
		name:  "Bethuel Dela Cruz",
		age:   28,
	}
	users = append(users, user)

	fmt.Println((users))
}
