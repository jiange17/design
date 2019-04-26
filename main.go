package main

import (
	"fmt"

	"xcq.org/design/decorator"
	"xcq.org/design/observer"
	"xcq.org/design/strategy"
)

func TestObserver() {
	kafka := observer.NewKafka()
	xcq := &observer.ObserverXcq{
		ID: 25,
	}
	kafka.AddObserver(xcq)
	kafka.SetMessage("你好呀，good boy！")
	kafka.RemoveObserver(xcq)
	kafka.SetMessage("拜拜，收不到了。")
}

func TestStrategy() {
	xcq := strategy.NewPeople()
	xcq.Weapon.Display()
	xcq.Weapon = strategy.NewSword()
	xcq.Weapon.Display()
}

func TestDecorator() {
	food := decorator.NewBreadCompoment(1)
	fmt.Println(food.Cost())
	food = decorator.WrapCreamDecorator(food, 2)
	fmt.Println(food.Cost())
	food = decorator.WrapSugarDecorator(food, 3)
	fmt.Println(food.Cost())
}

func main() {
	TestObserver()
	TestStrategy()
	TestDecorator()
}
