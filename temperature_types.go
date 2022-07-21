package main

import "fmt"
type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius  {
	return celsius(k - 273.15)
}
func celeiusToKelvin(m celsius) kelvin{
	return kelvin(m + 273.15)
}
func main()  {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k,"k is ",c ,"C")

	var n celsius = 127.0
	m := celeiusToKelvin(n)
	fmt.Println(n,"C is ",m,"K")
}




//
//import "fmt"
//
//type celsius float64
//type kelvin float64
//
//// kelvinToCelsius converts ºK to ºC
//func kelvinToCelsius(k kelvin) celsius {
//	return celsius(k - 273.15)  // ←--- 类型转换是必需的
//}
//func main() {
//	var k kelvin = 294.0  // ←--- 实参必须为kelvin类型
//	c := kelvinToCelsius(k)
//	fmt.Print(k, "ºK is ", c, "ºC")
//}
