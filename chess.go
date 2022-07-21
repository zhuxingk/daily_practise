package main

import "fmt"

func main()  {
	var board [8][8]string//创建一个8*8嵌套数组，其中数组的每个元素都是一个字符串
	board[0][0] = "r"
	board[0][7] = "r"//将“车”（rook）放置到指定的坐标上
	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Println(board)
}
