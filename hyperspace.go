package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string)  {
	for i := range worlds{
		worlds[i] =strings.TrimSpace(worlds[i])

	}
}
func main()  {
	planets := []string{"Venus","Earth","MArs"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets,""))
}
