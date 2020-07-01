package main

import (
	"go-tests/list"
)

func main() {
	
	var myList = list.New()
	println("List size: ", myList.Size())

	myList.Add(1)
	myList.Add(10)
	myList.Add(100)
	
	println("List size: ", myList.Size())

	printList(myList)

	if value, ok := myList.Pop().(int); ok {
		println("Pop: ", value)
		println("Size: ", myList.Size())
	}

	myList.Add(1000)
	printList(myList)

	myList.Del(2)
	printList(myList)

	myList.Pop()
	printList(myList)
}

func printList(myList *list.LinkedList) {
	println("Size: ", myList.Size())
	
	if value, ok := myList.Get(0).(int); ok {
		println("List get(0): ", value)
	}

	if value, ok := myList.Get(1).(int); ok {
		println("List get(1): ", value)
	}

	if value, ok := myList.Get(2).(int); ok {
		println("List get(2): ", value)
	}

	if value, ok := myList.Get(3).(int); ok {
		println("List get(3): ", value)
	}
}