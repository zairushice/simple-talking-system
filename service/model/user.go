package model

type User struct {
	UserId       int    `json:"userId"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName"`
}
