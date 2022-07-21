package main

import "fmt"
//sensor type function
type kelvin float64
type sensor func() kelvin
func realSensor() kelvin{
	return 0//待办事项，实现真正的传感器
}
func calibrate(s sensor,offset kelvin) sensor{
	return func() kelvin{//声明并返回匿名函数
		return s() + offset
	}
}
func main(){
	sensor := calibrate(realSensor,5)
	fmt.Println(sensor())
}