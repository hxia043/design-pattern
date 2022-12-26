package factory

import "method/phone"

type huaweiFactory struct{}

func (hf *huaweiFactory) CreatePhone() phone.Phone {
	return phone.NewHuawei()
}
