package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "0.0.0.0:9999")
	for {
		conn, _ := ln.Accept()
		go func(c net.Conn) {
			var allBytes []byte
			for {
				buf := make([]byte, 4)
				n, err := io.ReadFull(conn, buf)
				if err == io.EOF {
					break // client conn closed
				}
				fmt.Printf("[recv] len = %d, msg = %s\n", n, buf)
				allBytes = append(allBytes, buf[:n]...)
			}
			//allBytes, err := ioutil.ReadAll(conn)
			//if err != nil {
			//	fmt.Printf("error = %v\n", err)
			//}
			fmt.Printf("[recv_all] len = %d, cap = %d msg = %s\n", len(allBytes), cap(allBytes), allBytes)
		}(conn)
	}
}
