package main

import "fmt"

func main()  {
	launch := false
	launchtext :=fmt.Sprintf("%v\n",launch)
	fmt.Println("ready for launch",launchtext)

	var yesNo string
	if launch{
		yesNo = "yes"
	}else{
		yesNo = "no"
	}
	fmt.Println("Ready for launch",yesNo)

	launch1 := true
	var oneZero int
	if launch1{
		oneZero =1
	}else {
		oneZero = 0
	}
	fmt.Println("Reday for launch1: ",oneZero)
}