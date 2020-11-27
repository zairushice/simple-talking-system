package contoller

import (
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/service/process"
	"simple-talking-system/service/utils"
)

type Controller struct {
	Conn net.Conn
}

func (this *Controller) ServerProcessMessage(msg *message.Message) (err error) {
	up := &process.UserProcessor{Conn: this.Conn}
	switch msg.Type {
	case message.LoginMsgType:
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		err = up.ServerProcessRegister(msg)
	default:
		fmt.Println("invalid message type")
	}
	return
}

func (this *Controller) Process2() (err error) {
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		msg, err := tf.ReadBytes()
		if err != nil {
			fmt.Println("readBytes error:", err)
			return err
		}
		err = this.ServerProcessMessage(&msg)
		if err != nil {
			return err
		}
	}

}
