package factory

import "method/phone"

type xiaomiFactory struct{}

func (xf *xiaomiFactory) CreatePhone() phone.Phone {
	return phone.NewXiaomi()
}
