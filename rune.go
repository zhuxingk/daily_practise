package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Print(string(pi),string(alpha),string(omega),string(bang))


	countdown := 10
	str := "Launch in T minus" + strconv.Itoa(countdown) + " seconds"
	fmt.Println(str)

	countdown1 := 9
	str1 := fmt.Sprintf("Launch  in T minus %v seconds",countdown1)//作用与 Printf 函数基本相同，唯一的区别在于 Sprintf 函数会返回格式化之后的 string 而不是打印它
	fmt.Println(str1)

	countdown2, err := strconv.Atoi("10")//字符串转换为数值
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(countdown2)
}
