package main

import (
	"builtin-function/formats"
	"builtin-function/jsons"
	"builtin-function/logs"
	"builtin-function/paths"
	"builtin-function/times"
	"fmt"
)

func init() {
	fmt.Println("init")
	// 外部ライブラリで初期化処理を行う場合に使用されていることがある。
	// 例えば、データベースの接続処理など
}

func main() {
	// print系はデバッグ用のメッセージ出力関数という位置づけ
	print("Hello, World!")
	println("Hello, World!")

	// 標準出力に出力したい場合はfmtパッケージを使う(本番環境ではこちらを使う)
	fmt.Print("Hello, World!")

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// var a [2]int

	// n := 2

	// // 配列の範囲外アクセスはランタイムパニックを引き起こす
	// println(a[n])

	formats.Main()
	logs.Logs()
	jsons.Jsons()
	// nets.Main()
	times.Main()
	paths.Main()

}
