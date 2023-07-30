package main

import "fmt"

type visitor interface {
	manStatus(*man)
	womanStatus(*woman)
}

type happyVisitor struct{}
type sadVisitor struct{}
type successVisitor struct{}

func (h *happyVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I have a toy")
}

func (s *sadVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I lost my toy")
}

func (s *successVisitor) manStatus(m *man) {
	fmt.Println(m.name, "I have a great wife")
}

func (h *happyVisitor) womanStatus(w *woman) {
	fmt.Println(w.name, "I have a lovely husband")
}

func (s *sadVisitor) womanStatus(w *woman) {
	fmt.Println(w.name, "I lost my lovely husband")
}

func (s *successVisitor) womanStatus(w *woman) {
	fmt.Println(w.name, "I have a lovely husband")
}

type man struct {
	name string
}

type woman struct {
	name string
}

func (m *man) accept(s visitor) {
	s.manStatus(m)
}

func (w *woman) accept(s visitor) {
	s.womanStatus(w)
}

func main() {
	m := &man{"hxia"}
	s := &sadVisitor{}
	m.accept(s)
}
