package main

import (
	"fmt"
)

type Builder interface {
	createDoor() Builder
	createWindow() Builder
	createChair() Builder
	build() (*house, error)
}

type villaBuilder struct {
	house
}

type residenceBuilder struct {
	house
}

type house struct {
	door   int
	window int
	chair  int
}

func (vb *villaBuilder) createDoor() Builder {
	vb.door = 10
	return vb
}

func (vb *villaBuilder) createWindow() Builder {
	vb.window = 50
	return vb
}

func (vb *villaBuilder) createChair() Builder {
	vb.chair = 50
	return vb
}

func (vb *villaBuilder) validation() error {
	return nil
}

func (vb *villaBuilder) build() (*house, error) {
	// validate property of object houseBuilder, skip...
	err := vb.validation()

	vb.createDoor()
	vb.createWindow()
	vb.createChair()

	return &vb.house, err
}

func (rb *residenceBuilder) createDoor() Builder {
	rb.door = 2
	return rb
}

func (rb *residenceBuilder) createWindow() Builder {
	rb.window = 2
	return rb
}

func (rb *residenceBuilder) createChair() Builder {
	rb.chair = 1
	return rb
}

func (rb *residenceBuilder) validation() error {
	return nil
}

func (rb *residenceBuilder) build() (*house, error) {
	// validate property of object carBuilder, skip...
	err := rb.validation()

	rb.createDoor()
	rb.createWindow()
	rb.createChair()

	return &rb.house, err
}

func NewBuilder(typ string) Builder {
	switch typ {
	case "villa":
		return &villaBuilder{}
	case "residence":
		return &residenceBuilder{}
	default:
		return nil
	}
}

func main() {
	house, err := NewBuilder("villa").build()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(house.door)
}
