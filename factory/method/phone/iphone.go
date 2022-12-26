package phone

import "fmt"

type iPhone struct {
	name string
}

func (ip *iPhone) SayHello() {
	fmt.Println("hello", ip.name)
}

func NewIphone() *iPhone {
	return &iPhone{name: "iphone"}
}
