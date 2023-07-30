package main

import "fmt"

type status interface {
	manStatus()
	womanStatus()
}

type happy struct{}
type sad struct{}
type success struct{}

func (h *happy) manStatus() {
	fmt.Println("I have a toy")
}

func (s *sad) manStatus() {
	fmt.Println("I lost my toy")
}

func (s *success) manStatus() {
	fmt.Println("I have a great wife")
}

func (h *happy) womanStatus() {
	fmt.Println("I have a lovely husband")
}

func (s *sad) womanStatus() {
	fmt.Println("I lost my lovely husband")
}

func (s *success) womanStatus() {
	fmt.Println("I have a lovely husband")
}

type man struct {
}

type woman struct {
}

func (m *man) accept(s status) {
	s.manStatus()
}

func (w *woman) accept(s status) {
	s.womanStatus()
}

func main() {
	m := &man{}
	s := &sad{}
	m.accept(s)
}
