package main

import (
	"fmt"
	"os"
	"simple-talking-system/client/process"
)

var userId int
var passWord string
var userName string

func main() {
	var key int

	for {
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
				fmt.Print("userId:")
				_, err := fmt.Scanf("%d\n", &userId)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Print("passWord:")
				_, err = fmt.Scanf("%s\n", &passWord)
				if err != nil {
					fmt.Println(err)
					return
				} else {
					up := &process.UserProcess{}
					err := up.Login(userId, passWord)
					if err != nil {
						fmt.Println(err)
						return
					}
				}

			}
		case 2:
			{
				fmt.Print("create a userId:")
				_, err := fmt.Scanf("%d\n", &userId)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Print("create a password:")
				_, err = fmt.Scanf("%s\n", &passWord)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Print("create a nickname:")
				_, err = fmt.Scanf("%s\n", &userName)
				if err != nil {
					fmt.Println(err)
				}
				up := &process.UserProcess{}
				err = up.Register(userId, passWord, userName)
				if err != nil {
					fmt.Println(err)
				}

			}
		case 3:
			{
				os.Exit(0)
			}
		default:
			fmt.Println("invalid input, please try again")
		}
	}

}
