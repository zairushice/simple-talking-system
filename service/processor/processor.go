package processor

import (
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/service/process"
	"simple-talking-system/service/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMessage(msg *message.Message) (err error) {
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

func (this *Processor) Process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
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
