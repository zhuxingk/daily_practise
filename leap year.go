package main

import "fmt"

func main()  {
	fmt.Println("this year is 2100,should you leap?")
	year := 2100
	leap := year % 400 == 0 || (year%4==0 && year%100!=0)
	if leap{
		fmt.Println("look before you leap")
	}else {
		fmt.Println("keep your feet on the ground")
	}
}