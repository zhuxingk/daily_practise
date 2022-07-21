package main
import "fmt"
func main(){
	var a = "hello"
	switch a{
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(3)
		}

	var b ="dad"
	switch b {
	case "mum","dad":
		fmt.Println("family")

	}
	var r int  =11
	switch {
	case r > 10 && r <20:
		fmt.Println(r)
	}
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough//连续执行case语句
	case s != "world":
		fmt.Println("world")
	}
}
