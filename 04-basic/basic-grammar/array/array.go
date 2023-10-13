package array

import "fmt"

func Main() {
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
