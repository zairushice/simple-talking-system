package main

import (
	"fmt"
	"net"
	"simple-talking-system/service/processor"
)

func mainProcess(conn net.Conn) {
	defer conn.Close()
	p := processor.Processor{Conn: conn}
	err := p.Process2()
	if err != nil {
		return
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	fmt.Println("start listening on ", listen.Addr())
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		fmt.Println("loop accept message...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		} else {
			fmt.Println("get message from ", conn.RemoteAddr())
			go mainProcess(conn)
		}

	}

}