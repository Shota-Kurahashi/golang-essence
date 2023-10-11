package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(message string) {
	println(message)
}

func main() {
	message := "h1"
	go func() {
		sendMessage(message)
	}()

	message = "h2"

	waiting()
	mutex()
	channel()

}

func waiting() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// ここが実行される前にmain関数が終了してしまう
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(time.Second)
	}()

	fmt.Println("Waiting...")

	wg.Wait()

	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			fmt.Printf("%dth goroutine sleeping...\n", i+1)
			time.Sleep(time.Second)

		}(i)

	}

	wg2.Wait()
}

func mutex() {
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()

		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()

		}
	}()

	wg.Wait()

	fmt.Println(n)
}

func server(ch chan string) {
	defer close(ch)

	ch <- "one"
	ch <- "two"
	ch <- "three"

}

func channel() {
	var s string

	ch := make(chan string)
	go server(ch)
	s = <-ch
	fmt.Println(s)
	fmt.Println("----")
	s = <-ch
	fmt.Println(s)
	fmt.Println("----")
	s = <-ch
	fmt.Println(s)
}
