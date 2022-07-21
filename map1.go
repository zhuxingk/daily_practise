package main
import "fmt"
func main() {
	scene := make(map[string]int)
	//key类型为string,value类型为int
	scene["route"] = 66
	scene["brazil"] = 64
	scene["china"] = 960
	/*map填充内容的声明方式
	m := map[string]int{
	"W": "forward"
	"A": "left"
	"S": "right"
	"D": "backward"
	}
	 */
	fmt.Println(scene["route"])
	v1 := scene["route2"]
	fmt.Println(v1)
	//使用for range循环对map进行遍历
	for k ,v := range scene{
		fmt.Println(k,v)
	}

}
