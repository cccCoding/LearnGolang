package main

import "sync"

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

type entry struct {
	res result
	ready chan struct{}		//close when res is ready
}

type Func func(key string) (interface{}, error)	//要缓存的函数类型

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//第一次对这个key进行请求，计算值并广播ready事件
		e = &entry{ready:make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)	//关闭channel以广播ready
	} else {
		memo.mu.Unlock()

		<-e.ready		//等待ready事件，如果未关闭会阻塞，直到关闭
	}
	return e.res.value, e.res.err
}
