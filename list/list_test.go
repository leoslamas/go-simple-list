package list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func setup() *LinkedList {
	myList := New()
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

	if value, ok := myList.Get(0).(int); ok {
		assert.Equal(t, 1, value)
	}

	if value, ok := myList.Get(1).(int); ok {
		assert.Equal(t, 10, value)
	}

	if value, ok := myList.Get(2).(int); ok {
		assert.Equal(t, 100, value)
	}

	if _, ok := myList.Get(3).(int); ok {
		t.FailNow()
	}
	
}

func TestLinkedList_Pop(t *testing.T) {
	myList := setup()

	if pop, ok := myList.Pop().(int); ok {
		assert.Equal(t, 1, pop)
		assert.Equal(t, 2, myList.Size)
	}

	if pop, ok := myList.Pop().(int); ok {
		assert.Equal(t, 10, pop)
		assert.Equal(t, 1, myList.Size)
	}

	if pop, ok := myList.Pop().(int); ok {
		assert.Equal(t, 100, pop)
		assert.Equal(t, 0, myList.Size)
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

func BenchmarkLinkedList_Add(b *testing.B) {
	myList := New()
	
	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	myList := New()
	
	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		myList.Get(i)
	}
}

func BenchmarkLinkedList_Size(b *testing.B) {
	myList := New()
	var x int64

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		x += int64(myList.Size())
	}
}

func BenchmarkLinkedList_Pop(b *testing.B) {
	myList := New()
	var x int64

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		if val, ok := myList.Pop().(int); ok {
			x += int64(val)
		}
	}
}

func BenchmarkLinkedList_Del(b *testing.B) {
	myList := New()

	for i := 0; i < b.N; i++ {
		myList.Add(i)
	}

	for i := 0; i < b.N; i++ {
		myList.Del(i)
	}
}