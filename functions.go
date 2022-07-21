package main

import "fmt"

func KevlinToCellsius(k float64) float64 {
	k -= 273.15
	return k

}
func cesluisToFahrenheit(f float64) float64 {
	f = (f * 9.0 / 5.0) + 32
	return f
}
func kelvinToFahenrenheit(m float64)float64{
	m -= 459.67
	return m
}
func main() {
	kevlin := 233.0
	celsuis := KevlinToCellsius(kevlin)
	fmt.Println(kevlin, " k is ", celsuis, " C")

	celsuis = 233.0
	fahenrenheit := cesluisToFahrenheit(celsuis)
	fmt.Println(celsuis, "C is ", fahenrenheit, "F")

	kevlin = 0.0
	fahenrenheit = kelvinToFahenrenheit(kevlin)
	fmt.Println(kevlin, "K is ", fahenrenheit, "F")

	kev := 0.0
	fahenrenheit = cesluisToFahrenheit(KevlinToCellsius(kev))
	fmt.Println(kevlin, "K is ", fahenrenheit, "F")
}
