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

func (this *Transfer) WriteBytes(bytes []byte) (err error) {
	msgLength := uint32(len(bytes))
	binary.BigEndian.PutUint32(this.Buf[:4], msgLength)
	n, err := this.Conn.Write(this.Buf[:4])
	if err != nil {
		return
	}
	n, err = this.Conn.Write(bytes)
	if err != nil || n != int(msgLength) {
		return
	}
	return
}

func (this Transfer) ReadBytes() (msg message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err == io.EOF {
		fmt.Println("client has closed the connection")
		return
	} else if err != nil {
		return
	}
	lengthUint := binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:lengthUint])
	if err != nil || n != int(lengthUint) {
		fmt.Println("read message error:", err)
		return
	}

	err = json.Unmarshal(this.Buf[:lengthUint], &msg)
	if err != nil {
		return
	}

	return
}
