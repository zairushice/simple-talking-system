package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"simple-talking-system/service/model"
)

func main() {
	data := model.User{
		UserId:       200,
		UserPassword: "selangwudi",
		UserName:     "测试1",
	}
	marshal, _ := json.Marshal(data)
	encoding := base64.StdEncoding.EncodeToString(marshal)
	fmt.Println(encoding)

	conn, _ := redis.Dial("tcp", "192.168.10.94:6379")
	//do, _ := conn.Do("hset", "users", 200, encoding)
	s, err := redis.String(conn.Do("hget", "users", 100))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println(err)
			return
		}

	}
	fmt.Println(s)

}
