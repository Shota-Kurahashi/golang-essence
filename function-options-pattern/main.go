package main

import (
	"function-options-pattern/server"
	"log"
	"os"
)

func main() {
	f, err := os.Create("server.log")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	svr := server.New("localhost", 8080, server.WithLogger(logger), server.WithTimeout(30))

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}

}
