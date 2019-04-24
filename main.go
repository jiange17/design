package main

import (
	"xcq.org/design/observer"
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

func main() {
	TestObserver()
}