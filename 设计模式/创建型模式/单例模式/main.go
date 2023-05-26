package main

import (
	"fmt"
	"sync"
)

/**
单例模式介绍：
单例模式 --- 创建型模式
思想：保证一个类只有一个实例， 并提供一个访问该实例的全局节点（同时解决两个问题，违反单一职责原则）

适用场景：
1. 当一个类的实例在系统中只需要存在一个，以节省系统资源或确保一致性时，可以使用单例模式。例如，日志记录器、数据库连接池、线程池等。
2. 当需要限制某个类只能拥有一个实例，以便控制特定资源或实现特定功能时，可以使用单例模式。例如，配置信息管理类、全局计数器等。
3. 当一个类的对象需要被多个对象共享访问时，可以使用单例模式。例如，缓存管理器、身份验证管理器等。

具体应用场景：
1. 日志记录器：在一个应用程序中，通常只需要一个日志记录器实例来记录日志信息。通过使用单例模式，可以确保在整个应用程序中只有一个日志记录器实例存在，方便记录和管理日志。
2. 数据库连接池：在多线程环境下，数据库连接的创建和销毁是一项昂贵的操作。通过使用单例模式来管理数据库连接池，可以确保在整个应用程序中只有一个连接池实例存在，提供高效的数据库连接复用。
3. 配置信息管理器：应用程序中的配置信息通常是全局共享的，例如数据库连接参数、系统设置等。通过使用单例模式来管理配置信息，可以方便地在应用程序的各个部分访问和修改配置信息。
4. UI界面中的全局状态管理：在一些应用程序中，可能需要共享一些全局状态，例如用户登录状态、界面主题等。通过使用单例模式来管理这些全局状态，可以确保在整个应用程序中只有一个状态实例存在，并方便地在不同的界面模块之间进行状态的读取和修改。

单例模式的优缺点
优点：
-- 保证一个类只有一个实例并获得了一个指向该实例的全局访问节点
-- 节省系统资源，仅在首次请求单例对象时对其进行初始化，特别是在需要频繁创建和销毁对象的情况下
-- 实现线程安全，单例模式提供线程安全的访问，避免多个线程同时创建多个实例或访问不同的实例问题
-- 避免全局变量的滥用，使用单例模式可以避免过多使用全局变量，避免全局变量带来的命名冲突和作用域混乱的问题
缺点：
-- 违反了单一职责原则。 该模式同时解决了两个问题。
-- 难以拓展，由于单例模式只允许存在一个实例，因此扩展功能需要直接修改其代码，违反了开闭原则
-- 对象生命周期长，单例模式创建的实例往往在应用的整个生命周期都存在，可能导致对象持有的资源无法及时释放，增加内存占用
-- 高耦合性，单例对象的访问点是全局的，其他对象需要直接依赖或引用该单例对象，也容易被滥用，增加代码复杂性和维护难度
**/

// // 单例双重校验---常规Lock写法
// type single struct{}

// var (
// 	lock           sync.Mutex // 用于加锁保证多线程/多协程同时创建时不会出错
// 	singleInstance *single    // 全局访问单例，只能通过 GetInstance 获取
// )

// func GetInstance() *single {
// 	if singleInstance != nil {
// 		fmt.Println("Single instance already exist")
// 		return singleInstance
// 	}
// 	lock.Lock()
// 	defer lock.Unlock()
// 	if singleInstance == nil {
// 		singleInstance = &single{}
// 		fmt.Println("Creating single instance now")
// 	} else {
// 		fmt.Println("LOCKING!! Single instance already exist")
// 	}
// 	return singleInstance
// }

// func main() {
// 	for i := 0; i < 30; i++ {
// 		go GetInstance()
// 	}
// 	// Scanln is similar to Scan, but stops scanning at a newline and
// 	// after the final item there must be a newline or EOF.
// 	fmt.Scanln()
// }

// 单例双重校验---常规Once写法
type single struct{}

var (
	once           sync.Once
	singleInstance *single
)

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
