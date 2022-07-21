package main

import "fmt"

func main()  {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge + earthDays / marsDays
	fmt.Println("I am",marsAge,"years old on MArs!")
	fmt.Println(int(earthDays))

	red := 255
	red_conv := uint8(red)
	fmt.Println(red_conv)
	//fmt.Print(age > marsAge)
}