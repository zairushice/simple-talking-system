package processor

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/common/utils"
)

type UserProcessor struct {
	Conn net.Conn
}

func (this UserProcessor) ServerProcessLogin(msg *message.Message) (err error) {
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("unmarshal message.Data error:", err)
	}
	loginResMsg := new(message.LoginResMsg)
	if loginMsg.UserId == 200 && loginMsg.UserPassword == "selangwudi" {
		loginResMsg.Code = 200
		loginResMsg.Error = "success!!"
	} else {
		loginResMsg.Code = 500
		loginResMsg.Error = "unmatched userid and password"
	}
	marshal, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("marshal login response message.Data error:", err)
	}
	resMsg := new(message.Message)
	resMsg.Type = message.LoginResMsgType
	resMsg.Data = string(marshal)
	resMsgBytes, err := json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal login response message error:", err)
	}
	err = tf.WriteBytes(resMsgBytes)
	if err != nil {
		fmt.Println("send login response message error:", err)
	}
	fmt.Println("send back response message success!!")
	return

}
