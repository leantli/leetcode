package main

import "fmt"

/**
命令模式介绍：
命令模式 --- 行为设计模式
思想：将请求封装成一个对象，该转换让你能根据不同的请求将方法参数化、 延迟请求执行或将其放入队列中， 且能实现可撤销操作。
命令模式使得可以将请求的发送者与接收者解耦

适用场景：
1. 希望将请求发送者和接收者解耦：命令模式通过将请求封装成一个对象，使得发送者不需要知道请求的具体接收者，从而实现解耦。发送者只需通过命令对象来发起请求，而具体的接收者将由命令对象来处理。
2. 希望在不同的时间点、顺序或条件下请求执行：命令模式允许将请求参数化，可以根据不同的时间、顺序或条件来执行请求。这使得可以轻松地实现请求的排队、延迟执行、撤销和重做等功能。
3. 希望支持事务和日志记录：命令模式可以很好地支持事务的操作。将一系列操作封装成命令对象，可以在执行前进行验证和准备，并且可以通过事务管理器来管理多个命令的执行。同时，命令模式也可以方便地记录请求日志，以便后续分析、恢复或审计。
4. 希望支持回调和撤销操作：命令模式可以通过命令对象的回调方法来实现对请求的响应。这使得可以在命令对象中调用接收者的方法，并将结果返回给发送者。此外，命令模式还可以支持撤销操作，通过保存命令对象的历史记录，可以撤销已执行的命令。

具体应用场景：
1. 日志记录器链：不同类型的日志记录器可以按照顺序记录日志，每个记录器可以决定是否继续传递给下一个记录器。
2. 数据验证链：不同的验证器可以逐个验证数据的合法性，每个验证器可以决定是否继续传递给下一个验证器。
3. 过滤器/拦截器链：（过滤器链）多个请求过滤器可以按照顺序对请求进行过滤，每个过滤器可以决定是否继续传递给下一个过滤器。
4. 购买流程链：购买流程中的不同环节可以作为处理者，按照顺序处理购买请求，例如验证库存、计算价格、生成订单等。

命令模式的优缺点
优点：
-- 解耦处理逻辑，将请求发送者和接收者解耦，请求发送者无需知道具体的接受者和处理细节，只需统一调用接口，上层无需知道下层具体实现，这种松耦合有助于减少对象之间的依赖关系，符合迪米特法则和依赖倒置原则
-- 可重用的处理逻辑，每个处理者都专注于自己的处理逻辑，使得处理逻辑可以单独测试、维护和复用，可以减少重复代码，提高代码可读性和可维护性，同样符合单一职责原则
-- 灵活性和可拓展性，可以在不更改现有代码的情况下灵活地增加处理者，符合开闭原则
-- 支持撤销、恢复和事务(组合)，其可以记录命令的执行历史，从而支持撤销和重做操作。同时，它也可以支持将一系列命令组合成事务，保证多个命令的原子性操作
-- 支持请求的队列和延迟执行，其可以将请求进行排队，以便按照顺序执行。还可以实现延迟执行，即将命令对象保存起来，在需要时再执行
缺点：
-- 增加了类的数量：引入命令对象会增加系统中的类数量，特别是在有大量命令类的情况下，可能会导致类的膨胀
-- 可能引入额外的复杂性：命令模式引入了额外的抽象层，可能会增加代码的复杂性和理解的难度

和策略模式的区别：
命令模式强调将请求（命令）封装成对象，以及命令的执行者和发送者的解耦。它通过命令对象和命令调用者之间的关系来实现这一点
策略模式强调将算法或策略封装成对象，并使其可互换。它通过上下文对象和具体策略类之间的关系来实现这一点
**/

// // 例子背景
// // 下面我们通过电视机的例子来了解命令模式。 你可通过以下方式打开电视机：
// // 按下遥控器上的 ON 开关； 按下电视机上的 ON 开关。
// // 基于命令模式实现以上需求
// // 我们可以从实现 ON 命令对象并以电视机作为接收者入手。 当在此命令上调用 execute 执行方法时， 方法会调用 TV.on 打开电视函数。
// // 最后的工作是定义请求者： 这里实际上有两个请求者： 遥控器和电视机。 两者都将嵌入 ON 命令对象。
// // 注意我们是如何将相同请求封装进多个请求者的。 我们也可以采用相同的方式来处理其他命令。
// // 创建独立命令对象的优势在于可将 UI 逻辑与底层业务逻辑解耦。 这样就无需为每个请求者开发不同的处理者了。
// // 命令对象中包含执行所需的全部信息， 所以也可用于延迟执行。

// // 1. 命令接口及其具体实现
// type Command interface {
// 	execute()
// }
// type OnCommand struct {
// 	Device Device
// }
// type OffCommand struct {
// 	Device Device
// }

// func (c *OnCommand) execute() {
// 	c.Device.on()
// }

// func (c *OffCommand) execute() {
// 	c.Device.off()
// }

// // 2. 接收者接口及其具体实现
// type Device interface {
// 	on()
// 	off()
// }
// type TV struct {
// 	IsRunning bool
// }

// func (t *TV) on() {
// 	t.IsRunning = true
// 	fmt.Println("turn on")
// }
// func (t *TV) off() {
// 	t.IsRunning = false
// 	fmt.Println("turn off")
// }

// // 3. 请求者，一个结构体中带有命令的接口
// type Button struct {
// 	command Command
// }

// func (b *Button) press() {
// 	b.command.execute()
// }

// func main() {
// 	tv := &TV{} // 接收者
// 	// 两个具体命令实现类
// 	onCommand := &OnCommand{
// 		Device: tv,
// 	}
// 	offCommand := &OffCommand{
// 		Device: tv,
// 	}

// 	// 请求者
// 	onButton := &Button{
// 		command: onCommand,
// 	}
// 	onButton.press()
// 	offButton := &Button{
// 		command: offCommand,
// 	}
// 	offButton.press()
// }

/**
例子二：实现一个商品分类筛选功能，用户可以通过命令来选择不同的筛选条件，然后系统会根据用户的选择展示相应的商品列表
这个的实现和策略模式基本差不多，但注重场景不同
**/

// 1. 首先定义一个Command命令接口，包含一个执行命令的方法execute()
type Command interface {
	Execute()
}

// 2. 来实现多个具体的命令类，每个类代表一种筛选条件
type CategoryFilterCommand struct {
	// ...省略属性
}

func (c *CategoryFilterCommand) Execute() {
	// 执行商品分类筛选操作，比如从数据库或缓存中获取符合筛选条件的商品列表并进行展示
	fmt.Println("执行商品分类筛选操作")
}

type PriceFilterCommand struct {
	// ...
}

func (c *PriceFilterCommand) Execute() {
	fmt.Println("执行价格筛选操作")
}

// 3. 创建命令调用者(请求者)
type CommandInvoker struct {
	command Command
}

func (ci *CommandInvoker) SetCommand(command Command) {
	ci.command = command
}

func (ci *CommandInvoker) ExecuteCommand() {
	if ci.command != nil {
		ci.command.Execute()
	}
}

func main() {
	invoker := &CommandInvoker{}

	// 用户选择商品分类筛选
	categoryFilterCommand := &CategoryFilterCommand{}
	invoker.SetCommand(categoryFilterCommand)
	invoker.ExecuteCommand()

	// 用户选择价格筛选
	priceFilterCommand := &PriceFilterCommand{}
	invoker.SetCommand(priceFilterCommand)
	invoker.ExecuteCommand()
}
