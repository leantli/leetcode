package main

import "fmt"

/**
责任链模式介绍：
责任链模式(CoR、Chain of Responsibility) --- 行为设计模式
思想：允许请求沿着处理链进行发送，收到请求后，每个处理者都可以对请求进行处理，或将其传递给链上的下一个处理者

适用场景：
1. 多个对象可以处理同一种类型的请求，但具体由哪个对象处理请求是动态决定的。
2. 需要按照顺序逐个处理请求，并且每个处理者都可以决定是否继续传递请求。
3. 发送者不需要知道请求的具体处理者，只需将请求发送给链的起始点，由链中的处理者自行负责处理。
4. 希望动态地添加或移除处理者，并灵活组织处理流程
-- 如果在处理者类中有对引用成员变量的设定方法， 你将能动态地插入和移除处理者， 或者改变其顺序
-- 在例三的基础上
-- 先用一个map存储所有的handler，key为handler的name，然后热读配置中启用的name
-- 单独加入chain中的list，再正常运作即可

具体应用场景：
1. 日志记录器链：不同类型的日志记录器可以按照顺序记录日志，每个记录器可以决定是否继续传递给下一个记录器。
2. 数据验证链：不同的验证器可以逐个验证数据的合法性，每个验证器可以决定是否继续传递给下一个验证器。
3. 过滤器/拦截器链：（过滤器链）多个请求过滤器可以按照顺序对请求进行过滤，每个过滤器可以决定是否继续传递给下一个过滤器。
4. 购买流程链：购买流程中的不同环节可以作为处理者，按照顺序处理购买请求，例如验证库存、计算价格、生成订单等。

责任链模式的优缺点
优点：
-- 解耦处理逻辑，将请求发送者和接收者解耦，请求发送者无需知道具体的接收者和处理细节，只需统一调用接口，上层无需知道下层具体实现，这种松耦合有助于减少对象之间的依赖关系，符合迪米特法则和依赖倒置原则
-- 可重用的处理逻辑，每个处理者都专注于自己的处理逻辑，使得处理逻辑可以单独测试、维护和复用，可以减少重复代码，提高代码可读性和可维护性，同样符合单一职责原则
-- 灵活性和可拓展性，可以在不更改现有代码的情况下灵活地增加处理者，符合开闭原则
缺点：
-- 链中没有适当的处理者，请求可能被忽略或丢失(不觉得这个成立，如果每个处理者都不会处理这个请求，那就算不使用责任链模式，将所有处理都整合在一起，这个请求被处理也没有意义，说明这个请求根本就不该出现？)
-- 性能考虑：责任链模式需要遍历链知道找到能处理这个请求的处理者，链过长会影响性能(如果是个需要匹配的链，在一定程度上确实有影响)
-- 责任模糊：一个请求可能被多个处理器这处理，难以确定哪个处理者对请求做了最终处理(不成立，看log和责任链的具体实现，如果是匹配处理完就返回的，那需要确定吗？如果可以经过多个处理器处理的，那看log？)
**/

// 以下三种责任链思想实现的例子
// 1. golang 基于结构体中带函数的特点实现责任链
// 2. 接口实现责任链
// 3. 责任链结构+接口实现(这里接口实现也可以嵌套例1)

// 1. golang 基于结构体中可带函数的特性实现责任链
// 假设要对搜索请求进行 敏感词过滤、query词纠错、query词归一等(都属于query词改写相关)
// 使用责任链模式解决

// // golang可以用结构体+函数属性间接实现责任链模式，其他语言一般需要依靠接口+继承实现
// // 1.1. 简单的搜索请求和响应结构体
// type SearchReq struct {
// 	keyword        string
// 	rewriteKeyword string
// 	// ...
// }

// type SearchRes struct {
// 	results []string
// 	// ...
// }

// // 1.2. 责任链模式
// type HandleFn func(req *SearchReq)

// type SearchHandler struct {
// 	HandleFn HandleFn
// 	Next     *SearchHandler
// }

// // 1.3. 处理搜索请求责任链
// func ExecuteSearchHandlers(h *SearchHandler, req *SearchReq) {
// 	if h != nil {
// 		h.HandleFn(req)
// 		ExecuteSearchHandlers(h.Next, req)
// 	}
// }

// // 1.4. 实现敏感词过滤
// func FilterHandleFn(req *SearchReq) {
// 	// ...敏感词过滤
// 	// 算是query词改写的一部分，修改后直接改动req.keyword
// 	// req.rewriteKeyword = filter(req.rewriteKeyword)
// }

