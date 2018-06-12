package main

import "fmt"

//拓展业务接口,一个接口可以对应多个类型,拓展软件开发
type boytype interface {
	girlgo()
}

type 冷酷男 struct {
}

type 暖男 struct {
}

func (朱建涛 暖男)girlgo() {	//重写方法	//前面那个括号表示要重写暖男这个类中的girlgo方法
	fmt.Println("我是暖男啦啦啦")
}

func (莽子 冷酷男)girlgo() {
	fmt.Println("我是冷酷男呵呵呵")
}

func main() {
	var boy boytype		//变量类型定义为父类,或者说是接口

	boy = new(冷酷男)	//对象可以是任意子类
	boy.girlgo()

	boy = new(暖男)
	boy.girlgo()
}
