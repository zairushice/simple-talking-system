package main

import (
	"fmt"
	"os"
)

var userId int
var passWord string

func main() {
	var key int
	var loop = true

	for loop {
		fmt.Println("*********Welcome to CHAT ROOM*********")
		fmt.Println("\t\t\t1.Login")
		fmt.Println("\t\t\t2.Sign up")
		fmt.Println("\t\t\t3.Exit")
		fmt.Println("\t\t\tChoose<1-3>")

		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println(err)
		}
		switch key {
		case 1:
			{
				fmt.Println("turn to login")
				loop = false
			}
		case 2:
			{
				fmt.Println("turn to sign up")
				loop = false
			}
		case 3:
			{
				os.Exit(0)
			}
		default:
			fmt.Println("invalid input, please try again")
		}
	}

	if key == 1 {
		fmt.Print("userId:")
		_, err := fmt.Scanf("%d\n", &userId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("passWord:")
		_, err = fmt.Scanf("%s\n", &passWord)
		if err != nil {
			fmt.Println("err")
		} else {
			fmt.Println("转入用户名密码校验逻辑")
			_ = login(userId, passWord)
		}

	}

}
