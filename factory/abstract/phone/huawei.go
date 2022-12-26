package phone

import "fmt"

type huawei struct {
	name string
}

func (hw *huawei) SayHello() {
	fmt.Println("hello", hw.name)
}

func NewHuawei() *huawei {
	return &huawei{name: "huawei"}
}
