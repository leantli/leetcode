package main

import "fmt"

/**
装饰器模式介绍：
装饰器模式 --- 创建型模式
思想：通过创建一个包装器（装饰器）来包裹原始对象，以增强原始对象的功能，允许在不修改现有对象结构的情况下，动态地向对象添加额外的行为或功能

适用场景：
1. 当需要在不修改现有对象代码的情况下，动态地添加额外的功能或行为时，可以使用装饰器模式。它提供了一种灵活的方式来扩展对象的功能，而无需修改现有代码
2. 当存在多个不同的功能组合，且每种组合都需要独立扩展时，装饰器模式可以提供更好的组合灵活性，避免了类爆炸的问题
3. 当希望对一个对象的功能进行多层次的包装，以实现更复杂的功能组合时，装饰器模式可以提供更高层次的灵活性和可扩展性

具体应用场景：
1. I/O流处理，如Java中的BufferedInputStream、BufferedOutputStream等。
2. Web开发中的中间件，如身份验证、日志记录、缓存等功能的动态添加。
3. GUI界面中的装饰性控件，如添加边框、滚动条等。
4. 日志记录、性能监测、数据缓存等功能的添加。

装饰器模式的优缺点
优点：
-- 无需修改现有对象的结构和代码，即可实现功能的扩展和修改
-- 你可以在运行时添加或删除对象的功能。
-- 可以灵活地组合多个装饰器，以实现不同的功能组合
-- 符合开闭原则，代码可扩展性强
缺点：
-- 对于过多的装饰器叠加，代码可读性可能会降低
-- 使用不当可能导致装饰器与原始对象之间的紧耦合
-- 装饰器模式不适合对对象的核心功能进行修改，而应该关注对对象功能的扩展和增强

和代理模式的区别：
虽然两个模式都共用同一接口和包装实现，并且都能够不改原代码而拓展额外功能或责任
但是装饰器模式侧重于功能扩展和组合，而代理模式侧重于对对象访问的控制
以下例子装饰器可以反复叠加，代理模式则一般不这样
**/

// 以下以一个披萨为例使用装饰器模式
// 装饰对象和包装器都共用同一接口
type IPizza interface {
	getPrice() int
}

// 具体装饰对象
type VeggeMania struct{}

func (vm *VeggeMania) getPrice() int { return 15 }

// 具体装饰
type TomatoTopping struct{ pizza IPizza }

func (tt *TomatoTopping) getPrice() int { return tt.pizza.getPrice() + 7 }

type CheeseTopping struct{ pizza IPizza }

func (ct *CheeseTopping) getPrice() int { return ct.pizza.getPrice() + 10 }

func main() {
	pizza := &VeggeMania{}
	// 增加一个装饰
	pizzaWithCheese := &CheeseTopping{pizza}
	// 增加一个装饰
	pizzaWithCheeseAndTomato := &TomatoTopping{pizzaWithCheese}
	fmt.Println(pizzaWithCheeseAndTomato.getPrice()) // 32
}
