package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read length error:", err)
		return
	}
	lengthUint := binary.BigEndian.Uint32(buf[:4])
	fmt.Println("message length=", lengthUint)

	n, err := conn.Read(buf[:lengthUint])
	if err != nil {
		fmt.Println("read message error:", err)
		return
	}

	fmt.Println("read message length=", n)
	msgBytes := buf[:lengthUint]

	err = json.Unmarshal(msgBytes, &msg)
	if err != nil {
		fmt.Println("message unmarshal error:", err)
		return
	}
	fmt.Println("msg=", msg)

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
