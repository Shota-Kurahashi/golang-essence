package main

import "fmt"

func main() {

	// warning: Printf format %s reads arg #2, but call has 1 arg
	fmt.Println("%s %s", "Hello")
}
