package main

// 这些常见坑：符合Go语言语法的，可以正常的编译，但是可能是运行结果错误，或者是有资源泄漏的风险
// 来源：https://chai2010.cn/advanced-go-programming-book/appendix/appendix-a-trap.html
// 个人感觉这里面的坑其实大部分都没有那么容易遇到，不如这篇文章的坑，其实更容易遇到 https://juejin.cn/post/7241452578125824061?utm_source=gold_browser_extension#comment

/**
当参数的可变参数是空接口类型时，传入空接口的切片时需要注意参数展开的问题。
不管是否展开，编译器都无法发现错误，但是输出是不同的：
[1 2 3]
1 2 3
**/
// func main() {
// 	var a = []interface{}{1, 2, 3}
// 	fmt.Println(a)
// 	fmt.Println(a...)
// }

/**
数组是值传递
在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果
[7 2 3]
[1 2 3]
**/
// func main() {
// 	x := [3]int{1, 2, 3}
// 	func(arr [3]int) {
// 		arr[0] = 7
// 		fmt.Println(arr)
// 	}(x)
// 	fmt.Println(x)
// }

/**
map是一种hash表实现，每次遍历的顺序都可能不一样
**/
// func main() {
// 	m := map[string]string{
// 		"1": "1",
// 		"2": "2",
// 		"3": "3",
// 	}

// 	for k, v := range m {
// 		println(k, v)
// 	}
// }

/**
返回值被屏蔽--在局部作用域中，命名的返回值内同名的局部变量屏蔽
**/
// func Foo() (err error) {
// 	if err := Bar(); err != nil {
// 		return
// 	}
// 	return
// }

/**
recover必须在defer函数中运行
**/
// func main() {
// 	// // 成功的：
// 	// defer func() {
// 	// 	recover()
// 	// }()
// 	// panic(1)
// 	// 下面的 recover 都是失败的不成功
// 	// // 1. recover捕获的是祖父级调用时的异常，直接调用时无效
// 	// recover()
// 	// panic(1)
// 	// // 2. 直接defer调用也是无效
// 	// defer recover()
// 	// panic(1)
// 	// // 3. defer调用时多层嵌套依然无效
// 	// defer func() {
// 	// 	func() { recover() }()
// 	// }()
// 	// panic(1)
// }

/**
main函数提前退出--后台Goroutine无法保证完成任务
**/
// func main() {
// 	go println("hello")
// }

/**
通过Sleep来回避并发中的问题--休眠/调度语句并不能保证输出完整的字符串
**/
// func main() {
// 	go println("hello")
// 	// time.Sleep(time.Second)
// 	// runtime.Gosched()
// }

/**
独占CPU导致其它Goroutine饿死----Goroutine 是协作式抢占调度（Go1.14版本之前），Goroutine本身不会主动放弃CPU
**/
// func main() {
// 	runtime.GOMAXPROCS(1)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(i)
// 		}
// 	}()

// 	for {
// 		// runtime.Gosched() // 加入这个语句可以解决 goroutine 饿死的情况
// 	} // 占用CPU
// }

/**
不同Goroutine之间不满足顺序一致性内存模型
因为在不同的Goroutine，main函数中无法保证能打印出hello, world
很可能 done 已经成功赋值 true，而 msg 还未赋值，我们需要显性保证同步
	var msg string
	var done bool
	func setup() {
		msg = "hello, world"
		done = true
	}
	func main() {
		go setup()
		for !done {
		}
		println(msg)
	}

可以通过 chan 保证同步，也可以通过 sync.Lock 保证
**/
// var msg string
// var done = make(chan bool)

// func setup() {
// 	msg = "hello, world"
// 	done <- true
// }

// func main() {
// 	go setup()
// 	<-done
// 	println(msg)
// }

/**
	闭包错误引用同一个变量
	// func main() {
	// 	for i := 0; i < 5; i++ {
	// 		defer func() {
	// 			println(i)
	// 		}()
	// 	}
	// }
	只会输出 5 个 5，而不会以此输出 4~0

改进的方法： 在每轮迭代中生成一个局部变量 或 通过函数参数传入
**/
// func main() {
// 	for i := 0; i < 5; i++ {
// 		i := i
// 		defer func() {
// 			println(i)
// 		}()
// 	}
// }
// func main() {
// 	for i := 0; i < 5; i++ {
// 		defer func(i int) {
// 			println(i)
// 		}(i)
// 	}
// }

/**
在循环内部执行defer语句---defer在函数退出时才能执行，在for执行defer会导致资源延迟释放
func main() {
	for i := 0; i < 5; i++ {
		f, err := os.Open("temp")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}
// 以上不会报错，但不够好，正常应该是 for 循环中每次 open 操作完后都 close f，而不是最后重复执行 f.close
**/
// 解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer
// func main() {
// 	for i := 0; i < 5; i++ {
// 		func() {
// 			f, err := os.Open("temp")
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			defer f.Close()
// 		}()
// 	}
// }

