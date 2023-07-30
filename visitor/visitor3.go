package main

import "fmt"

type VisitorFunc func(man)

func happyVisitor(m man) {
	fmt.Println(m.name, "I have a good thing")
}

func sadVisitor(m man) {
	fmt.Println(m.name, "I have a bad thing")
}

func successVisitor(m man) {
	fmt.Println(m.name, "I finish a thing")
}

type man struct {
	name string
}

func (m man) accept(v VisitorFunc) {
	v(m)
}

func main() {
	m := man{"hxia"}
	m.accept(sadVisitor)
}
