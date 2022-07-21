package main

import "fmt"

func main()  {
	yesNO := "yes"
	var launch bool
	switch yesNO {
	case "true", "yes","1":
		launch = true
	case "false","no","0":
		launch = false
	default:
		fmt.Println(yesNO,"is not valid")
	}
	fmt.Println("Ready for launch:",launch)
}