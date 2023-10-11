package maps

import (
	"fmt"
	"sort"
)

func Main() {
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
