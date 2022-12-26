package main

import "fmt"

type Phone interface {
	sayHello()
}

type iPhone struct {
	name string
}

type huawei struct {
	name string
}

type xiaomi struct {
	name string
}

func (ip *iPhone) sayHello() {
	fmt.Println("hello", ip.name)
}

func (hw *huawei) sayHello() {
	fmt.Println("hello", hw.name)
}

func (xm *xiaomi) sayHello() {
	fmt.Println("hello", xm.name)
}

func NewIphone() Phone {
	return &iPhone{name: "iphone"}
}

func NewHuawei() Phone {
	return &huawei{name: "huawei"}
}

func NewXiaomi() Phone {
	return &xiaomi{name: "xiaomi"}
}

func main() {
	var phone Phone

	phone = NewIphone()
	phone.sayHello()

	phone = NewHuawei()
	phone.sayHello()

	phone = NewXiaomi()
	phone.sayHello()
}
