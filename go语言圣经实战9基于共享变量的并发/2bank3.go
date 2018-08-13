package bank

import "sync"

var (
	mu sync.Mutex
	balance int	//惯例来说,被mutex保护的变量是在mutex变量声明之后立刻声明的
)

//互斥锁
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()	//defer乃神器也
	return balance
}
