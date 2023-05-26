package main

import "fmt"

/**
工厂模式介绍：
工厂模式 --- 创建型模式
思想：定义了一个方法， 且必须使用该方法代替通过直接调用构造函数来创建对象 （ new操作符） 的方式。
解决了在不指定具体类的情况下创建产品对象的问题；；子类可重写该方法来更改将被创建的对象所属类。

适用场景：
1. 当无法预知对象确切类别及其依赖关系，需要根据特定条件创建不同对象时，可以使用工厂模式
2. 需要通过工厂类来创建对象，以降低客户端与具体对象之间的耦合度时，可以使用工厂模式
3. 当需要扩展或变化对象的创建逻辑时，可以使用工厂模式

具体应用场景：
1. 在一个图形绘制应用中，根据用户选择的图形类型（如圆形、矩形、三角形）创建相应的图形对象。
2. 在一个日志记录器的应用中，根据配置文件中的设置创建不同类型的日志记录器（如文件日志记录器、数据库日志记录器）
3. 在一个电商平台中，客户端需要通过工厂类创建不同类型的支付方式对象（如支付宝支付、微信支付），从而与具体的支付方式实现解耦。
4. 在一个游戏中，通过工厂类创建不同类型的敌人对象，使得游戏逻辑与具体敌人类型的实现解耦。
5. 在一个汽车制造工厂中，根据客户的订单和规格要求，使用具体的工厂类创建不同型号和配置的汽车对象。
6. 在一个食品配送系统中，根据不同地区的需求和规定，使用具体的工厂类创建适合该地区的食品对象

工厂模式的优缺点
优点：
-- 你可以避免创建者和具体产品之间的紧密耦合
-- 单一职责原则。 你可以将产品创建代码放在程序的单一位置， 从而使得代码更容易维护
-- 开闭原则。 无需更改现有客户端代码， 你就可以在程序中引入新的产品类型
缺点：
-- 应用工厂方法模式需要引入许多新的子类， 代码可能会因此变得更复杂。 最好的情况是将该模式引入创建者类的现有层次结构中
**/

// 由于 Go 中缺少类和继承等 OOP 特性， 所以无法使用 Go 来实现经典的工厂方法模式。 不过， 我们仍然能实现模式的基础版本， 即简单工厂
// 例子背景：使用工厂结构体来构建多种类型的武器

// // 一！简单工厂模式
// // 在下面的示例中，我们定义了一个 Gun 接口，包含了 Fire 方法。然后，我们实现了 Pistol 和 Rifle 两种类型的枪支，它们都满足 Gun 接口。
// // 接着，我们创建了一个 GunFactory 工厂，其中的 CreateGun 方法接收一个 gunType 参数，根据不同的类型创建相应的枪支对象。
// // 最后，在 main 函数中通过工厂创建了不同类型的枪支并进行了射击。
// // 通过工厂模式，我们可以根据需要创建不同类型的枪支对象，而不需要直接与具体的枪支类型进行耦合。这种设计方式符合开闭原则，使得代码更加灵活和可扩展。
// // 1. 枪支接口及其实现类
// type Gun interface {
// 	Fire()
// }
// type Pistol struct{}

// func (p Pistol) Fire() {
// 	fmt.Println("Pistol: Bang!")
// }

// type Rifle struct{}

// func (r Rifle) Fire() {
// 	fmt.Println("Rifle: Bang bang bang!")
// }

// // 2. 工厂类及其方法实现
// type GunFactory struct{}

// // CreateGun 是 GunFactory 的方法，根据传入的 gunType 创建相应的枪支对象
// func (f GunFactory) CreateGun(gunType string) Gun {
// 	switch gunType {
// 	case "pistol":
// 		return Pistol{}
// 	case "rifle":
// 		return Rifle{}
// 	default:
// 		return nil
// 	}
// }

// func main() {
// 	factory := GunFactory{}

// 	pistol := factory.CreateGun("pistol")
// 	pistol.Fire() // Output: Pistol: Bang!

// 	rifle := factory.CreateGun("rifle")
// 	rifle.Fire() // Output: Rifle: Bang bang bang!
// }

// 二！工厂方法模式
// 在这个示例代码中，我们定义了 Gun 接口来表示枪支的行为，其中只有一个 Fire 方法。然后，我们定义了 GunFactory 接口，它包含一个 CreateGun 方法，用于创建具体的枪支对象
// 接下来，我们实现了两种具体的枪支类型：AK47 和 M4，它们分别实现了 Gun 接口的 Fire 方法
// 然后，我们创建了两个工厂类：AK47Factory 和 M4Factory，它们分别实现了 GunFactory 接口的 CreateGun 方法，用于创建对应的枪支对象
// 在 main 函数中，我们分别使用 AK47Factory 和 M4Factory 创建了具体的枪支对象，然后调用它们的 Fire 方法来进行开火操作
// 通过将创建对象的责任委托给具体的工厂子类，我们将示例代码从简单工厂模式扩展为工厂方法模式。每个具体的工厂子类负责创建一种类型的枪支对象，实现了更高的灵活性和可扩展性
// 乍看之下， 这种更改可能毫无意义： 我们只是改变了程序中调用构造函数的位置而已。 但是， 仔细想一下， 现在你可以在子类中重写工厂方法， 从而改变其创建产品的类型。
// 但有一点需要注意:仅当这些产品具有共同的基类或者接口时， 子类才能返回不同类型的产品， 同时基类中的工厂方法还应将其返回类型声明为这一共有接口。
// 1. 枪的接口及其实现类
type Gun interface {
	Fire()
}

type AK47 struct{}

func (a *AK47) Fire() {
	fmt.Println("AK47 开火！")
}

type M4 struct{}

func (m *M4) Fire() {
	fmt.Println("M4 开火！")
}

// 2. 枪支工厂的接口及其实现类
type GunFactory interface {
	CreateGun() Gun
}

type AK47Factory struct{}

func (f *AK47Factory) CreateGun() Gun {
	return &AK47{}
}

type M4Factory struct{}

func (f *M4Factory) CreateGun() Gun {
	return &M4{}
}

func main() {
	ak47Factory := &AK47Factory{}
	m4Factory := &M4Factory{}

	ak47 := ak47Factory.CreateGun()
	m4 := m4Factory.CreateGun()

	ak47.Fire()
	m4.Fire()
}
