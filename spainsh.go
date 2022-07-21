package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	question := "¿Cómo estás?"
	fmt.Println(len(question),"bytes")
	fmt.Println(utf8.RuneCountInString(question))//以符文而不是以字节为单位返回字符串的长度
	c, size := utf8.DecodeRuneInString(question)//解码字符串的首个字符并返回解码后的符文占用的字节数量。
	fmt.Printf("first rune:%c %v bytes",c,size)

	for i,c := range question{
		fmt.Printf("%v %c\n",i,c)
	}
	test := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(utf8.RuneCountInString(test))
	test1 := "¿"
	fmt.Println(utf8.DecodeRuneInString(test1))

}
