package stack

import (
	"fmt"
	"log"
)

type StackManager interface {
	PrintStack()
	Pop() interface{}
	Push(input interface{})
	Peek() interface{}
	IsEmpty() bool
	CountStackElement() int
}

type implStack struct {
	storageManger []interface{}
}

func (i *implStack) PrintStack() {
	fmt.Printf("current storage state: %v\n", i)
}

func (i *implStack) Pop() interface{} {
	var result interface{}
	var tempStorage []interface{}
	if len(i.storageManger) != 0 {
		result = i.Peek()
		tempStorage = i.storageManger[:len(i.storageManger)-1]
	}
	i.storageManger = tempStorage
	log.Printf("storage [pop: %v]: after pop: %v\n", result, i.storageManger)

	return result
}

func (i *implStack) Push(input interface{}) {
	i.storageManger = append(i.storageManger, input)
	log.Printf("storage generic [push: %v]: after push: %v\n", input, i.storageManger)

}

func (i *implStack) Peek() interface{} {
	n := i.storageManger[len(i.storageManger)-1 : len(i.storageManger)]
	log.Printf("storage peek: %v\n", n)

	return n[0]
}

func (i *implStack) IsEmpty() bool {
	rs := len(i.storageManger) == 0
	log.Printf("storage generic-isEmplty: %v\n", rs)

	return rs
}

func (i *implStack) CountStackElement() int {
	return len(i.storageManger)
}

func NewStackStructure() StackManager {
	return &implStack{}
}
