package main

import "fmt"

/**
抽象工厂模式介绍：
抽象工厂模式 --- 创建型模式
思想：能创建一系列相关的对象， 而无需指定其具体类

适用场景：
1. 需要创建一族相关或相互依赖的产品对象：抽象工厂模式适用于创建多个相关的产品对象，这些产品对象之间存在一定的关联或依赖关系
2. 需要保证产品对象的一致性和兼容性：抽象工厂模式可以确保创建的产品对象属于同一族系列，保证它们在使用时的一致性和兼容性
3. 需要灵活替换具体工厂类：抽象工厂模式可以方便地替换具体工厂类，从而改变所创建的产品族。如果需要新增或替换一套产品族，只需新增或替换对应的具体工厂类即可，不需要修改已有的客户端代码。这种扩展性和灵活性使得抽象工厂模式在变化较频繁的系统中应用广泛

具体应用场景：
1. 创建不同操作系统下的图形界面组件，如按钮、文本框和下拉列表等，每个操作系统都有自己的一套组件风格和行为，抽象工厂模式可以为每个操作系统提供一个具体工厂，用于创建对应的图形界面组件。
2. 创建不同数据库类型的连接对象和操作对象，每个数据库类型有自己的实现方式，通过抽象工厂模式可以为每个数据库类型提供一个具体工厂，用于创建对应的连接对象和操作对象。
3. 日志记录器：根据配置，创建文件日志记录器或数据库日志记录器。
4. 加密算法：根据需求，提供不同类型的加密算法对象

抽象工厂模式的优缺点
优点：
-- 提供了一种封装对象创建的方式，将具体产品的创建逻辑与客户端代码解耦，使得客户端代码更加灵活和可维护。
-- 可以确保一族产品对象的兼容性，因为一族产品对象是由同一个工厂创建的，保证了它们之间的一致性。
-- 可以方便地替换具体工厂类，从而改变所创建的产品族，满足不同的需求。
-- 符合开闭原则，对扩展开放，对修改关闭。当需要增加新的产品族时，只需要扩展抽象工厂和具体工厂，而无需修改已有代码。
缺点：
-- 增加了系统的复杂性，引入了更多的类和接口，增加了代码的量和维护的难度。
-- 对于新增的产品对象的支持不够灵活，如果需要新增一个产品对象，需要修改抽象工厂和所有的具体工厂类
**/

// 在工厂方法模式的背景基础上，多引入子弹这个东西，因为抽象工厂模式的目的主要在于创建一族相关或相互依赖的产品对象，而不是单独创建某个产品对象
// 多家一类关联对象，更符合实际情况以及展示示例

// 在这个示例代码中，我们定义了 Gun 接口和 Bullet 接口来分别表示枪支和子弹的行为。
// 然后，我们定义了 AbstractFactory 接口，它包含了创建枪支和子弹的方法。
// 接下来，我们实现了具体的枪支类型 AK47 和 M4，它们分别实现了 Gun 接口的 Fire 方法。
// 然后，我们实现了具体的子弹类型 AK47Bullet 和 M4Bullet，它们分别实现了 Bullet 接口的 Load 方法。
// 接着，我们创建了两个工厂类：AK47Factory 和 M4Factory，它们分别实现了 AbstractFactory 接口。AK47Factory 负责创建 AK47 枪支对象和 AK47 对应的子弹对象，而 M4Factory 则负责创建 M4 枪支对象和 M4 对应的子弹对象。
// 在 main 函数中，我们通过具体的工厂类分别创建了 AK47 枪支对象和子弹对象，以及 M4 枪支对象和子弹对象。然后，我们分别调用它们的方法，模拟了开火和装填子弹的过程。
// 通过使用抽象工厂模式，我们可以在不修改客户端代码的情况下，根据具体的工厂类创建一系列相关的产品对象（比如枪支和子弹）。这种模式可以提供更高的灵活性和可扩展性，并且符合开闭原则，使得系统的设计更具可维护性和可扩展性。
// 1. 声明 Gun 和 Bullet 接口，并分别实现
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

type Bullet interface {
	Load()
}

type AK47Bullet struct{}

func (b *AK47Bullet) Load() {
	fmt.Println("AK47 子弹装填完成！")
}

type M4Bullet struct{}

func (b *M4Bullet) Load() {
	fmt.Println("M4 子弹装填完成！")
}

// 2. 抽象工厂接口及其实现类
type AbstractFactory interface {
	CreateGun() Gun
	CreateBullet() Bullet
}

type AK47Factory struct{}

func (f *AK47Factory) CreateGun() Gun {
	return &AK47{}
}

func (f *AK47Factory) CreateBullet() Bullet {
	return &AK47Bullet{}
}

type M4Factory struct{}

func (f *M4Factory) CreateGun() Gun {
	return &M4{}
}

func (f *M4Factory) CreateBullet() Bullet {
	return &M4Bullet{}
}

func main() {
	ak47Factory := &AK47Factory{}
	m4Factory := &M4Factory{}

	ak47 := ak47Factory.CreateGun()
	ak47Bullet := ak47Factory.CreateBullet()

	m4 := m4Factory.CreateGun()
	m4Bullet := m4Factory.CreateBullet()

	ak47.Fire()
	ak47Bullet.Load()

	m4.Fire()
	m4Bullet.Load()
}
