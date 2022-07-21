package main

import "fmt"
func main(){
	command := "go east"
	if command == "go east" {
		fmt.Println("you head further up the mountain.")
	}else if command == "go inside"{

		fmt.Println("you enter the cave where you live out the rest of your life.")
	}else {
		fmt.Println("didn't quite get that.")
	}
}


