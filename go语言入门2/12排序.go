package main

import (
	"sort"
	"fmt"
)

func main() {
	scores := []int {77,89,78,56,194}
	issorted := sort.IntsAreSorted(scores)	//判断是否排序
	fmt.Println(issorted)
	sort.Ints(scores)		//按大小排序
	issorted = sort.IntsAreSorted(scores)
	fmt.Println(issorted,scores)

	names := []string {"v","c","a","z"}
	sort.Strings(names)
	fmt.Println(names)

	//自定义排序 按字符串长度
	values := []string {"as","c","asadfgasdfd","zadsf"}
	sort.Sort(ByLength(values))
	fmt.Println(values)
}

type ByLength []string		//接口	//要重写下面几个方法

func (s ByLength)Len() int {	//长度	//重写方法
	return len(s)
}

func (s ByLength)Swap(i int, j int)  {	//交换
	s[i],s[j] = s[j],s[i]
}

func (s ByLength)Less(i, j int) bool {	//比较大小
	return len(s[i]) < len(s[j])
}
