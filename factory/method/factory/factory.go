package factory

import (
	"log"
	"method/phone"
)

type Factory interface {
	CreatePhone() phone.Phone
}

func NewFactory(name string) Factory {
	switch name {
	case "iphone":
		return &iphoneFactory{}
	case "huawei":
		return &huaweiFactory{}
	case "xiaomi":
		return &xiaomiFactory{}
	default:
		log.Fatalln("unknown factory")
	}

	return nil
}
