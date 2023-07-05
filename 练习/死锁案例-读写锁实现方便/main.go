package main

import (
	"fmt"
	"sync"
)

// // 实现一个死锁看看

type Temp struct {
	lock sync.RWMutex
	name string
}

func (t *Temp) String() string {
	var res string
	t.lock.RLock()
	defer t.lock.RUnlock()
	res = t.name
	return res
}

func (t *Temp) updateName() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.name = "nihao"
	fmt.Printf("t.Name:%s", t)
}

func main() {
	var temp Temp
	temp.updateName()
	fmt.Println("success")
}

// // 没有加锁对 map 进行并发操作----会直接报错 deadlock
// // 写操作中需要获取读锁
// func main() {
// 	m := make(map[int]int)
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			m[i] = i
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			fmt.Println(m[i])
// 		}
// 	}()
// 	select {}
// }
