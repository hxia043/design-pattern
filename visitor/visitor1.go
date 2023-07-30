package main

import "fmt"

type visitor interface {
	manStatus()
	womanStatus()
}

type happyVisitor struct{}
type sadVisitor struct{}
type successVisitor struct{}

func (h *happyVisitor) manStatus() {
	fmt.Println("I have a toy")
}

func (s *sadVisitor) manStatus() {
	fmt.Println("I lost my toy")
}

func (s *successVisitor) manStatus() {
	fmt.Println("I have a great wife")
}

func (h *happyVisitor) womanStatus() {
	fmt.Println("I have a lovely husband")
}

func (s *sadVisitor) womanStatus() {
	fmt.Println("I lost my lovely husband")
}

func (s *successVisitor) womanStatus() {
	fmt.Println("I have a lovely husband")
}

type man struct {
}

type woman struct {
}

func (m *man) accept(s visitor) {
	s.manStatus()
}

func (w *woman) accept(s visitor) {
	s.womanStatus()
}

func main() {
	m := &man{}
	s := &sadVisitor{}
	m.accept(s)
}
