package main

import "fmt"

func main()  {
	message := "shalom"
	for i:=0; i<6;i++{
		c := message[i]
		fmt.Printf("the %c word is %c\n", i,c)
	}

}
