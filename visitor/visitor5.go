package main

import (
	"errors"
	"fmt"
)

type VisitorFunc func(*man) error

func happyVisitor(m *man) error {
	fmt.Println(m.name)
	return nil
}

func sadVisitor(m *man) error {
	fmt.Println(m.age)
	return nil
}

func successVisitor(m *man) error {
	fmt.Println(m.sex)
	return nil
}

func validationFunc(m *man) error {
	if m.name == "" {
		return errors.New("empty name")
	}

	return nil
}

type man struct {
	name string
	age  int
	sex  string
}

type Visitor interface {
	Visit(VisitorFunc) error
}

func (m *man) Visit(fn VisitorFunc) error {
	fmt.Println("in man")

	if err := fn(m); err != nil {
		return err
	}

	fmt.Println("out man")
	return nil
}

type validationVisitor struct {
	visitor Visitor
}

func (v validationVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in validation")

		if m.name == "" {
			return errors.New("empty name")
		}

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println("out validation")
		return nil
	})
}

type errorVisitor struct {
	visitor Visitor
}

func (v errorVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in error")

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println("out error")
		return nil
	})
}

type ageVisitor struct {
	visitor Visitor
}

func (v ageVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(m *man) error {
		fmt.Println("in age")

		if err := fn(m); err != nil {
			return err
		}

		fmt.Println(m.name, m.age)

		fmt.Println("out age")
		return nil
	})
}

type VisitorList []Visitor

func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(fn); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var visitor Visitor

	m1 := &man{name: "hxia", age: 18}
	m2 := &man{name: "huyun", age: 29}
	m3 := &man{name: "troy", age: 25}

	visitors := []Visitor{m1, m2, m3}

	visitor = VisitorList(visitors)
	visitor.Visit(happyVisitor)

	fmt.Println("================")

	visitor = validationVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)

	fmt.Println("================")
	visitor = errorVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)

	fmt.Println("================")

	visitor = ageVisitor{visitor: visitor}
	visitor.Visit(happyVisitor)
}
