package char

import "fmt"

func Main() {
	s := "Hello"
	s += "World"

	fmt.Printf("s: %v\n", s)

	// 添字でアクセスできる
	fmt.Printf("s[0]: %c\n", s[0])

	// 書き換える場合は、一度[]byteに変換する
	b := []byte(s)

	b[0] = 'h'

	s = string(b)

	fmt.Printf("s: %v\n", s)

	var content = `Hello
	World
	!
	`

	fmt.Printf("content: %v\n", content)
}
