package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	fmt.Println("start listening...")

	for true {
		conn, err := listen.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("listem.Accept error:", err)
		}

	}
}