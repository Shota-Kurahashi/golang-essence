package anys

import (
	"fmt"
	"reflect"
)

func Main() {
	var v interface{}

	v = 1
	// 型アサーション
	n := v.(int)

	v = "foo"
	fmt.Printf("n: %v\n", n)
	s := v.(string)

	fmt.Printf("s: %v\n", s)

	s, ok := v.(string)

	if !ok {
		panic("v is not string")
	}

	fmt.Printf("s: %v\n", s)

	if c, ok := v.(string); !ok {
		fmt.Printf("v is not string\n")
	} else {
		fmt.Printf("v is string: %v\n", c)
	}

	PrintDetails(1)

}

func PrintDetails(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("args is int: %v\n", v)
	case string:
		fmt.Printf("args is string: %v\n", v)
	default:
		fmt.Printf("args is unknown type\n")
	}
}

func PrintSelfDetails(v interface{}) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int:
		fmt.Printf("args is int: %v\n", v)
	case reflect.String:
		fmt.Printf("args is string: %v\n", v)
	default:
		fmt.Printf("args is unknown type\n")
	}

}
