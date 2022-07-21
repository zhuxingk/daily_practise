package main

import (
	"fmt"
	"math"
	"os"
)

func main()  {
	var bh float64 = 32765
	var h = int16(bh)//todo
	if bh < math.MinInt16 || bh > math.MaxInt16{
		fmt.Println("error")
		os.Exit(0)
	}
	fmt.Println(h)
}
