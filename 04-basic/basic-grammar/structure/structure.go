package structure

import (
	"fmt"
)

type User struct {
	Name string
	// 小文字は外部からアクセスできない
	age int
}

func (u User) AddAge(n User) User {
	return User{
		Name: u.Name,
		age:  u.age + n.age,
	}
}

func showName(user *User) {
	fmt.Printf("user.Name: %v\n", user.age)
	user.Name = "bar"
}

func (u *User) Add(v int) {
	u.age += v
}

func Main() {

	var user User

	user.Name = "foo"
	user.age = 10

	// 関数の引数に渡すときは、コピーが作成されるので、参照渡しにならない
	// つまり、関数内で値を変更しても、呼び出し元の値は変更されない
	// 更新したい場合は、ポインタを渡す

	showName(&user)

	fmt.Printf("user.Name: %v\n", user.Name)

	user2 := User{
		Name: "foo",
		age:  10,
	}

	user2 = user2.AddAge(user2)

	fmt.Printf("user2: %v\n", user2.age)
}
