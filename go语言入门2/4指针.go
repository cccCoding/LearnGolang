package main

import "fmt"

func main() {
	var num = 10
	var pnum *int = &num
	fmt.Println(num, &num, pnum)	//数据及其地址
	*pnum = 11						//通过指针修改数据
	fmt.Println(num, &num, pnum)

	var count = 10
	var pcount *int			//空指针	<nil>
	pcount = &count
	fmt.Println(pcount)
	if pcount == nil {
		fmt.Println("空指针")
	} else {
		fmt.Println("不是空指针")
	}
}
