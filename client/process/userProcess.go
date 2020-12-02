package process

import (
	"encoding/json"
	"fmt"
	"net"
	"simple-talking-system/client/utils"
	"simple-talking-system/common/message"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, passWord string) (err error) {
	msg := message.Message{
		Type: message.LoginMsgType,
	}
	data := message.LoginMsg{
		UserId:       userId,
		UserPassword: passWord,
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		return
	}
	msg.Data = string(marshalData)

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", "192.168.10.230:8888")
	if err != nil {
		return
	}
	defer conn.Close()
	tf := &utils.Transfer{Conn: conn}
	err = tf.WriteBytes(bytes)
	if err != nil {
		fmt.Println("write message error:", err)
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
		fmt.Println("登录成功!!")
		go processServerMsg(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Printf("error code:%v, error message:%v\n", loginResMsg.Code, loginResMsg.Error)
	}

	return
}

func (this *UserProcess) Register(userId int, password string, userName string) (err error) {
	registerMsg := message.Message{
		Type: message.RegisterMsgType,
	}
	user := message.User{
		UserId:       userId,
		UserPassword: password,
		UserName:     userName,
	}
	registerData := message.RegisterMsg{
		User: user,
	}
	registerDataBytes, err := json.Marshal(registerData)
	if err != nil {
		fmt.Println("marshal registerData error:", err)
		return
	}
	registerMsg.Data = string(registerDataBytes)
	bytes, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("marshal registerMsg error:", err)
		return
	}

	conn, err := net.Dial("tcp", "192.168.10.230:8888")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	tf := &utils.Transfer{Conn: conn}
	err = tf.WriteBytes(bytes)
	if err != nil {
		fmt.Println("write message error:", err)
	}

	resMsg, err := tf.ReadBytes()
	if err != nil {
		fmt.Println("read message error:", err)
		return
	}
	registerResMsg := new(message.RegisterResMsg)
	err = json.Unmarshal([]byte(resMsg.Data), registerResMsg)
	if err != nil {
		fmt.Println("unmarshal login response message error:", err)
		return
	}
	if registerResMsg.Code == 200 {
		fmt.Println("注册成功,请重新登录!!")
	} else {
		fmt.Printf("error code:%v, error message:%v\n", registerResMsg.Code, registerResMsg.Error)
	}

	return
}
