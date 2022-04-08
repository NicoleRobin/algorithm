package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("net")
	listener, err := net.Listen("tcp", ":9988")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	defer conn.Close()
	data := []byte{}
	count, err := conn.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("recv data:%s, len:%d\n", string(data), count)
}
