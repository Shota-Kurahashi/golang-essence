package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.Copy(w, f)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHandler(w, r)
		case http.MethodPost:
			postHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
