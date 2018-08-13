package bank

var (
	sema = make(chan struct{}, 1)
	balance int
)

//用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量
func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
