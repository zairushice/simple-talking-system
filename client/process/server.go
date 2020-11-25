package process

import (
	"fmt"
	"net"
	"os"
	"simple-talking-system/client/utils"
)

func ShowMenu() {
	fmt.Println("*********LOGIN SUCCESS!!*********")
	fmt.Println("*********1.LIST ALL ONLINE USERS*********")
	fmt.Println("*********2.SEND MESSAGE*********")
	fmt.Println("*********3.LIST CHAT RECORD*********")
	fmt.Println("*********4.EXIT*********")
	fmt.Println("*********CHOOSE<1-4>*********")
	var key int
	_, err := fmt.Scanf("%d\n", &key)
	if err != nil {
		fmt.Println(err)
	}
	switch key {
	case 1:
		fmt.Println("list all online users")
	case 2:
		fmt.Println("turn to chat box")
	case 3:
		fmt.Println("list chat record")
	case 4:
		os.Exit(0)
	default:
		fmt.Println("invalid choose")
	}
}

func processServerMsg(conn net.Conn) {
	tf := utils.Transfer{Conn: conn}
	for {
		fmt.Println("后台正在监听服务器端")
		msg, err := tf.ReadBytes()
		if err != nil {
			fmt.Println("processServerMsg read messsage error:", err)
			return
		}
		fmt.Println("processServerMsg read message:", msg)
	}

}
