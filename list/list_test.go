package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *LinkedList[int] {
	myList := New[int]()
	myList.Add(1)
	myList.Add(10)
	myList.Add(100)

	return myList
}

func TestLinkedList_Size(t *testing.T) {
	myList := setup()

	assert.Equal(t, 3, myList.Size())

	myList.Add(1000)
	myList.Add(10000)

	assert.Equal(t, 5, myList.Size())
}

func TestLinkedList_Add(t *testing.T) {
	myList := setup()

	myList.Add(90)
	assert.Equal(t, 90, myList.Get(3))
	assert.Equal(t, 4, myList.Size())
}

func TestLinkedList_Get(t *testing.T) {
	myList := setup()

	var value int

	value = myList.Get(0)
	assert.Equal(t, 1, value)

	value = myList.Get(1)
	assert.Equal(t, 10, value)

	value = myList.Get(2)
	assert.Equal(t, 100, value)

	if val := myList.Get(3); !isZero(val) {
		t.FailNow()
	}
}

func TestLinkedList_Pop(t *testing.T) {
	myList := setup()

	var pop int

	pop = myList.Pop()
	assert.Equal(t, 1, pop)
	assert.Equal(t, 2, myList.Size())

	pop = myList.Pop()
	assert.Equal(t, 10, pop)
	assert.Equal(t, 1, myList.Size())

	pop = myList.Pop()
	assert.Equal(t, 100, pop)
	assert.Equal(t, 0, myList.Size())

	if pop := myList.Pop(); !isZero(pop) {
		t.FailNow()
	}
}

func TestLinkedList_Del(t *testing.T) {
	myList := setup()

	assert.Equal(t, 3, myList.Size())

	myList.Del(-1)
	assert.Equal(t, 3, myList.Size())

	myList.Del(3)
	assert.Equal(t, 3, myList.Size())

	myList.Del(2)
	assert.Equal(t, 2, myList.Size())

	myList.Del(1)
	assert.Equal(t, 1, myList.Size())

	myList.Del(0)
	assert.Equal(t, 0, myList.Size())
}

func TestLinkedList_Iter(t *testing.T) {
	myList := setup()
	var counter = 0
	var x = [3]int{1, 10, 100}

	for i := range myList.Iter() {
		assert.Equal(t, x[counter], i)
		counter++
	}

	emptyList := New[int]()
	for i := range emptyList.Iter() {
		t.Error("Should be unreachable ", i)
	}
}

func TestLinkedList_Iter2(t *testing.T) {
	myList := setup()
	var counter = 0
	var x = [3]int{1, 10, 100}

	for i := range myList.Iter2() {
		assert.Equal(t, x[counter], i)
		counter++
	}

	emptyList := New[int]()
	for i := range emptyList.Iter2() {
		t.Error("Should be unreachable ", i)
	}
}

func TestLinkedList_Filter(t *testing.T) {
	myList := setup()

	filteredList1 := Filter(myList, func(val int) bool {
		return val < 100
	})
	
	list1 := []int{}
	for i := range filteredList1.Iter() {
		list1 = append(list1, i)
	}

	assert.ElementsMatch(t, []int{1,10}, list1)

	filteredList2 := Filter(myList, func(val int) bool {
		return val > 10
	})
	
	list2 := []int{}
	for i := range filteredList2.Iter() {
		list2 = append(list2, i)
	}

	assert.ElementsMatch(t, []int{100}, list2)
}

func TestLinkedList_Map(t *testing.T) {
	myList := setup()

	mappedList := Map(myList, func(val int) int {
		return val+1
	})

	list := []int{}
	for i := range mappedList.Iter() {
		list = append(list, i)
	}

	assert.ElementsMatch(t, []int{2,11,101}, list)
}

func TestLinkedList_Fold(t *testing.T) {
	myList := setup()

	reduced := Fold(myList, 0, func(acc, val int) int {
		return acc + val
	})

	assert.Equal(t, 111, reduced)

	reduced = Fold(myList, 1, func(acc, val int) int {
		return acc * val
	})

	assert.Equal(t, 1000, reduced)
}

func BenchmarkLinkedList_Add(b *testing.B) {
	myList := New[int]()

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}
}

func BenchmarkLinkedList_Add_Get(b *testing.B) {
	myList := New[int]()

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		myList.Get(i)
	}
}

func BenchmarkLinkedList_Add_Size(b *testing.B) {
	myList := New[int]()
	var x int64

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		x += int64(myList.Size())
	}
}

func BenchmarkLinkedList_Add_Pop(b *testing.B) {
	myList := New[int]()
	var x int64

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		val := myList.Pop()
		x += int64(val)
	}
}

func BenchmarkLinkedList_Add_Del(b *testing.B) {
	myList := New[int]()

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		myList.Del(i)
	}
}

func BenchmarkLinkedList_Iter(b *testing.B) {
	myList := New[int]()
	var x int64

	for i := 0; i < 10000; i++ {
		myList.Add(i)
	}

	for j := 0; j < b.N; j++ {
		for i := range myList.Iter() {
			x += int64(i)
		}
	}
}

func BenchmarkLinkedList_Iter2(b *testing.B) {
	myList := New[int]()
	var x int64

	for i := 0; i < 10000; i++ {
		myList.Add(i)
	}

	for j := 0; j < b.N; j++ {
		for i := range myList.Iter2() {
			x += int64(i)
		}
	}
}
