package process

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/service/utils"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this *UserProcessor) ServerProcessLogin(msg *message.Message) (err error) {
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WriteBytes(resMsgBytes)
	if err != nil {
		fmt.Println("send login response message error:", err)
		return
	}
	fmt.Println("successfully send back login response message to client!!")
	return
}

func (this *UserProcessor) ServerProcessRegister(msg *message.Message) (err error) {
	return
}
