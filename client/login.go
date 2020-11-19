package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
)

func writeBytes(conn net.Conn, bytes []byte) (err error) {
	n, err := conn.Write(bytes)
	if err != nil {
		fmt.Println("write bytes error:", err)
	}

	fmt.Println("write bytes=", n)
	return
}

func login(userId int, passWord string) (err error) {
	msg := message.Message{
		Type: message.LoginMsgType,
		Data: message.LoginMsg{
			UserId:       userId,
			UserPassword: passWord,
			UserName:     "test1",
		},
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal error:", err)
		return err
	}
	msgLength := uint32(len(bytes))
	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, msgLength)

	conn, err := net.Dial("tcp", "192.168.50.81:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	err = writeBytes(conn, lenBytes)
	if err != nil {
		fmt.Println("write message length error:", err)
		return
	}
	err = writeBytes(conn, bytes)
	if err != nil {
		fmt.Println("write message error:", err)
	}
	fmt.Println("message length=", msgLength)
	fmt.Println("message content=", string(bytes))

	return
}
