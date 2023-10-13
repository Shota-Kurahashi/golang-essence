package nets

import (
	"fmt"
	"log"
	"net"
)

func Main() {
	lis, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatal(err)
	}

	defer lis.Close()

	for {
		fmt.Println("Waiting for connection...")
		conn, err := lis.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleData(conn)
	}
}

func handleData(conn net.Conn) {
	defer conn.Close()
	var b [512]byte

	for {
		n, err := conn.Read(b[:])

		if err != nil {
			break
		}

		fmt.Println(string(b[:n]))
	}
}
