package main

import (
	"fmt"
	"sync"
	"time"
)

func download(url string) {
	println("Downloading", url)
	time.Sleep(6 * time.Second)
}

func main() {
	before := time.Now()
	limit := make(chan struct{}, 20)

	var wg sync.WaitGroup

	// 100個のgoroutineを起動する。すると、100個のgoroutineが同時に実行されてしまい、取得先のサーバーに負荷がかかる。
	// そのため、goroutineの数を制限する必要がある。
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			limit <- struct{}{}

			defer wg.Done()
			u := fmt.Sprintf("https://example.com/%d", i)
			download(u)

			<-limit
		}(i)

	}

	wg.Wait()
	fmt.Printf("Total time: %s", time.Since(before))
}
