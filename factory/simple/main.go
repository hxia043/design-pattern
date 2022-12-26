package main

import (
	"fmt"
	"simple/phone"
)

func main() {
	iphone := phone.NewIphone()
	fmt.Printf("iphone: %v\n", iphone)

	huawei := phone.NewHuawei()
	fmt.Printf("hauwei: %v\n", huawei)

	xiaomi := phone.NewXiaomi()
	fmt.Printf("xiaomi: %v\n", xiaomi)
}
