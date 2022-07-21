package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)
func main(){
	tip1 := "gnejia is a nijia"
	fmt.Println(len(tip1))
	tip2 := "忍者"
	fmt.Println(len(tip2))
	//获取ASCII码字符
	fmt.Println(utf8.RuneCountInString("忍者"))
	fmt.Println(utf8.RuneCountInString("hold the fire"))
	//获取unicode字符串长度



	//遍历字符串，遍历ascii码
	theme := "sniper start"
	for i := 0;i < len(theme); i++{
		fmt.Printf("ascii: %c %d\n", theme[i],theme[i])
	}
	//遍历字符串，遍历unicode码
	theme1 := "sniper start"
	for _,s := range theme1{
		fmt.Printf("unicode: %c %d\n", s, s)
	}



	//获取字符串中的某一段字符
	tracer := "god is so excellent !"
	comma := strings.Index(tracer, "g")
	pos := strings.Index(tracer[comma:],"o")
	fmt.Println(comma,pos,tracer[comma+pos:])



	//修改字符串
	angel := "heros never die"
	angleBytes := []byte(angel)
	for i := 5;i <= 10;i++{
		angleBytes[i] = ' '
	}


	fmt.Println(string(angleBytes))
	//连接字符串
	hammer := "吃我一锤"
	sickle := "死吧"
	//声明字节缓冲
	var stringBuilder bytes.Buffer
	//将字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	//将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())


	//字符串格式化
	var progress = 2
	var target = 8
	//参数格式化
	title := fmt.Sprintf("已采集%d个药草，还需要完成%d个任务", progress, target)
	fmt.Println(title)
	pi := 3.14159
	variant := fmt.Sprintf("%v %v %v","月球基地",pi,true)
	fmt.Println(variant)
	//匿名结构体声明，并赋予初值
	profile := &struct{
		Name string
		HP int
	}{
		Name: "rat",
		HP: 150,
	}
	fmt.Printf("使用‘%%+v’ %+v\n",profile)
	fmt.Printf("使用‘%%#v’ %#v\n",profile)
	fmt.Printf("使用‘%%T’ %T\n",profile)
}