/**
切片会导致整个底层数组被锁定---切片会导致整个底层数组被锁定，底层数组无法释放内存。如果底层数组较大会对内存产生很大的压力
// func main() {
// 	headerMap := make(map[string][]byte)
// 	for i := 0; i < 5; i++ {
// 		name := "temp"
// 		data, err := ioutil.ReadFile(name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		headerMap[name] = data[:1]
// 	}
// 	// do some thing
// }
**/
// 下面的代码主要改进在---并不直接 data[:1]赋值，而是通过 append 操作拷贝了 data[:1]
// append 操作追加的元素是值传递而非引用传递
// func main() {
// 	headerMap := make(map[string][]byte)
// 	for i := 0; i < 5; i++ {
// 		name := "temp"
// 		data, err := ioutil.ReadFile(name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		headerMap[name] = append([]byte{}, data[:1]...)
// 	}
// 	// do some thing
// }

/**
空指针和空接口不等价--比如返回了一个错误指针，但是并不是空的error接口
空指针可以被视为一个特殊的接口值，但并不完全等价
**/
// type MyStruct struct {
// 	Name string
// }
// func main() {
// 	var p *MyStruct = nil      // 空指针 p，指向 MyStruct 类型
// 	var i interface{} = p      // 将 p 赋值给一个接口变量 i
// 	fmt.Println(p == nil)      // 输出: true 空指针正常等于空指针
// 	fmt.Println(i == nil)      // 输出: false 空接口不等于空指针，接口值是非空的其实，空接口值由两部分组成：动态类型和动态值。在这种情况下，动态类型是 *MyStruct，动态值是 nil
// 	fmt.Println(i == nilPtr()) // 输出: true 动态类型都为 *MyStruct，动态值都为 nil
// }
// func nilPtr() *MyStruct {
// 	return nil
// }

/**
内存地址会变化---Go语言中对象的地址可能发生变化，因此指针不能从其它非指针类型的值生成
**/
// func main() {
// 	var x int = 42
// 	var p uintptr = uintptr(unsafe.Pointer(&x))
// 	runtime.GC()
// 	var px *int = (*int)(unsafe.Pointer(p))
// 	println(*px)
// 	// 即 p 原先是指向 x 的地址，但 GC 后，p 还是那个地址，但地址对应的值不一定还是 x 了
// 	// 当内存发生变化的时候，相关的指针会同步更新，但是非指针类型的uintptr不会做同步更新。
// 	// 同理CGO中也不能保存Go对象地址
// }

/**
Goroutine泄露
Go语言是带内存自动回收的特性，因此内存一般不会泄漏
但是 Goroutine 却存在泄漏的情况，同时泄漏的 Goroutine 引用的内存同样无法被回收
**/
// func main() {
// 	ch := func() <-chan int {
// 		ch := make(chan int)
// 		go func() {
// 			for i := 0; ; i++ {
// 				ch <- i
// 			}
// 		}()
// 		return ch
// 	}()
//		for v := range ch {
//			fmt.Println(v)
//			if v == 5 {
//				break
//			}
//		}
//	}
//
// 上面的程序中后台Goroutine向管道输入自然数序列，main函数中输出序列
// 但是当break跳出for循环的时候，后台Goroutine就处于无法被回收的状态了，因为 ch <- 7 被阻塞住了，ch 中还有一个 6 等待被消费
// 可以通过context+select来避免这个问题
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	ch := func(ctx context.Context) <-chan int {
// 		ch := make(chan int)
// 		go func() {
// 			for i := 0; ; i++ {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case ch <- i:
// 				}
// 			}
// 		}()
// 		return ch
// 	}(ctx)
// 	for v := range ch {
// 		fmt.Println(v)
// 		if v == 5 {
// 			cancel()
// 			break
// 		}
// 	}
// }

/**
容器中的 GOMAXPROCS
自 Go 1.5 开始， Go 的 GOMAXPROCS 默认值已经设置为 CPU 的核数，但是在 Docker 或 k8s 容器中 runtime.GOMAXPROCS() 获取的是 宿主机的 CPU 核数 。
这样会导致 P 值设置过大，导致生成线程过多，会增加上下文切换的负担，导致严重的上下文切换，浪费 CPU
所以可以使用 uber 的 automaxprocs 库，大致原理是读取 CGroup 值识别容器的 CPU quota，计算得到实际核心数，并自动设置 GOMAXPROCS 线程数量
**/

// // slice 扩容地址发生改变
// func main() {
// 	s := []int{1, 2}
// 	for i := 0; i < 16; i++ {
// 		s = append(s, 3, 4, 5)
// 		fmt.Printf("add:%p, len:%d, cap:%d\n", s, len(s), cap(s))
// 	}
// }

// // golang 的 mutex 不可重入
// func main() {
// 	m := sync.Mutex{}
// 	m.Lock()
// 	m.Lock()
// 	m.Lock()
// 	m.Unlock()
// 	m.Unlock()
// 	m.Unlock()
// }

// 生成固定数量的协程，接收某一 chan 的消息处理
// func main() {
// 	c := make(chan bool)
// 	// 固定死只有 10 个处理协程
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			for v := range c {
// 				fmt.Println(v)
// 			}
// 		}()
// 	}
// }

// 基于带缓冲的 chan，缓冲满时阻塞的原理->实现缓冲长度的并发协程数量
// func main() {
// 	c := make(chan bool, 3)
// 	// 有 1000 个请求，耦合生产和消费，管道长度即最大并发数
// 	for i := 0; i < 1000; i++ {
// 		c <- true
// 		i := i
// 		go func() {
// 			fmt.Println(i)
// 			<-c
// 		}()
// 	}
// }
