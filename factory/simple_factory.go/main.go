package main

import "fmt"

type iPhone struct{}
type huawei struct{}
type xiaomi struct{}

func NewIphone() *iPhone {
	return &iPhone{}
}

func NewHuawei() *huawei {
	return &huawei{}
}

func NewXiaomi() *xiaomi {
	return &xiaomi{}
}

func main() {
	iphone := NewIphone()
	fmt.Printf("iphone: %v\n", iphone)

	huawei := NewHuawei()
	fmt.Printf("hauwei: %v\n", huawei)

	xiaomi := NewXiaomi()
	fmt.Printf("xiaomi: %v\n", xiaomi)
}
