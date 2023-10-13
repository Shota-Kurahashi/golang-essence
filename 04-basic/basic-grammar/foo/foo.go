package foo

import "fmt"

var index = 1

func UpdateUser() {
	fmt.Printf("foo UpdateUser() %d\n", index)
}

func FindUser(name string) (string, error) {
	fmt.Printf("foo FindUser() %d\n", index)
	return "foo", nil
}
