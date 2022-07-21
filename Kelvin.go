package main

import (
	"fmt"

)
func kevlinToCellsius(k float64) float64{
	k -= 273.15
	return k

}
func main()  {
	kevlin := 294.0
	celsuis := kevlinToCellsius(kevlin)
	fmt.Println(kevlin," k is ",celsuis," C")
}
