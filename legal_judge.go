package main

import (
	"fmt"
	"math"
	"os"
)

func main()  {
	var v =42

	if v <= math.MaxUint8 && v >= 0{
		v_conv := uint8(v)
		fmt.Println(v_conv)
	}else{
		fmt.Println("error")
		os.Exit(0)
	}
}