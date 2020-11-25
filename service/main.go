package main

import (
	"fmt"
	"net"
	"simple-talking-system/service/processor"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	fmt.Println("start listening on ", listen.Addr())
	if err != nil {
		fmt.Println("listen error:", err)
	}
	for {
		fmt.Println("loop accept message...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		}
		fmt.Println("get message from ", conn.RemoteAddr())
		sp := processor.ServerProcessor{Conn: conn}
		err = sp.MainProcess()
	}

}
