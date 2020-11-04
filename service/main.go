package main

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	readBytes := make([]byte, 8096)
	n, err := conn.Read(readBytes)
	if err != nil {
		fmt.Println("read bytes error:", err)
		return
	}
	fmt.Println("read bytes n=", n)
	err = json.Unmarshal(readBytes[:n], &msg)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println("unmarshal result=", msg)
	return
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := readPkg(conn)
		if err != nil {
			fmt.Println("readPkg error:", err)
			return
		}
	}

}

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
		go process(conn)

	}

}
