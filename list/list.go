package list

type node struct {
	value interface{}
	next  *node
}

// LinkedList is a collection of linked nodes
type LinkedList struct {
	head *node
}

// New creates a new list
func New() *LinkedList {
	return &LinkedList{}
}

// Size returns list length
func (list LinkedList) Size() int {
	var tmp = list.head
	var counter = 0

	for {
		if tmp != nil {
			tmp = tmp.next
			counter++
		} else {
			break
		}
	}

	return counter
}

// Add inserts an element into the list
func (list *LinkedList) Add(item interface{}) {
	var tmp = list.head

	var node = &node{
		value: item,
		next:  nil,
	}

	if tmp == nil {
		list.head = node
		return
	}

	for {
		if tmp.next != nil {
			tmp = tmp.next
		} else {
			break
		}
	}

	tmp.next = node
}

// Get returns an element by index
func (list *LinkedList) Get(index int) interface{} {
	var tmp = list.head
	var counter = 0

	if tmp == nil {
		return nil
	}

	if index == 0 && tmp != nil {
		return tmp.value
	}

	for {
		if tmp.next != nil {
			tmp = tmp.next
			counter++
			if counter == index {
				return tmp.value
			}
		} else {
			break
		}
	}

	return nil
}

// Pop removes first element and returns its value
func (list *LinkedList) Pop() interface{} {
	if list.head != nil {
		list.head = list.head.next
	}

	return list.head
}

// Del removes an element by index
func (list *LinkedList) Del(index int) {
	var tmp = list.head
	var target *node
	var counter = 0
	
	if index < 0 {
		return
	}

	if index == 0 {
		list.Pop()
		return
	}

	for {
		if tmp != nil {
			if counter == index-1 {
				target = tmp.next
				if target != nil {
					tmp.next = target.next
					return
				}
			}
		} else {
			break
		}

		tmp = tmp.next;
		counter++;
	}
}
