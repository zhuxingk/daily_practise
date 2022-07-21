package main

import (
	"fmt"
	"sync"
)
func main(){
	/*
	//创建map实例，一个从int到int的映射
	m := make(map[int]int)
	//开启一段并发代码
	go func(){
		//不停输入
		for {
			m[1] = 1
		}
	}()
	go func(){
		//不停输入
		for {
			_ = m[1]
		}
	}()
	for{

	}
	 */
	var scene sync.Map
	//将keyvalue对保存到sync.map
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)
	//从sync.map中根据键取值
	fmt.Println(scene.Load("london"))
	//delete key value
	scene.Delete("london")
	//遍历所有sync.map中的键值对
	scene.Range(func(k,v interface{}) bool{
		fmt.Println("iterate:", k, v)
		return true
	})
}