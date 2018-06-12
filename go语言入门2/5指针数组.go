package main

import "fmt"

func main() {

	var arr = []int {7,8,9}
	var parr [3]*int		//指针数组
	var pparr [3]**int		//指针的指针
	for i := 0; i < len(arr); i++ {
		parr[i] = &arr[i]
		pparr[i] = &parr[i]
		fmt.Println(i, arr[i], parr[i], pparr[i], *parr[i], **pparr[i])
	}
	
	var num = 100
	fmt.Println("main",num)
	mvnum(num)
	fmt.Println("main",num)
	mvpnum(&num)
	fmt.Println("main",num)
}

func mvnum(num int)  {
	num = 1000
	fmt.Println("mvnum",num)
}

func mvpnum(pnum *int)  {
	*pnum = 1000
	fmt.Println("mvnum",*pnum)
}
