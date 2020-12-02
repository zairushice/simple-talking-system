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
	err = json.Unmarshal(dataBytes, loginMsg)
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
		fmt.Println(*user, "login success!!")
	}
	loginResMsgBytes, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("marshal login response message.Data error:", err)
		return
	}
	resMsg.Data = string(loginResMsgBytes)
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
	dataBytes := []byte(msg.Data)
	registerData := new(message.RegisterMsg)
	err = json.Unmarshal(dataBytes, registerData)
	if err != nil {
		return
	}
	err = model.MyUserDao.Register(&registerData.User)
	resMsg := new(message.Message)
	resMsg.Type = message.RegisterResMsgType
	registerResMsg := new(message.RegisterResMsg)
	if err != nil {
		if err == model.ErrorExistUserId {
			registerResMsg.Code = 505
			registerResMsg.Error = model.ErrorExistUserId.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "register unknown error:" + err.Error()
		}
	} else {
		registerResMsg.Code = 200
		registerResMsg.Error = "注册成功!!"
	}

	marshal, err := json.Marshal(registerResMsg)
	if err != nil {
		return
	}
	resMsg.Data = string(marshal)
	resMsgBytes, err := json.Marshal(resMsg)
	if err != nil {
		return
	}
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WriteBytes(resMsgBytes)
	if err != nil {
		fmt.Println("write registerResMsg error:", err)
		return
	}
	return

}
