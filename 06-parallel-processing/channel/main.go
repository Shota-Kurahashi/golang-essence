package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func download(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {

		resp, err := http.Get(u)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			log.Fatalln(err)
		}

		ch <- body
	}
}

func main() {
	urls := []string{"https://jsonplaceholder.typicode.com/posts", "https://jsonplaceholder.typicode.com/comments", "https://jsonplaceholder.typicode.com/albums", "https://jsonplaceholder.typicode.com/photos", "https://jsonplaceholder.typicode.com/todos", "https://jsonplaceholder.typicode.com/users"}

	ch := make(chan []byte)

	var wg sync.WaitGroup

	wg.Add(1)
	go download(&wg, urls, ch)

	// 呼ばれる順番は保証されない。
	for body := range ch {
		log.Println(string(body))
	}
}
