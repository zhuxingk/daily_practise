package main

import (
	"fmt"
)

func main()  {
	planets :=[...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestial := planets[0:4]
	gasgiants := planets[4:6]
	icegiants := planets[6:8]
	fmt.Println(terrestial,gasgiants,icegiants)

	fmt.Println(gasgiants[0])

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants,gas,ice)

	icegiantsmarkII := icegiants// ←--- 复制 iceGiants切片（指向planets数组的视图）
	icegiants[1] = "Poesidon"
	fmt.Println(planets)// ←--- 打印出“[Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]”
	fmt.Println(icegiants,icegiantsmarkII,ice)// ←--- 打印出“[Uranus Poseidon] [Uranus Poseidon][Uranus Poseidon]”
}
