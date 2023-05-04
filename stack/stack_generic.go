package stack

import (
	"fmt"
	"log"
)

var storageManger []interface{}

type StorageType interface {
	int64 | int | float64 | float32 | string
}

func PrintListGeneric() {
	storageGenericInfo(storageManger)
	pushGeneric(90)
	pushGeneric("test")
	pushGeneric(22.5)
	pushGeneric(float32(22.589))

	print("\n")
	storageGenericInfo(storageManger)
	peekGeneric()
	isEmptyGeneric()
	print("\n")
	popGeneric()
	print("\n")
	popGeneric()
	print("\n")
	popGeneric()
	print("\n")
	popGeneric()
	isEmptyGeneric()
}

func storageGenericInfo(input interface{}) {
	fmt.Printf("current storage state: %v\n", input)
}

func pushGeneric[V StorageType](input V) []interface{} {
	storageManger = append(storageManger, input)
	log.Printf("storage generic [push: %v]: after push: %v\n", input, storageManger)

	return storageManger
}

func popGeneric() interface{} {
	var result interface{}
	var tempStorage []interface{}
	if len(storageManger) != 0 {
		result = peekGeneric()

		tempStorage = storageManger[:len(storageManger)-1]
	}
	storageManger = tempStorage
	log.Printf("storage [pop: %v]: after pop: %v\n", result, storageManger)

	return result
}

func peekGeneric() interface{} {
	n := storageManger[len(storageManger)-1 : len(storageManger)]
	log.Printf("storage peek: %v\n", n)

	return n[0]
}

func isEmptyGeneric() bool {
	rs := len(storageManger) == 0
	log.Printf("storage generic-isEmplty: %v\n", rs)

	return rs
}
