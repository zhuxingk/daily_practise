package main

import (
	"fmt"
	"math/big"
)

func main()  {
	lightspeed := big.NewInt(299792)
	secondPerday := big.NewInt(86400)
	distance := new(big.Int)
	distance.SetString("24000000000000000000",10)
	fmt.Println("Andromeda Galaxy is",distance,"km away.")
	seconds := new(big.Int)
	seconds.Div(distance,lightspeed)
	days := new(big.Int)
	days.Div(seconds,secondPerday)
	fmt.Println("That is",days,"days of travel at light speed.")
}
