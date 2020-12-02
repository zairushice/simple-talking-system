package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"simple-talking-system/common/message"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{Pool: pool}
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	userInfo, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorNoSuchUserId
		}
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(userInfo)
	if err != nil {
		fmt.Println("base64 decode error:", err)
		return
	}
	user = new(User)
	err = json.Unmarshal(decodeBytes, user)
	if err != nil {
		fmt.Println("unmarshal userInfo error:", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPassword string) (user *User, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPassword != userPassword {
		err = ErrorWrongPassword
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ErrorExistUserId
		return
	}
	bytes, err := json.Marshal(user)
	encoding := base64.StdEncoding.EncodeToString(bytes)
	_, err = conn.Do("hset", "users", user.UserId, encoding)
	if err != nil {
		fmt.Println("write redis error:", err)
		return err
	}
	return
}
