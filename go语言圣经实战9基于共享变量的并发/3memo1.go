package main

// 缓存函数，
// 需要缓存函数的返回结果，这样在对函数进行调用的时候，就只需要进行一次计算，
// 之后只要返回计算的结果就可以
// 并发安全且会避免对整个缓存加锁而导致所有操作都去争一个锁的设计

type Memo struct {
	f Func					//函数
	cache map[string]result	//参数和返回结果
}

type Func func(key string) (interface{}, error)	//要缓存的函数类型

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]	//先尝试从map中获取值
	if !ok {					//如果没有值，说明这个方法用这个参数没有执行过，执行然后保存值到map
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
