package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getvalue(filename,expectSection,expectKey string) string {
	//读取文件
	file, err := os.Open(filename) //打开文件
	if err != nil {
		return " " //文件找不到，返回空
	}
	defer file.Close() //函数结束时，关闭文件

	//读取行文本
	//使用读取器读取文件
	reader := bufio.NewReader(file)
	//当前读取的段的名字
	var sectionName string
	for {
		//读取文件的一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//切掉行左右两边的空白字符
		linestr = strings.TrimSpace(linestr)
		//忽略空行
		if linestr == "" {
			continue
		}
		//忽略注释
		if linestr[0] == ';' {
			continue
		}
		//读取段
		//行首和尾巴分别时方括号的，说明是段标记的起止符
		if linestr[0] == '[' && linestr[len(linestr)-1] ==']'{
			//将段名取出
			sectionName = linestr[1 : len(linestr)-1]
			//这个段是希望读取的
		}else if sectionName == expectSection{//读取键值
			//切开等号分割的键值对
			pair := strings.Split(linestr,"=")
			//保证切开只有1个等号分割的键值情况
			if len(pair) == 2{
				//去掉多余的空白字符
				key := strings.TrimSpace(pair[0])
				//是期望的建
				if key == expectKey{
					//返回去掉空白字符的值
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return strings.TrimSpace(getvalue(filename,expectSection,expectKey))
}
func main(){
	fmt.Println(getvalue("example.ini","remote \"origin\"","fetch"))
	fmt.Println(getvalue("example.ini","core","hideDotFiles"))
}
