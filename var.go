package main

import (

	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main()  {
	const size = 300//图片大小
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//根据给定大小创建灰度图
	//遍历像素
	for x := 0;x < size; x++{
		for y := 0;y < size; y++ {
			pic.SetGray(x, y,color.Gray{255})
		}
	}
	//从0到最大像素生成x坐标
	for x := 0;x < size; x++ {
		//让sin的值范围在0到2pi之间
		s := float64(x) * 2 * math.Pi / size
		//让sin的幅度为一半的像素，向下偏移一半像素并反转
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y),color.Gray{0})
	}
	//创建文件
	file, err := os.Create("sin.png")
	if err != nil{
		log.Fatal(err)
	}
	png.Encode(file,pic)//将image信息写入文件中
	file.Close()
}
