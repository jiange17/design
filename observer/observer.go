package observer

import (
	"fmt"
)

// 采用观察者模式
// 简单模拟消息队列，发消息给订阅者

type Kafka struct {
	Observers []Observer
	Message   string
}

func NewKafka() *Kafka {
	return &Kafka{
		Observers: make([]Observer, 0),
		Message:   "hello world",
	}
}

func (k *Kafka) AddObserver(ob Observer) {
	k.Observers = append(k.Observers, ob)
	ob.Update(k.Message)
}

func (k *Kafka) RemoveObserver(ob Observer) {
	tempObservers := make([]Observer, 0)
	for _, observer := range k.Observers {
		if observer.GetID() == ob.GetID() {
			continue
		}
		tempObservers = append(tempObservers, observer)
	}
	k.Observers = tempObservers
}

func (k *Kafka) Notify() {
	for _, observer := range k.Observers {
		observer.Update(k.Message)
	}
}

func (k *Kafka) GetMessage() string {
	return k.Message
}

func (k *Kafka) SetMessage(message string) {
	k.Message = message
	k.Notify()
}

type Observer interface {
	Update(message string)
	GetID() int
}

type ObserverXcq struct {
	ID int
}

func (o *ObserverXcq) Update(message string) {
	fmt.Printf("xcq get message: %v\n", message)
}

func (o *ObserverXcq) GetID() int {
	return o.ID
}

// func main() {
// 	kafka := NewKafka()
// 	xcq := &ObserverXcq{
// 		ID: 25,
// 	}
// 	kafka.AddObserver(xcq)
// 	kafka.SetMessage("你好呀，good boy！")
// 	return
// }
