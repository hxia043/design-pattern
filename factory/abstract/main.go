package main

import "abstract/phone"

func main() {
	var p phone.Phone

	p = phone.NewPhone("iphone")
	p.SayHello()

	p = phone.NewPhone("xiaomi")
	p.SayHello()

	p = phone.NewPhone("huawei")
	p.SayHello()
}
