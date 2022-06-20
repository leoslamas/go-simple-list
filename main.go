package main

import (
	"go-list/list"
)

func main() {
	var myList = initialize()
	println("List size: ", myList.Size())


	println("List size: ", myList.Size())

	printList(myList)

	value := myList.Pop()
	println("Pop: ", value)
	println("Size: ", myList.Size())

	myList.Add(1000)
	printList(myList)

	myList.Del(2)
	printList(myList)

	myList2 := initialize()

	filtered := list.Filter(myList2, func(x int) bool {
		return x % 2 == 0
	})
	print("Filter even: ")
	printList(filtered)

	mapped := list.Map(myList2, func(a int) int {
		return a + 1
	})
	print("Map +1: ")
	printList(mapped)

	reduced := list.Fold(myList2, 0, func(a, b int) int {
		return a + b
	})
	println("Reduce +1: ", reduced)
}

func initialize() *list.LinkedList[int] {
	myList := list.New[int]()

	myList.Add(1)
	myList.Add(10)
	myList.Add(100)

	return myList
}

func printList[T any](myList *list.LinkedList[T]) {
	println("Size: ", myList.Size())
	for val := range myList.Iter() {
		println("List Iter: ", val)
	}
}
