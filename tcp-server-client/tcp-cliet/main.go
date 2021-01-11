package main

import "net"

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hello"))
	conn.Write([]byte("world"))
}
