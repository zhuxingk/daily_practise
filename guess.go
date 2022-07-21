package main

import (

	"math/rand"
	"fmt"
)

func main()  {
	number := 5
	var guessnum int
	for {
		guessnum = rand.Intn(100)
		fmt.Println(guessnum)
		if guessnum == number{
			fmt.Println("bingo!")
			break
		}else if guessnum > number {
			fmt.Println("more than number！")
		}else if guessnum < number{
			fmt.Println("little than number")
		}else{
			fmt.Println("err")
		}
	}
}