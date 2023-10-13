package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
