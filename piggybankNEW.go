package main

import (
	"fmt"
	"math/rand"
)
var piggybank int64

func main()  {
	piggybank = 0
	fmt.Printf("%5v\n","extra")
	for piggybank < 2000{
		switch rand.Intn(3) {
		case 1:
			piggybank += 5

		case 2:
			piggybank += 10

		default:
			piggybank += 25

		}
		dollars := piggybank / 100
		cents := piggybank % 100
		fmt.Printf("$%d.%02d\n",dollars,cents)
	}
}
