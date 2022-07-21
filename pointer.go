package main

import "fmt"

func main(){
	var house = "Malibu Point 10880, 90625"
	ptr := &house
	fmt.Printf("ptr type:%T \n",ptr)
	fmt.Printf("address:%p\n",ptr)
	value := *ptr
	fmt.Printf("value type:%T\n",value)
	fmt.Printf("value:%s\n",value)
}