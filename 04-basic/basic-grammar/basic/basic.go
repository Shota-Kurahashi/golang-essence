package basic

import (
	"basic-grammar/foo"
	"fmt"
)

func Main() {
	fmt.Println("こんにちは、世界！")

	var any interface{} = "これは任意の型を持てる"
	fmt.Printf("any (任意の型): %v\n", any)

	var _bool bool = true
	fmt.Printf("_bool (真偽値): %v\n", _bool)

	var _byte byte = 65 // ASCII for 'A'
	fmt.Printf("_byte (バイト型、0-255): %v\n", _byte)

	var _complex64 complex64 = 1 + 2i
	fmt.Printf("_complex64 (複素数、単精度): %v\n", _complex64)

	var _complex128 complex128 = 3 + 4i
	fmt.Printf("_complex128 (複素数、倍精度): %v\n", _complex128)

	var _error error = fmt.Errorf("これはエラーです")
	fmt.Printf("_error (エラー型): %v\n", _error)

	var _float32 float32 = 3.14
	fmt.Printf("_float32 (浮動小数点数、単精度): %v\n", _float32)

	var _float64 float64 = 2.71828
	fmt.Printf("_float64 (浮動小数点数、倍精度、基本的にはこちら): %v\n", _float64)

	var _int int = 42
	fmt.Printf("_int (符号あり整数): %v\n", _int)

	var _int8 int8 = -128
	fmt.Printf("_int8 (8ビット符号あり整数): %v\n", _int8)

	var _int16 int16 = 32767
	fmt.Printf("_int16 (16ビット符号あり整数): %v\n", _int16)

	var _int32 int32 = 2147483647
	fmt.Printf("_int32 (32ビット符号あり整数): %v\n", _int32)

	var _int64 int64 = 9223372036854775807
	fmt.Printf("_int64 (64ビット符号あり整数): %v\n", _int64)

	var _rune rune = '日'
	fmt.Printf("_rune (文字、UTF-8): %v\n", _rune)

	var _string string = "こんにちは"
	fmt.Printf("_string (文字列): %v\n", _string)

	var _uint uint = 42
	fmt.Printf("_uint (符号なし整数): %v\n", _uint)

	var _uint8 uint8 = 255
	fmt.Printf("_uint8 (8ビット符号なし整数): %v\n", _uint8)

	var _uint16 uint16 = 65535
	fmt.Printf("_uint16 (16ビット符号なし整数): %v\n", _uint16)

	var _uint32 uint32 = 4294967295
	fmt.Printf("_uint32 (32ビット符号なし整数): %v\n", _uint32)

	var _uint64 uint64 = 18446744073709551615
	fmt.Printf("_uint64 (64ビット符号なし整数): %v\n", _uint64)

	var _uintptr uintptr = 0xFFFFFFF
	fmt.Printf("_uintptr (ポインタの数値表現): %v\n", _uintptr)

	const _const = "これは定数です"

	fmt.Printf("_const (定数): %v\n", _const)

	const (
		Apple = iota
		Orange
		Banana
	)

	fmt.Printf("iota (連番): %v, %v, %v\n", Apple, Orange, Banana)

	const (
		Apple2 = iota + iota
		Orange2
		Banana2 = iota + 3
		Mikan
	)

	fmt.Printf("iota (連番): %v, %v, %v, %v\n", Apple2, Orange2, Banana2, Mikan)

	type Animal int

	const (
		Dog Animal = iota
		Cat
		Monkey
	)

	fmt.Printf("iota type (連番): %v, %v, %v\n", Dog, Cat, Monkey)

	foo.UpdateUser()
	name, err := foo.FindUser("foo")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("name: %v\n", name)
	}
}
