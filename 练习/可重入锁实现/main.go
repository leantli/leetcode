package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"net/http"
	_ "net/http/pprof"
)

// golang 中没有可重入锁，那么如何实现？
// 1. 记住持有锁的线程
// 2. 统计重入次数

// // 没有可重入锁---会直接报错
// func main() {
// 	var lock sync.Mutex
// 	var a int
// 	lock.Lock()
// 	a = 1
// 	print(a)
// 	lock.Lock()
// 	a = 2
// 	print(a)
// 	lock.Unlock()
// 	a = 3
// 	print(a)
// 	lock.Unlock()
// }

// 获取协程 id
func GetGoroutineID() int64 {
	var buf [64]byte
	idx := runtime.Stack(buf[:], false)
	// fmt.Println(string(buf[:idx]))
	idField := strings.Fields(string(buf[:idx]))[1] // buf 的string 输出为 goroutine [id] [status:running]...
	// fmt.Println(idField)
	id, _ := strconv.Atoi(idField)
	return int64(id)
}

type ReentrantLock struct {
	sync.Mutex
	owner     int64 // 当前持有锁的 goroutine id
	recursion int32 // 重入次数
}

func (l *ReentrantLock) Lock() {
	gid := GetGoroutineID()
	if atomic.LoadInt64(&l.owner) == gid {
		atomic.AddInt32(&l.recursion, 1)
		return
	}
	l.Mutex.Lock()
	atomic.StoreInt64(&l.owner, gid)
	atomic.StoreInt32(&l.recursion, 1)
}

func (l *ReentrantLock) UnLock() {
	gid := GetGoroutineID()
	// 非持有锁的 goroutine 释放锁会错误
	if gid != atomic.LoadInt64(&l.owner) {
		panic(fmt.Sprintf("wrong owner[%d]:%d\n", l.owner, gid))
	}
	atomic.AddInt32(&l.recursion, -1)
	if atomic.LoadInt32(&l.recursion) == 0 {
		// goroutine 最后一次调用该锁，释放锁
		atomic.StoreInt64(&l.owner, -1)
		l.Mutex.Unlock()
	}
}

func main() {
	l := &ReentrantLock{}
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			l.Lock()
			l.Lock()
			fmt.Println(i)
			l.UnLock()
			l.UnLock()
			l.Lock()
			l.UnLock()
		}()
	}
	for i := 0; i < 3; i++ {
		m := sync.Mutex{}
		go func() {
			defer func() {
				err := recover()
				fmt.Println(err)
				return
			}()
			m.Lock()
			m.Lock()
			fmt.Println("222")
			m.Unlock()
			m.Unlock()
			m.Lock()
			m.Unlock()
		}()
	}
	go func() {
		var c chan bool
		c <- true
	}()
	go func() {
		http.ListenAndServe("localhost:4040", nil)
	}()
	var c chan int
	c <- 1
}
