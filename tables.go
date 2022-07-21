package main

import (

	"fmt"

)
//温度转换方法声明
type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}
type fahrenheit float64

func (f fahrenheit) celsius() celsius  {
	return celsius((f - 32.0) * 5.0 /9.0)

}
//格式打印变量声明，固定格式（const）
const (
	line = "================================="
	//rowFormat = "| %8s | %8s |\n"
	//numberFormat = "%.1f"
	rowFormat = "|  %8s  |  %8s   |\n"
	numberFormat = "%.1f"
)
type getRowFn func(row int) (string,string)//声明方法

//drawtable 绘画出两列的表格
func drawTable(hdr1,hdr2 string,rows int,getRow getRowFn){
	fmt.Println(line)
	fmt.Printf(rowFormat,hdr1,hdr2)
	fmt.Println(line)
	for row := 0;row<rows;row++{
		cell,cell2 := getRow(row)
		fmt.Printf(rowFormat,cell,cell2)
	}
	fmt.Println(line)
}
func ctof(row int) (string,string)  {
	c := celsius(row*5 -40)
	f:=c.fahrenheit()
	//fmt.Sprintf(格式化样式, 参数列表…)
	//格式化样式：字符串形式，格式化符号以 % 开头， %s 字符串格式，%d 十进制的整数格式。
	//参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。
	cell := fmt.Sprintf(numberFormat,c)
	cell2 := fmt.Sprintf(numberFormat,f)
	return cell,cell2
}
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	//fmt.Sprintf(格式化样式, 参数列表…)
	//格式化样式：字符串形式，格式化符号以 % 开头， %s 字符串格式，%d 十进制的整数格式。
	//参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}
func main() {
	drawTable("ºC", "ºF", 29, ctof)
	fmt.Println()
	drawTable("ºF", "ºC", 29, ftoc)
}