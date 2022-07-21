package main

import (
	"fmt"
	"math/rand"
)

type KelVin float64

func fakesensor() KelVin  {
	return KelVin(rand.Intn(151) + 150)
}
func realsensor() KelVin  {
	return 0//待办事项，实现传感器
}
func main(){
	sensor := fakesensor//函数值赋值给变量
	fmt.Println(sensor())
	sensor = realsensor
	fmt.Println(sensor())

}