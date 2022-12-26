package phone

import "fmt"

type xiaomi struct {
	name string
}

func (xm *xiaomi) SayHello() {
	fmt.Println("hello", xm.name)
}

func NewXiaomi() *xiaomi {
	return &xiaomi{name: "xiaomi"}
}
