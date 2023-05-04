package main

import (
	"log"

	"github.com/yotzapon/data-stack/stack"
)

func main() {
	log.Println("--------------------------- Stack Data Structure")
	stack := stack.NewStackStructure()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	print("\n")
	stack.PrintStack()
	stack.Peek()
	stack.IsEmpty()
	print("\n")
	stack.Pop()
	print("\n")
	stack.Pop()
	print("\n")
	stack.Pop()
	print("\n")
	stack.Pop()
	stack.IsEmpty()
}
