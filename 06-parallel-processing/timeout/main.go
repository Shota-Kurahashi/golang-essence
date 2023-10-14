package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generator("Hello!")

	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
		}
	}

	time := time.After(5 * time.Second)

	ch2 := generator("Bye!")

	for {
		select {
		case s := <-ch2:
			fmt.Println(s)
		case <-time:
			fmt.Println("Timeout!")
			return
		}
	}

}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}
