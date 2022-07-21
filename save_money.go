package main

import (
	"fmt"
	"math/rand"
)
var save_money float64

func main()  {
	save_money = 0.0
	fmt.Printf("%5v\n","extra")
	for save_money < 20.0{
		switch rand.Intn(3) {
		case 1:
			save_money += 0.05

		case 2:
			save_money += 0.1

		default:
			save_money += 0.25

		}
		fmt.Printf("$%5.2f\n",save_money)
	}
	}
