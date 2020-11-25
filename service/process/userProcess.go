package process

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
)

func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
	dataBytes := []byte(msg.Data)
	loginMsg := new(message.LoginMsg)
	err = json.Unmarshal(dataBytes, &loginMsg)
	if err != nil {
		fmt.Println("unmarshal message.Data error:", err)
		return
	}

	resMsg := new(message.Message)
	resMsg.Type = message.LoginResMsgType
	loginResMsg := new(message.LoginResMsg)
	if loginMsg.UserId == 200 && loginMsg.UserPassword == "selangwudi" {
		loginResMsg.Code = 200
		loginResMsg.Error = "success"

	} else {
		loginResMsg.Code = 500
		loginResMsg.Error = "userId does not exist"
	}
	loginResMsgBytes, err := json.Marshal(loginResMsg)
	resMsg.Data = string(loginResMsgBytes)
	if err != nil {
		fmt.Println("marshal login response message.Data error:", err)
		return
	}
	resMsgBytes, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal login response message error:", err)
	}
	err = utils.WriteBytes(conn, resMsgBytes)
	if err != nil {
		fmt.Println("send login response message error:", err)
		return
	}
	fmt.Println("successfully send back login response message to client!!")
	return
}

func serverProcessRegister(conn net.Conn, msg *message.Message) (err error) {
	return
}

func serverProcessMessage(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		err = serverProcessLogin(conn, msg)
	case message.RegisterMsgType:
		err = serverProcessRegister(conn, msg)
	default:
		fmt.Println("invalid message type")
	}
	return
}
