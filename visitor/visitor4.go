package main

import "fmt"

type VisitorFunc func(person)

func happyVisitor(p person) {
	fmt.Println(p.getName(), "I have a good thing")
}

func sadVisitor(p person) {
	fmt.Println(p.getName(), "I have a bad thing")
}

func successVisitor(p person) {
	fmt.Println(p.getName(), "I finish a thing")
}

type person interface {
	getName() string
}

type man struct {
	name string
}

type woman struct {
	name string
}

func (m man) getName() string {
	return m.name
}

func (w woman) getName() string {
	return w.name
}

func (m *man) accept(v VisitorFunc) {
	v(m)
}

func (w *woman) accept(v VisitorFunc) {
	v(w)
}

func main() {
	m := &man{"hxia"}
	m.accept(sadVisitor)
}
