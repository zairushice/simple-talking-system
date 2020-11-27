package process

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/service/model"
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
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPassword)
	if err != nil {
		if err == model.ErrorWrongPassword {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else if err == model.ErrorNoSuchUserId {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "server internal error"
		}
	} else {
		loginResMsg.Code = 200
		loginResMsg.Error = "login success!!"
		fmt.Println(user, "login success!!")
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
