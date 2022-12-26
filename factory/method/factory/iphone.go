package factory

import "method/phone"

type iphoneFactory struct{}

func (iP *iphoneFactory) CreatePhone() phone.Phone {
	return phone.NewIphone()
}
