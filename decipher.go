package main

import "fmt"

func main()  {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0
	for i := 0;i < len(cipherText);i++{
		//A=0,.....,Z=25
		c := cipherText[i] -'A'
		k :=keyword[keyIndex] - 'A'
		//fmt.Println(c,k)
		//加密字母，关键字字母
		c = (c-k+26)%26 + 'A'
		message += string(c)
	//	fmt.Println(c)
		//对keyIndex自增
		keyIndex++
		keyIndex %= len(keyword)
		//fmt.Println(keyIndex,len(keyword))t
	}
	fmt.Println(message)
}
