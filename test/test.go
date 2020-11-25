package main

import "fmt"

func main() {
	var i int
	for {
		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			fmt.Println(err)
		}
	}

}
