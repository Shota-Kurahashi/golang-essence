package pointer

import "fmt"

type User struct {
	Name string
	age  int
}

func (u *User) Add(v int) {
	u.age += v
}

func Main() {
	v := 1
	p := &v
	*p = 2
	fmt.Printf("v: %v\n", v)

	user := new(User)
	fmt.Printf("user: %v\n", user)
	user.Name = "foo"
	user.age = 10
	fmt.Printf("user: %v\n", user)

	user.Add(10)
	fmt.Printf("user: %v\n", user.age)

}
