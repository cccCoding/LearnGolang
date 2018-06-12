package main

import "fmt"

func main2() {
	var ids = [5]int{1,2,3,4,5}	//定义一个数组有5个int元素
	var names = [5]string{"a","b","c","d","e"}
	fmt.Println(ids,names)
	fmt.Println(names[1])	//访问数组元素

	for i, id := range ids {
		fmt.Println(i,id)
	}

	fmt.Println("-----------")

	for i := 0; i < len(names); i++ {
		fmt.Println(i,names[i])
	}

	fmt.Println("-----------")

	var data[10] int	//定义数组有10个int元素
	fmt.Println(data,len(data))
	for i := 0; i < len(data); i++ {	//数组元素初始化
		data[i] = i * 10
	}
	for i := 0; i < len(data); i++ {	//数组元素初始化
		fmt.Println(data[i])
	}
}

func main3() {
	var matrix[3][4] int		//2维数组	3行4列
	fmt.Println(matrix)

	var matrixplus[3][2][2] int		//3维数组
	fmt.Println(matrixplus)

	var matrix2 = [3][4] int {{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	matrix3 := [3][2][2] int {{{2,2},{3,4}},{{5,6},{7,8}},{{9,10},{11,12}}}
	fmt.Println(matrix2)
	fmt.Println(matrix3)
}

func main() {
	//数组作为参数,有副本机制
	var names = [5]string{"a","b","c","d","e"}	//不是同一个数组
	showArr(names)
	fmt.Println("main",names)

	//切片,没有副本机制,是同一个地址
	var ids = []int{1,2,3,4,5}
	showArr2(ids)
	fmt.Println("main",ids)
}

func showArr(para [5] string) {
	para[0] = "你好"
	fmt.Println("showArr",para)
}

func showArr2(para [] int) {
	para[0] = 100
	fmt.Println("showArr2",para)
}

