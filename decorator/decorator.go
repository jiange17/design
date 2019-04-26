package decorator

// 装饰模式
// 给面包加奶油，糖计算加完之后的总价钱。

type Compoment interface {
	Cost() int
}

type BreadCompoment struct {
	money int
}

func (c *BreadCompoment) Cost() int {
	return c.money
}

func NewBreadCompoment(money int) Compoment {
	return &BreadCompoment{money: money}
}

type CreamDecorator struct {
	Compoment
	money int
}

func (c *CreamDecorator) Cost() int {
	return c.Compoment.Cost() + c.money
}

func WrapCreamDecorator(c Compoment, money int) Compoment {
	return &CreamDecorator{
		Compoment: c,
		money:     money,
	}
}

type SugarDecorator struct {
	Compoment
	money int
}

func (s *SugarDecorator) Cost() int {
	return s.Compoment.Cost() + s.money
}

func WrapSugarDecorator(c Compoment, money int) Compoment {
	return &SugarDecorator{
		Compoment: c,
		money:     money,
	}
}
