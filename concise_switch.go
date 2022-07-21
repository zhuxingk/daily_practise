package main

import "fmt"

func main()  {
	fmt.Println("there is a cavern entrance here and a path to the east")
	command := "go inside"
	switch command {
	case "go east":
		fmt.Println("you head further up the mountain.")
	case "enter cave","go inside":
		fmt.Println("you find yourself in a dimly lit cavern")
	case "read sign":
		fmt.Println("The sign reads 'no minors'")
	default:
		fmt.Println("didn't quite get there")
	}
	room := "lake"
	switch  {
	case room == "cave":
		fmt.Println("you find yourself in a dimly cavern")
	case room == "lake":
		fmt.Println("the ice seems solid enough")
		fallthrough
	case room == "underwater":
		fmt.Println("the water is freezing cold")
	}
}