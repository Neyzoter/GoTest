package main

import "fmt"

type Fruit interface {
	GetWeight() float32
}

type FruitPlus interface {
	GetWeight() float32
	GetName() string
}

type Apple struct {
	Weight float32
	Name string
}

func (a Apple) GetWeight() float32 {
	return a.Weight
}

func (a Apple) GetName() string {
	return a.Name
}

// 类型断言是一个使用在接口值上的操作
// x.(T)：x为接口，T为类型或者接口
// 如果T为类型，通过断言则返回T类型，可以调用T类型的所有方法
// 如果T为接口，通过断言则返回T接口，可以调用T接口的所有方法
func main() {

	var fruit Fruit
	fruit = Apple{2, "Apple"}
	// 无法调用GetName，因为fruit是一个Fruit类型的接口，而不是Apple类型
	fruit.GetWeight()
	// 检查fruit的动态类型是否是Apple
	fruitCopy := fruit.(Apple)
	fmt.Printf("%T\n", fruitCopy)
	// 断言后，fruitCopy实际类型为Apple，可以调用GetName
	fmt.Println(fruitCopy.GetWeight(), " ", fruitCopy.GetName())

	fruitPlus := fruit.(FruitPlus)
	fmt.Printf("%T\n", fruitPlus) // 虽然显示的是Apple类型，但是实际上是接口
	// 断言后，fruitCopy实际类型为Apple，可以调用GetName
	fmt.Println(fruitPlus.GetWeight(), " ", fruitPlus.GetName())
}
