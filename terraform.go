package main

import "fmt"

// Planets类型会将方法绑定到[]string类型
type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)  // ←--- 打印出“[Mercury Venus Earth New Mars Jupiter Saturn New Uranus New Neptune]”
}
