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

	conn, err := net.Dial("tcp", "192.168.50.59:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	err = utils.WriteBytes(conn, bytes)
	if err != nil {
		fmt.Println("write message error:", err)
	}

	resMsg, err := utils.ReadBytes(conn)
	if err != nil {
		fmt.Println("read message error:", err)
		return
	}
	loginResMsg := new(message.LoginResMsg)
	err = json.Unmarshal([]byte(resMsg.Data), loginResMsg)
	if err != nil {
		fmt.Println("unmarshal login response message error:", err)
		return
	}
	if loginResMsg.Code == 200 {
		fmt.Println("successfully login!!")
	} else {
		fmt.Printf("error code:%v, error message:%v\n", loginResMsg.Code, loginResMsg.Error)
	}

	return
}
