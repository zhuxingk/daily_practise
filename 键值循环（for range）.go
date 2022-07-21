package main
import "fmt"
func main(){
	//遍历数组、切片
	for key1, value1 := range []int{1,2,3,4}{
		fmt.Printf("key1: %d\n, value1: %d\n",key1,value1)
	}
	//遍历字符串
	var str ="hello world"
	for key2,value2 := range str{
		fmt.Printf("key2 :%d\n,value2:%d\n",key2 ,value2)
	}
	//遍历map
	m := map[string]int{
		"hello world"  :100,
		"你好" :200,
	}
	for key,value := range m{
		fmt.Println(key,value)
	}
	//遍历通道
	c := make(chan int)
	go func(){//启动goroutine
		c <- 1
		c <- 3
		c <- 2
		close(c)
	}()
	for v := range c{
		fmt.Println(v)
	}
}
