package main

import (
	"builder-pattern/server"
	"log"
	"os"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	svr := server.NewBuilder("localhost", 8080).Timeout(30).Logger(logger).Build()

	if err := svr.Start(); err != nil {
		log.Fatalln(err)
	}

}
