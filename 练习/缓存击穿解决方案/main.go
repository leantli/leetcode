package main

import (
	"fmt"
	"sync"
)

// 缓存击穿：高并发查询某数据，无缓存(缓存过期、缓存被删除)直接打在数据库上，对数据库造成巨大压力，严重情况导致数据库宕机
// 解决方案：
// 1. 热点数据永不过期(但是总有需要更新缓存的时候，因此我们需要考虑下面的方案2)
// 2. 加(读写)互斥锁，缓存数据过期后，即便同时有多个请求进来，也只有一个线程去查数据库

// // 第一个热身
// // 测试 golang 读写锁的优先级---写锁进来时，会等前面的读锁都解开才会成功上写锁；而写锁后面想要加读锁的会阻塞住（即验证不会写饥饿）
// // 总共12秒的读锁，在第三秒时会上一个写锁1，等待前面三个读锁解锁，3+4=7，在第7秒的时候成功上写锁1，再等待3秒，在 10 秒时写锁1解开；
// // 此时4-9秒的读锁都在排队中，写锁1一解开就全部上读锁；写锁2虽然在第5秒就开始尝试取写锁，但是要等上一个写锁释放才能成功排队；
// // 因此在第10秒开始写锁2正式抢到，后面的10-12的读锁都得等排队
// // 顺序 (1 2 3 读锁) (写锁1) (4 5 6 7 8 9 读锁) (写锁2) (10 11 12 读锁)
// func main() {
// 	var cacheLock sync.RWMutex
// 	var wg sync.WaitGroup
// 	wg.Add(12)
// 	// 等 3 秒再上写锁1
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(time.Second * 3)
// 		fmt.Println("尝试上写锁1")
// 		cacheLock.Lock()
// 		fmt.Println("成功上写锁1")
// 		time.Sleep(3 * time.Second)
// 		cacheLock.Unlock()
// 		fmt.Println("解除写锁1")
// 	}()

// 	// 等5秒再上写锁2
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(time.Second * 5)
// 		fmt.Println("尝试上写锁2")
// 		cacheLock.Lock()
// 		fmt.Println("成功上写锁2")
// 		time.Sleep(3 * time.Second)
// 		cacheLock.Unlock()
// 		fmt.Println("解除写锁2")
// 	}()

// 	// 前 12 秒，每秒上一个读锁
// 	for i := 1; i <= 12; i++ {
// 		go func() {
// 			defer wg.Done()
// 			temp := i
// 			cacheLock.RLock()
// 			fmt.Printf("%d已经锁上读锁\n", temp)
// 			time.Sleep(time.Second * 4)
// 			cacheLock.RUnlock()
// 			fmt.Printf("%d已经解除读锁\n", temp)
// 		}()
// 		time.Sleep(time.Second)
// 	}
// 	wg.Wait()
// 	fmt.Println("全部结束")
// }

// // 例子二：加读写锁，缓存数据不存在(过期、删除)后，即便同时有多个请求进来，也只有一个线程去查数据库
// var (
// 	cache        map[string]string
// 	cacheLock    sync.RWMutex
// 	databaseData map[string]string
// )

// func main() {
// 	// 模拟初始化缓存和数据库数据
// 	cache = make(map[string]string)
// 	cache["key1"] = "value1"
// 	databaseData = make(map[string]string)
// 	databaseData["key1"] = "value1"

// 	// 启动多个并发请求进行读取
// 	for i := 1; i <= 10; i++ {
// 		go readData(fmt.Sprintf("key%d", i))
// 	}
// 	time.Sleep(time.Second * 5)
// 	fmt.Println("----第二轮请求进来------")
// 	// 启动多个并发请求进行读取
// 	for i := 1; i <= 10; i++ {
// 		go readData(fmt.Sprintf("key%d", i))
// 	}

// 	// 等待请求处理完成
// 	fmt.Scanln()
// }

// func readData(key string) {
// 	// 获取读锁，然后尝试从缓存中读取数据
// 	fmt.Printf("尝试获取--%s--的缓存数据\n", key)
// 	cacheLock.RLock()
// 	fmt.Printf("--%s--获取到读锁\n", key)
// 	value, ok := cache[key]
// 	cacheLock.RUnlock()

// 	if ok {
// 		fmt.Printf("缓存命中，直接返回数据: %s - %s\n", key, value)
// 		return
// 	}

// 	// 缓存未命中，获取写锁
// 	cacheLock.Lock()
// 	fmt.Printf("--%s--未命中，获取到写锁\n", key)

// 	// 注意!!获取到写锁后并不是直接就读数据库并且重载缓存
// 	// 而是先看看缓存是否已经重载过了，避免同时多请求发现没缓存后到抢写锁这一步
// 	if value, ok := cache[key]; ok {
// 		cacheLock.Unlock()
// 		fmt.Printf("缓存数据重载成功: %s - %s\n", key, value)
// 		return
// 	}

// 	// 从数据库中查询数据 并 更新缓存
// 	valueFromDatabase := queryDatabase(key)
// 	cache[key] = valueFromDatabase
// 	// 解除写锁
// 	cacheLock.Unlock()
// 	fmt.Printf("缓存更新: %s - %s\n", key, valueFromDatabase)
// }

// func queryDatabase(key string) string {
// 	// 模拟从数据库中查询数据的逻辑
// 	time.Sleep(time.Second)
// 	return "ValueOf" + key
// }

// 例子三：黑名单缓存服务提供
var rwLock sync.RWMutex
var blackList []string

func getBlackList() ([]string, error) {
	rwLock.RLock()
	if blackList != nil {
		res := make([]string, len(blackList))
		copy(res, blackList)
		rwLock.RUnlock()
		return res, nil
	}
	rwLock.RUnlock()

	rwLock.Lock()
	defer rwLock.Unlock()
	if blackList != nil {
		res := make([]string, len(blackList))
		copy(res, blackList)
		return res, nil
	}
	// 模拟访问数据库获取 blackList
	blackList = []string{"test"}
	res := make([]string, len(blackList))
	copy(res, blackList)
	return res, nil
}

func updateBlackList() error {
	rwLock.Lock()
	defer rwLock.Unlock()
	// 模拟访问数据库获取blackList
	blackList = []string{"temp"}
	return nil
}

func main() {
	fmt.Println(getBlackList())
}
