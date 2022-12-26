package main

import (
	"fmt"
	"log"
)

type Phone interface {
	sayHello()
}

type Factory interface {
	createPhone() Phone
}

type iPhoneFactory struct{}
type huaweiFactory struct{}
type xiaomiFactory struct{}

func (ifact *iPhoneFactory) createPhone() Phone {
	return &iPhone{name: "iphone"}
}

func (hfact *huaweiFactory) createPhone() Phone {
	return &huawei{name: "huawei"}
}

func (xfact *xiaomiFactory) createPhone() Phone {
	return &xiaomi{name: "xiaomi"}
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

func NewFactory(name string) Factory {
	switch name {
	case "iphone":
		return &iPhoneFactory{}
	case "huawei":
		return &huaweiFactory{}
	case "xiaomi":
		return &xiaomiFactory{}
	default:
		log.Fatalln("unknow factory", name)
	}

	return nil
}

func main() {
	var phone Phone
	var factory Factory

	factory = NewFactory("iphone")
	phone = factory.createPhone()
	phone.sayHello()

	factory = NewFactory("huawei")
	phone = factory.createPhone()
	phone.sayHello()

	factory = NewFactory("xiaomi")
	phone = factory.createPhone()
	phone.sayHello()
}
