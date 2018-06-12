package main

import "fmt"

func main2() {
	data := []int{1,2,3,4,5}
	fmt.Println(data)
	fmt.Println(data[0:3])	//包左不包右
	fmt.Println(data[:2])
	fmt.Println(data[3:])
	fmt.Println(data[:])
	fmt.Println(len(data))	//长度,数组当前的长度
	fmt.Println(cap(data))	//最大切片长度,就是数组可以扩展的最大的长度
}

func main() {
	var myslice[] int = make([]int ,4, 5)	//make函数创建一个用于切片的数组,len为4,cap为5
	fmt.Println(myslice)
	showSlice(myslice)

	var data[] int		//空切片
	fmt.Println(data)
	if data == nil {
		fmt.Println("空切片")
	}
	data = append(data, 1)	//追加
	data = append(data, 2)
	data = append(data, 3,4,5)
	fmt.Println(data)
	var num[] int = make([] int, 5, 6)	//要先开辟空间才能拷贝数据
	fmt.Println(num)
	copy(num,data)
	fmt.Println(num)
}

func showSlice(x[] int) {
	fmt.Println(len(x), cap(x), x)
}

func main1() {
	//锯齿数组
	data := make([][]int, 3)
	fmt.Println(data)
	for i := 0; i < 3; i++ {
		length := i + 1
		data[i] = make([]int ,length,length)
		for j := 0; j < length; j++ {
			data[i][j] = i + j
		}
	}
	fmt.Println(data)
}
