package processor

import (
	"fmt"
	"net"
	"simple-talking-system/common/message"
	"simple-talking-system/common/utils"
)

type ServerProcessor struct {
	Conn net.Conn
}

func (this *ServerProcessor) ServerProcessMsg(msg *message.Message) (err error) {
	up := UserProcessor{Conn: this.Conn}
	switch msg.Type {
	case message.LoginMsgType:
		err = up.ServerProcessLogin(msg)
	default:
		fmt.Println("invalid message type")
	}
	return
}

func (this *ServerProcessor) MainProcess() (err error) {
	defer this.Conn.Close()
	tf := utils.Transfer{Conn: this.Conn}
	for {
		msg, err := tf.ReadBytes()
		if err != nil {
			fmt.Println("readPkg error:", err)
			return err
		}
		err = this.ServerProcessMsg(&msg)
		if err != nil {
			return err
		}
	}
}
