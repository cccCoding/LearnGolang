package main

import "fmt"

func main() {
	var names = []string {"啦啦","嘻嘻","呵呵","哈哈"}
	fmt.Println(index(names,"呵呵"))
	fmt.Println(include(names,"芭芭拉小魔仙"))
	nums := []int {1,2,3,4,5,6,7,8,9,10,11,12}
	fmt.Println(allData(nums, func(num int) bool {
		return num%3 == 0
	}))

	var a = 10
	var b = 20
	fmt.Println(getab(false)(a,b))
}

//查找索引的方法
func index(names []string, name string) int {
	for i, getname := range names {
		if name == getname {
			return i
		}
	}
	return -1
}

//判断是否存在
func include(names []string, name string) bool {
	return index(names,name) >= 0
}

//查找数组中所有能整除3的数的个数
// 函数作为参数
func allData(nums []int, check func(num int) bool) int {
	count := 0
	for _, num := range nums {
		if check(num) {
			count++
		}
	}
	return count
}

//函数作为返回值
func getab(flag bool) func(a int, b int) int {
	if flag {
		return func(a int, b int) int {
			return a + b
		}
	} else {
		return func(a int, b int) int {
			return a * b
		}
	}

}
