package main

import (
	"fmt"
	"strings"
)

var a = fmt.Println		//函数别名

func main() {
	a("哈哈哈")
	a(strings.Contains("大傻子","傻子"))	//是否包含
	a(strings.Count("hahahah","a"))	//出现次数
	a(strings.HasPrefix("nihao","ni"))//是否以某个字符串开头
	a(strings.HasSuffix("nihao","ao"))//是否以某个字符串结尾
	a(strings.Index("nihao","ih"))//索引
}
