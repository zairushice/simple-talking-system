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
	loginResMsg := new(message.LoginResMsg)
	err = json.Unmarshal([]byte(resMsg.Data), loginResMsg)
	if err != nil {
		fmt.Println("unmarshal login response message error:", err)
		return
	}
	if loginResMsg.Code == 200 {
		fmt.Println("successfully login!!")
		go processServerMsg(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Printf("error code:%v, error message:%v\n", loginResMsg.Code, loginResMsg.Error)
	}

	return
}
