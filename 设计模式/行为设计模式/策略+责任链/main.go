package main

import "fmt"

// 实现一个基于责任链和策略模式的搜索重排框架demo
// 这里责任链模式中的处理方法，限制为策略模式接口
// 因此该责任链相关的处理接口都是同一职能(重排)，与其他责任链实现不同，其他责任链中的处理接口性质不一定相同(职能不一定相同)

// 1. 策略模式方法接口
type ReorderStrategy interface {
	StrategyFunc(products []string) []string
}

// 第一个实现策略接口的策略类
type SalesReorderStrategy struct{}

var srs SalesReorderStrategy

func (srs *SalesReorderStrategy) StrategyFunc(products []string) []string {
	fmt.Println("按销量重排")
	// ...对产品销量的排序处理
	return products
}

// 第二个实现策略接口的策略类
type PriceReorderStrategy struct{}

var prs PriceReorderStrategy

func (prs *PriceReorderStrategy) StrategyFunc(products []string) []string {
	fmt.Println("按产品价格重排")
	// ...对产品价格的排序处理
	return products
}

// 2. 责任链模式基础实现
type ResponsibilityNode struct {
	Strategy ReorderStrategy // 当前节点要处理的策略
	Next     *ResponsibilityNode
}

// 责任链处理
func HandleNodes(node *ResponsibilityNode, products []string) []string {
	if node != nil {
		products = node.Strategy.StrategyFunc(products)
		products = HandleNodes(node.Next, products)
	}
	return products
}

func main() {
	products := []string{"123", "234", "456", "678"}
	r2 := &ResponsibilityNode{&prs, nil}
	r1 := &ResponsibilityNode{&srs, r2}
	HandleNodes(r1, products)
	fmt.Print("end")
}
