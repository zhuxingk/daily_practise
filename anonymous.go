package main

import "fmt"

func main()  {
	func() {//声明匿名函数
		fmt.Println("functions anonymous")
	}()//调用匿名函数
}
