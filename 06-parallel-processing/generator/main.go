package main

import (
	"fmt"
	"time"
)

func generator(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}

func main() {
	ch := generator("Hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	ch2 := fanIn(fanInGenerator("fanIn Hello"), fanInGenerator("fanIn Bye"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch2)
	}

}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()

	go func() {
		for {
			new_ch <- <-ch2
		}
	}()

	return new_ch
}

func fanInGenerator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}
