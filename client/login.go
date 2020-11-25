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
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal message.Data error:", err)
	}
	msg.Data = string(dataBytes)
	bytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal error:", err)
		return err
	}

	conn, err := net.Dial("tcp", "192.168.10.230:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	tf := utils.Transfer{Conn: conn}
	err = tf.WriteBytes(bytes)
	if err != nil {
		fmt.Println("write bytes error:", err)
		return
	}
	resMsg, err := tf.ReadBytes()
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
