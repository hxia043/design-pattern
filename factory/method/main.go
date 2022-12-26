package main

import (
	"method/factory"
	"method/phone"
)

func main() {
	var p phone.Phone
	var f factory.Factory

	f = factory.NewFactory("iphone")
	p = f.CreatePhone()
	p.SayHello()

	f = factory.NewFactory("huawei")
	p = f.CreatePhone()
	p.SayHello()

	f = factory.NewFactory("xiaomi")
	p = f.CreatePhone()
	p.SayHello()
}
