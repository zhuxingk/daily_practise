package main

import "fmt"

//
//func main()  {
//	var board [8][8]string//创建一个8*8嵌套数组，其中数组的每个元素都是一个字符串
//	board[0][0] = "r"
//	board[0][7] = "r"//将“车”（rook）放置到指定的坐标上
//	for column := range board[1] {
//		board[1][column] = "p"
//	}
//	fmt.Println(board)
//}
func display(board [8][8]rune)  {
	for _,row := range board{
		for _,column := range row{
			if column == 0{
				fmt.Print(" ")
			}else {
				fmt.Printf("%c ",column)
			}
		}
		fmt.Println()
	}
}
func main()  {
	var board [8][8]rune

	//摆放黑棋
	board[0][0] ='r'
	board[0][1] ='n'
	board[0][2] ='b'
	board[0][3] ='q'
	board[0][4] ='k'
	board[0][5] ='b'
	board[0][6] ='b'
	board[0][7] ='r'


	// 摆放“兵”
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	// 摆放白棋
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	display(board)
}