package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
)

func writeBytes(bytes []byte, network string, address string) (err error) {
	conn, err := net.Dial(network, address)
	defer conn.Close()
	if err != nil {
		fmt.Println("dial error:", err)
		return err
	}
	n, err := conn.Write(bytes)
	if err != nil {
		fmt.Println("write error:", err)
		return err
	}
	fmt.Println("write bytes:", n)
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
	fmt.Println(string(bytes))
	fmt.Println(msgLength)
	_ = writeBytes(lenBytes, "tcp", "localhost:8888")

	return
}
