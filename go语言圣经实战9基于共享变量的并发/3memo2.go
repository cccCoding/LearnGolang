package main

import "sync"

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)	//要缓存的函数类型

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]result)}
}

//加锁解决了并行带来的问题，但是性能又下降了
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}