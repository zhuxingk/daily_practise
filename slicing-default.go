package main

import (
	"fmt"
)

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	allplanets := planets[:]
	fmt.Println(allplanets)

	/*使用切片切分数组，同样可以用来切分字符串
	 */

	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)

	neptune = "Poseidon"
	fmt.Println(tune) // ←--- 打印出“tune”

	question := "¿Cómo estás?"
	fmt.Println(question[:6]) // ←--- 打印出“¿Cóm”

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Printf("%T %T %T", dwarfArray, dwarfSlice, dwarfs)
	//terrestial := planets[0:4]
	//gasgiants := planets[4:6]
	//icegiants := planets[6:8]
	//fmt.Println(terrestial,gasgiants,icegiants)
	//
	//fmt.Println(gasgiants[0])
	//
	//giants := planets[4:8]
	//gas := giants[0:2]
	//ice := giants[2:4]
	//fmt.Println(giants,gas,ice)
	//
	//icegiantsmarkII := icegiants// ←--- 复制 iceGiants切片（指向planets数组的视图）
	//icegiants[1] = "Poesidon"
	//fmt.Println(planets)// ←--- 打印出“[Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]”
	//fmt.Println(icegiants,icegiantsmarkII,ice)// ←--- 打印出“[Uranus Poseidon] [Uranus Poseidon][Uranus Poseidon]”
}
