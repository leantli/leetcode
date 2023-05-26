package main

import "fmt"

/**
代理模式介绍：
代理模式 --- 结构型模式
思想：能够提供对象的替代品或其占位符给客户端使用
代理控制着对于原对象的访问，并允许在将请求提交给对象前后进行一些处理

适用场景：
1. 延迟初始化(虚拟代理)：如果你有一个偶尔使用的重量级服务对象， 一直保持该对象运行会消耗系统资源时， 可使用代理模式
2. 访问控制(安全/保护代理)：代理对象可以控制对实际对象的访问权限，例如，在某些情况下，只允许特定的用户或角色访问实际对象，代理对象可以根据权限进行访问控制
3. 日志记录代理：代理对象可以在调用实际对象的方法前后进行日志记录，用于跟踪和分析系统的行为和性能
4. 缓存请求结果(缓存代理)：代理对象可以在访问实际对象前检查缓存，如果缓存中已经存在所需的结果，直接返回缓存结果，减少对实际对象的访问。这样可以提高系统的性能和响应速度
5. 智能引用：可在没有客户端使用某个重量级对象时立即销毁该对象
6. 本地执行远程服务(远程代理)：适用于服务对象位于远程服务器上的情形

代理模式的优缺点
优点：
-- 你可以在客户端毫无察觉的情况下控制服务对象
-- 如果客户端对服务对象的生命周期没有特殊要求， 你可以对生命周期进行管理
-- 即使服务对象还未准备好或不存在， 代理也可以正常工作
-- 开闭原则。 你可以在不对服务或客户端做出修改的情况下创建新代理
缺点：
-- 增加了系统复杂性：引入代理对象会增加代码的复杂性，需要额外的类和接口来管理代理和实际对象之间的关系
-- 增加了运行时开销：在访问实际对象之前或之后，代理对象可能需要执行一些额外的逻辑，导致运行时开销的增加
**/

// // 例一：利用代理模式记录日志
// // 定义实际对象和代理对象的共同接口，接着代理对象实现对应的接口，并额外实现before()和after()函数
// // Subject 定义了实际对象和代理对象的共同接口
// type Subject interface {
// 	DoSomething()
// }

// // RealSubject 实际对象，执行具体的操作
// type RealSubject struct{}

// func (rs *RealSubject) DoSomething() {
// 	fmt.Println("RealSubject: DoSomething")
// }

// // Proxy 代理对象，在执行操作前后记录日志
// type Proxy struct {
// 	realSubject *RealSubject
// }

// func (p *Proxy) DoSomething() {
// 	p.before()
// 	p.realSubject.DoSomething()
// 	p.after()
// }

// func (p *Proxy) before() {
// 	fmt.Println("Proxy: Before")
// }

// func (p *Proxy) after() {
// 	fmt.Println("Proxy: After")
// }

// func main() {
// 	// 创建实际对象
// 	realSubject := &RealSubject{}
// 	// 创建代理对象
// 	proxy := &Proxy{
// 		realSubject: realSubject,
// 	}
// 	// 使用代理对象调用方法
// 	proxy.DoSomething()
// }

// 例二：Nginx 提供对应用服务的受控访问权限控制，且可缓存请求
// 1. 代理和被代理类的共有接口
type server interface {
	handleReq(url, method string) (code int, resp string)
}

// 2. 被代理类的接口实现
type Application struct{}

func (a *Application) handleReq(url, method string) (code int, resp string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

// 3. 代理类的接口实现
type Nginx struct {
	application *Application
	maxAllowReq int
	rateLimiter map[string]int
}

func newNginx() *Nginx {
	return &Nginx{
		application: &Application{},
		maxAllowReq: 2,
		rateLimiter: map[string]int{},
	}
}

func (n *Nginx) handleReq(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleReq(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowReq {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func main() {
	nginxServer := newNginx()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"
	httpCode, body := nginxServer.handleReq(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleReq(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleReq(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleReq(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleReq(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
