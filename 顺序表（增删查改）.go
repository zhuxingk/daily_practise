package main

import (
	"debug/dwarf"
	"errors"
	"fmt"
	"go/ast"
)

const MAXSIZE = 6 //存储空间初始分配容量
type ElementType int //*ElementType类型根据实际情况确定，此处定为*int

var data [MAXSIZE]ElementType //数组存储数据元素，最大值为MAXSIZE
var length int //线性表当前长度

//初始化操作
func InitList(){
	data = [MAXSIZE]ElementType{}//数组存储数据元素，最大值为MAXSIZE
	length =0
}

//线性表存储结构的插入和删除
/**
获取第一个元素
GetElem(L,i,*e);将线性表中的第I个元素值返回给e
 */
func GetElem(i int) (bool ,ElementType){
	if i<0||i>len(data){
		return false,-1
	}
	return true,data[i-1]
}
/**
查询指定元素所在位置索引
 */

func LocateElem(e ElementType) int{
	for i,v := range data{
		if v == e {
			return i
		}
	}
	return -1
}
/**
插入元素
 */
func ListInsert(e ElementType) (bool,error){
	if length == len(data){//顺序线性表已满，返回FALSE表示插入失败
		return false,errors.New(fmt.Sprintf("存储空间已满，无法插入元素：v%",e))
	}
	data[length] = e
	length++
	return true,nil
}
/**
如果插入位置不合理，排除异常
如果线性表长度大于等于数组长度，则排除异常或动态增加容量
将指定位置的数据向后移动，并将新的数据插入当前位置
长度加一
 */
func ListInsertByPosition(i int,e ElementType) (bool,error)  {
	if length == len(data){
		return false,errors.New("线性表空间已经蛮，无法继续插入新的数据元素")
	}
	if i < 0||i>len(data){
		return false,errors.New(fmt.Sprintf("索引越界"))
	}
	for k := length -1;k >=i;k--{
		data[i-1] =data[k]
	}
	data[i-1] =e
	length++
	return true,nil
}
/**
删除指定元素
 */
func ListDelete(e ElementType) bool{
	//获取改元素的索引位置
	var index int = LocateElem(2)
	if index >=0{
		for k := index;k < length;k++{
			if k == length -1{
				data[k] =0
			}
		}
		length--
		return true
	}
	return false
}
/**
删除指定位置的元素
 */
