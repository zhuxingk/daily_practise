package main

import (
	"fmt"
	"sort"
)
func main() {
	scene := make(map[string]int)//创建map实例
	//准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	//声明一个切片保存数据
	var sceneListone []string
	//将map数据遍历复制到slice中
	for k := range scene{
		sceneListone = append(sceneListone,k)
	}
	//对slice进行排序
	sort.Strings(sceneListone)
	//输出
	fmt.Println(sceneListone)
	//删除map键值对中的值

	delete(scene,"brazil")

	for k, v := range scene{
		fmt.Println(k,v)
	}
}