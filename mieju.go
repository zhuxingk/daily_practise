package main

import (
	"fmt"

)

	type ChipType int
	const (
		None ChipType = iota
		CPU
		GPU
	)
	func (c ChipType) string() string {
		switch c {
		case None :
			return "None"
		case CPU :
			return "CPU"
		case GPU :
			return "GPU"	
	}
	return "N/A"
	}
	func main(){
		fmt.Printf("%s %d",CPU, CPU)
	}
