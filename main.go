package main

import (
	"go-list/list"
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
	
	for val := range myList.Iter() {
		if value, ok := val.(int); ok {
			println("List Get: ", value)
		}
	}
}