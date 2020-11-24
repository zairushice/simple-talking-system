package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"simple-talking-system/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) WriteBytes(conn net.Conn, bytes []byte) (err error) {
	msgLength := uint32(len(bytes))
	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, msgLength)
	n, err := conn.Write(lenBytes)
	if err != nil {
		fmt.Println("write message length error:", err)
		return
	}
	n, err = conn.Write(bytes)
	if err != nil || n != int(msgLength) {
		fmt.Println("write message error:", err)
		return
	}
	fmt.Println("write message length and message success!!")
	fmt.Printf("message length=%v\n", msgLength)
	fmt.Printf("message content=%v\n", string(bytes))
	return
}

func (t Transfer) ReadBytes(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err == io.EOF {
		fmt.Println("client has closed the connection")
		return
	} else if err != nil {
		fmt.Println("read message length error:", err)
	}
	lengthUint := binary.BigEndian.Uint32(buf[:4])
	fmt.Println(lengthUint)

	n, err := conn.Read(buf[:lengthUint])
	if err != nil || n != int(lengthUint) {
		fmt.Println("read message error:", err)
		return
	}

	fmt.Println("read message length and message success!!")
	fmt.Printf("read meassage length=%v\n", lengthUint)
	fmt.Printf("read message=%v\n", string(buf[:lengthUint]))

	err = json.Unmarshal(buf[:lengthUint], &msg)
	if err != nil {
		fmt.Println("message unmarshal error:", err)
		return
	}
	fmt.Println("msg=", msg)

	return
}
