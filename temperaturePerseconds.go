package main

import (
	"fmt"
	"math/rand"
	"time"
)
type KELvin float64
func measureTemperature(samples int,sensor func() KELvin){//mesaure 函数接受另一个函数作为他的第二个参数
	for i :=0; i<  samples;i++{
		k := sensor()
		fmt.Printf("%v K\n",k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() KELvin{
	return KELvin(rand.Intn(151) + 150)
}
func main(){
	measureTemperature(3,fakeSensor)//把函数的名字传给另一个函数
}
