package main

import (
	"go-list/list"
)

func main() {

	var myList = list.New[int]()
	println("List size: ", myList.Size())

	myList.Add(1)
	myList.Add(10)
	myList.Add(100)

	println("List size: ", myList.Size())

	printList(myList)

	value := myList.Pop()
	println("Pop: ", value)
	println("Size: ", myList.Size())

	myList.Add(1000)
	printList(myList)

	myList.Del(2)
	printList(myList)

	myList.Pop()
	printList(myList)

}

func printList[T any](myList *list.LinkedList[T]) {
	println("Size: ", myList.Size())

	for val := range myList.Iter() {
		println("List Iter: ", val)
	}

	for val := range myList.Iter2() {
		println("List Iter2: ", val)
	}
}
