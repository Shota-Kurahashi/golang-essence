package defers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b [512]byte
	n, err := f.Read(b[:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bytes read: ", n)

	doSomething("data")
}

func doSomething(dir string) error {
	err := os.Mkdir(dir, 0755)

	if err != nil {
		return err
	}

	// 関数の呼び出しを遅延させる
	// 引数はすぐに評価されるが、関数の呼び出しは遅延される
	defer os.RemoveAll(dir)

	f, err := os.Create(filepath.Join(dir, "temp.txt"))

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString("data")

	if err != nil {
		return err
	}

	return nil
}
