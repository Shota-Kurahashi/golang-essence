package forloop

import "fmt"

func Main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	i := 0

	for i < 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}

	for {
		fmt.Printf("i: %v\n", i)
		i++

		if i == 15 {
			break
		}
	}

	obj := []string{"foo", "bar", "baz"}

	for k, v := range obj {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}

finished:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break finished
			}
		}
	}
}
