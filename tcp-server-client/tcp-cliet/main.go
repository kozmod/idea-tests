package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte("world"))
	if err != nil {
		log.Fatal(err)
	}
}
