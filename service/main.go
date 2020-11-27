package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net"
	"simple-talking-system/service/contoller"
	"simple-talking-system/service/model"
	"time"
)

var pool *redis.Pool

func initPool(network string, address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(network, address)
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func process(conn net.Conn) {
	defer conn.Close()
	p := contoller.Controller{Conn: conn}
	err := p.Process2()
	if err != nil {
		return
	}
}

func main() {
	initPool("tcp", "192.168.10.94:6379", 16, 0, 300*time.Second)
	initUserDao()
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	fmt.Println("start listening on ", listen.Addr())
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		fmt.Println("loop accept message...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
		} else {
			fmt.Println("get message from ", conn.RemoteAddr())
			go process(conn)
		}

	}

}
