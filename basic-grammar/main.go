package main

import (
	"basic-grammar/foo"
	"fmt"
	"reflect"
	"sort"
)

func main() {

	typeF()
	forLoop()
	arraySlice()
	stringF()
	mapF()
	structF()
	pointF()
	anyF()
}

func typeF() {
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

func forLoop() {
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

func arraySlice() {
	// 配列は固定長
	var a [4]int
	a[0] = 1
	fmt.Printf("a: %v\n", a)

	var b []int

	fmt.Printf("b: %v\n", b)

	// スライスは可変長
	c := make([]int, 3)
	// c := []int{1, 2, 3} でもOK

	fmt.Printf("c: %v\n", c)

	c[0] = 1
	c[1] = 2
	c[2] = 3

	// c[3] = 4 はエラー

	fmt.Printf("c: %v\n", c)

	// 値の追加
	c = append(c, 4, 5, 6)

	fmt.Printf("c: %v\n", c)
	fmt.Printf("cの長さは: %v\n", len(c))

	// 長さ0 で容量100のスライス
	d := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		// 容量を超えると倍の容量で再確保される。パファーマンスに影響するので、あらかじめ容量を確保しておく
		d = append(d, i)
	}

	fmt.Printf("d: %v\n", d)
	// スライスの一部を取得 (2番目から4番目まで)
	fmt.Printf("d: %v\n", d[2:5])

	// * スライスから要素を消去する方法
	// 1. 新しく同じ型のスライスを作成し、詰め直す
	// 2. appendを使う
	// 3. 部分参照とcopyを使用する

	base := []int{1, 2, 3, 4, 5}

	//? 1. 新しく同じ型のスライスを作成し、詰め直す
	deleted := make([]int, 0, len(base)/2)

	for i := 0; i < len(base); i++ {
		if i%2 == 0 {
			deleted = append(deleted, base[i])
		}
	}

	base = deleted

	fmt.Printf("base: %v\n", base)

	//? 2. appendを使う

	base2 := []int{1, 2, 3, 4, 5}
	n := 2

	base2 = append(base2[:n], base2[n+1:]...)

	fmt.Printf("base: %v\n", base2)

	//? 3. 部分参照とcopyを使用する

	n2 := 2

	base3 := []int{1, 2, 3, 4, 5}

	base3 = base3[:n2+copy(base3[n2:], base3[n2+1:])]

	fmt.Printf("base: %v\n", base3)
}

func stringF() {
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

func mapF() {
	// tsのobjectと同じ
	var m map[string]int

	fmt.Printf("m: %v\n", m)

	m2 := make(map[string]int)

	m2["John"] = 1
	m2["Paul"] = 2
	m2["George"] = 3

	fmt.Printf("m2: %v\n", m2)

	// 順番は保証されない
	for k, v := range m2 {
		fmt.Printf("k: %v, v: %v \n", k, v)
	}

	keys := []string{}

	for k := range m2 {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("k: %v, v: %v \n", k, m2[k])
	}

	// 消去
	delete(m2, "John")

	m3 := m2

	delete(m2, "Paul")

	// m3も消える。つまり、参照渡し
	fmt.Printf("m2: %v\n", m2)
	fmt.Printf("m3: %v\n", m3)

	// キーの存在確認
	_, ok := m2["Paul"]

	fmt.Printf("ok: %v\n", ok)
}

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

func structF() {

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

func (u *User) Add(v int) {
	u.age += v
}

func pointF() {
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

func anyF() {
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
