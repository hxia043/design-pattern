package phone

import "log"

type Phone interface {
	SayHello()
}

func NewPhone(name string) Phone {
	switch name {
	case "iphone":
		return NewIphone()
	case "huawei":
		return NewHuawei()
	case "xiaomi":
		return NewXiaomi()
	default:
		log.Fatalln("unknown phone name", name)
	}

	return nil
}