// // 1.5 实现query词纠错
// func CorrectHandleFn(req *SearchReq) {
// 	// ...纠错操作
// 	// req.rewriteKeyword = correct(req.rewriteKeyword)
// }

// // 1.6 实现query词归一
// func UnitHandleFn(req *SearchReq) {
// 	// ...归一操作，这里具体实现unit(word string)
// 	// req.rewriteKeyword = unit(req.rewriteKeyword)
// }

// func main() {
// 	unitHandler := &SearchHandler{FilterHandleFn, nil}
// 	correctHandler := &SearchHandler{FilterHandleFn, unitHandler}
// 	filterHandler := &SearchHandler{FilterHandleFn, correctHandler}
// 	req := &SearchReq{keyword: "test"}
// 	ExecuteSearchHandlers(filterHandler, req)
// }

// // 2. 接口实现责任链
// // 背景：病人来访时， 他们首先都会去前台， 然后是看医生、 取药， 最后结账。 也就是说， 病人需要通过一条部门链， 每个部门都在完成其职能后将病人进一步沿着链条输送

// // 病人信息(请求)
// type Patient struct {
// 	name              string
// 	registrationDone  bool
// 	doctorCheckUpDone bool
// 	medicineDone      bool
// 	paymentDone       bool
// }

// // 责任链接口
// type Department interface {
// 	Handle(p *Patient)
// 	SetNext(d Department) // 接口本身是指针，因此注意这里不要声明称了指针
// }

// // 前台、医生、取药、结账实现
// type Reception struct{ Next Department }

// func (r *Reception) Handle(p *Patient) {
// 	if p.registrationDone {
// 		fmt.Println("Patient registration already done")
// 		r.Next.Handle(p)
// 		return
// 	}
// 	fmt.Println("Reception registering patient")
// 	p.registrationDone = true
// 	r.Next.Handle(p)
// }

// func (r *Reception) SetNext(next Department) {
// 	r.Next = next
// }

// type Doctor struct{ Next Department }

// func (r *Doctor) Handle(p *Patient) {
// 	if p.doctorCheckUpDone {
// 		fmt.Println("Doctor already done")
// 		return
// 	}
// 	fmt.Println("Doctor checkup patient")
// 	p.doctorCheckUpDone = true
// }

// func (r *Doctor) SetNext(next Department) {
// 	r.Next = next
// }

// // ...省略取药、结账两个实现

// func main() {
// 	d := &Doctor{}
// 	r := &Reception{}
// 	// ...
// 	r.SetNext(d)
// 	r.Handle(&Patient{name: "test"})
// }

// 3. 责任链结构+接口实现(这里接口实现也可以嵌套例1)

// 3.1 责任链中节点的接口声明
// 并且增加一个匹配函数，在进一步，匹配责任链，匹配上了就不继续看下一个了，直接处理后返回
// 这里的接口可以不用 SetNext(next ResponsibilityHandler) 函数，因为我们额外使用责任链结构
type ResponsibilityHandler interface {
	IsMatch(args ...string) bool // 判断当前处理器是否匹配，不匹配则由下一个处理器处理
	Handle(args ...string) string
}

// 3.2 责任链结构声明以及基本函数声明
type ResponsibilityHandlerChain struct {
	chain []ResponsibilityHandler
}

func (rhc *ResponsibilityHandlerChain) RegisterHandler(handler ResponsibilityHandler) {
	if len(rhc.chain) == 0 {
		rhc.chain = []ResponsibilityHandler{handler}
		return
	}
	rhc.chain = append(rhc.chain, handler)
}

func (rhc *ResponsibilityHandlerChain) Handle(args ...string) string {
	for _, hanlder := range rhc.chain {
		if hanlder.IsMatch(args...) {
			return hanlder.Handle(args...)
		}
	}
	return ""
}

type FirstHandler struct{}

func (fh *FirstHandler) IsMatch(args ...string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}

func (fh *FirstHandler) Handle(args ...string) string {
	fmt.Println("第一个处理")
	return "first done"
}

type SecondHandler struct{}

func (sh *SecondHandler) IsMatch(args ...string) bool {
	if len(args) > 0 {
		return true
	}
	return false
}

func (sh *SecondHandler) Handle(args ...string) string {
	fmt.Println("第二个处理")
	return "second done"
}

func main() {
	f := &FirstHandler{}
	s := &SecondHandler{}
	c := &ResponsibilityHandlerChain{}
	c.RegisterHandler(f)
	c.RegisterHandler(s)
	c.Handle()
	c.Handle("123")
}
