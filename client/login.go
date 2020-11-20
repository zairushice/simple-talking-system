package main

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/common/utils"
)

func login(userId int, passWord string) (err error) {
	msg := message.Message{
		Type: message.LoginMsgType,
	}
	data := message.LoginMsg{
		UserId:       userId,
		UserPassword: passWord,
		UserName:     "test1",
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal message.data error:", err)
	}
	msg.Data = string(marshalData)

	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal error:", err)
		return err
	}

	conn, err := net.Dial("tcp", "192.168.50.81:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	err = utils.WriteBytes(conn, bytes)
	if err != nil {
		fmt.Println("write message error:", err)
	}

	return
}
