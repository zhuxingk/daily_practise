package main

import (
	"fmt"
	"math/rand"
)
var onedaysecond = 60*60*24

func main(){
	distance := 62100000
	trip_type := ""
	company := ""
	fmt.Printf("%-20v %5v %15v %20v\n", "Spaceline", "Days", "Trip", "Price")
   for count:=0;count<10;count++{
	   switch rand.Intn(3) {
	   case 1:
	   	company = "Space Adventure"
	   case 2:
	   	company = "Space"
	   default:
		   company = "Virgin Galactic"
	   }
	   speed :=  rand.Intn(16) +15
	   duration := distance / speed /onedaysecond
	   price := speed + 20.0
	   if rand.Intn(2) ==1{
	   	trip_type = "oneway"
	   }else {
	   	trip_type = "double_way"
	   	price = price * 2
		   }
	   fmt.Printf("%-20v %5v %15v %20v\n",company,duration,trip_type,price)
   }
}
