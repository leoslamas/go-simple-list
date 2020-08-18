package list

type node struct {
	value interface{}
	next  *node
}

type iterator struct {
	node *node
}

func (i *iterator) hasNext() bool {
	return i.node != nil && i.node.next != nil
}

func (i *iterator) next() *node {
	if i.node != nil {
		n := i.node.next
		i.node = i.node.next
		return n
	}

	return nil
}

func (list *LinkedList) newIter() *iterator {
	node := &node{value: nil, next:list.head}
	return &iterator{node: node}
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
 
// Iter returns a channel to iterate using `for ... range` using an iterator.
func (list *LinkedList) Iter() <-chan interface{} {
	var channel = make(chan interface{}, list.Size())
	defer close(channel)

	var iter = list.newIter()

	for{
		if iter.hasNext() {
			channel <- iter.next().value
		} else {
			break
		}
	}

	return channel

}

// Iter2 returns a channel to iterate using `for ... range`
// but it uses index, thus traversing the entire list for each element 
func (list *LinkedList) Iter2() <-chan interface{} {
	var channel = make(chan interface{}, list.Size())
	defer close(channel)

	for i:=0; i<list.Size(); i++ {
		channel <- list.Get(i)
	}

	return channel
}