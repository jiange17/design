package strategy

import (
	"fmt"
)

// 策略模式
// 模拟给一个人装备不同的装备(武器)
type People struct {
	Weapon Weapon
}

func NewPeople() *People {
	return &People{
		Weapon: NewKnife(),
	}
}

func NewKnife() *Knife {
	return &Knife{}
}

type Weapon interface {
	Display()
}

type Knife struct {
}

func (k *Knife) Display() {
	fmt.Println("this is a knife.")
}

func NewSword() *Sword {
	return &Sword{}
}

type Sword struct {
}

func (s *Sword) Display() {
	fmt.Println("this is a sword.")
}
