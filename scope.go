package main

import "fmt"
import "math/rand"
func main(){
	count := 0
	for count < 10{
		num := rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}
	//fmt.Println(num)
}