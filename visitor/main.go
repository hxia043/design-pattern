package main

import "fmt"

type man struct {
	action string
}

type woman struct {
	action string
}

func (m *man) status() {
	if m.action == "happy" {
		fmt.Println("I have a toy")
	}

	if m.action == "sad" {
		fmt.Println("I lost my toy")
	}

	if m.action == "success" {
		fmt.Println("I have a great wife")
	}
}

func (w *woman) status() {
	if w.action == "happy" {
		fmt.Println("I have a lovely husband")
	}

	if w.action == "sad" {
		fmt.Println("I lost my lovely husband")
	}

	if w.action == "success" {
		fmt.Println("I have a lovely husband")
	}
}

func main() {
	m := man{
		action: "sad",
	}

	m.status()
}
