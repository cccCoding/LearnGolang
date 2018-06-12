package main

import "fmt"

func main() {
	//goToSwim()
	fmt.Println(fac(4))

	//腾讯面试题:台阶问题,一次一级或者二级台阶,走上50级台阶,有多少种可能
	fmt.Println(tencent(50))
}

func goToSwim() {		//死循环
	fmt.Println("goToSwim")
	goToSwim()
}

func fac(n uint64) (result uint64) {	//从1乘到n
	if n > 0 {
		result = n * fac(n - 1)
		return result
	} else {
		return 1
	}
}

func tencent(n int) int {
	if n <= 2 {
		return n
	} else {
		return tencent(n - 1) + tencent(n - 2)
	}
}