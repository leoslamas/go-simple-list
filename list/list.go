package list

// workarounds for zero values of non nil types
func zero[T any]() T {
	return *new(T)
}

func isZero[T comparable](v T) bool {
	return v == *new(T)
}

// LinkedList is a collection of linked nodes
type LinkedList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

type iterator[T any] struct {
	node *node[T]
}

func (i *iterator[T]) hasNext() bool {
	return i.node != nil && i.node.next != nil
}

func (i *iterator[T]) next() T {
	if i.node != nil {
		n := i.node.next
		i.node = i.node.next
		return n.value
	}

	return zero[T]()
}

func (list *LinkedList[T]) newIter() *iterator[T] {
	node := &node[T]{value: zero[T](), next: list.head}
	return &iterator[T]{node: node}
}

// New creates a new list
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Size returns list length
func (list LinkedList[T]) Size() int {
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
func (list *LinkedList[T]) Add(item T) {
	var tmp = list.head

	var node = &node[T]{
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
func (list *LinkedList[T]) Get(index int) T {
	var tmp = list.head
	var counter = 0

	if tmp == nil {
		return zero[T]()
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

	return zero[T]()
}

// Pop removes first element and returns its value
func (list *LinkedList[T]) Pop() T {
	var value = zero[T]()

	if list.head != nil {
		value = list.head.value
		list.head = list.head.next
	}

	return value
}

// Del removes an element by index
func (list *LinkedList[T]) Del(index int) {
	var tmp = list.head
	var target *node[T]
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

		tmp = tmp.next
		counter++
	}
}

// Iter returns a channel to iterate using `for ... range` using an iterator.
func (list *LinkedList[T]) Iter() <-chan T {
	var channel = make(chan T, list.Size())
	defer close(channel)

	var iter = list.newIter()

	for {
		if iter.hasNext() {
			channel <- iter.next()
		} else {
			break
		}
	}

	return channel

}

// Iter2 returns a channel to iterate using `for ... range`
// but it uses index, thus traversing the entire list for each element
func (list *LinkedList[T]) Iter2() <-chan T {
	var channel = make(chan T, list.Size())
	defer close(channel)

	for i := 0; i < list.Size(); i++ {
		channel <- list.Get(i)
	}

	return channel
}

// Filter filters the list by values for which 'pred' returns true
func Filter[T any](list *LinkedList[T], pred func(val T) bool) *LinkedList[T] {
	var newList = New[T]()

	var iter = list.newIter()

	for {
		if iter.hasNext() {
			nextVal := iter.next()
			if pred(nextVal) {
				newList.Add(nextVal)
			}
		} else {
			break
		}
	}

	return newList
}

// Map returns a new list applying 'f' to each value of 'list'
func Map[T any, U any](list *LinkedList[T], f func(val T) U) *LinkedList[U] {
	var newList = New[U]()

	var iter = list.newIter()

	for {
		if iter.hasNext() {
			nextVal := iter.next()
			newList.Add(f(nextVal))
		} else {
			break
		}
	}

	return newList
}

// Fold accumulates value starting with the first element and applying 'f' 
// from left to right to current accumulator value and each element.
func Fold[T any](list *LinkedList[T], init T, f func(acc, val T) T) T {
	var iter = list.newIter()

	var reduced = init

	for {
		if iter.hasNext() {
			nextVal := iter.next()
			reduced = f(reduced, nextVal)
		} else {
			break
		}
	}

	return reduced
}