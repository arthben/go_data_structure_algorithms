/*
Implementasi linked list untuk menghindari collision
*/

package hashtable

import "fmt"

type dataHash struct {
	val int
	key string
}

type node struct {
	val  *dataHash
	next *node
}

func newNode(val *dataHash) *node {
	return &node{val: val, next: nil}
}

type Chain interface {
	AddChain(key string, val int)
	DeleteChain(key string)
	PeekChain(key string) int
	Print()
}

type linkedList struct {
	head *node
	tail *node
}

func NewChain() Chain {
	return &linkedList{
		head: nil,
		tail: nil,
	}
}

func (l *linkedList) AddChain(key string, val int) {
	// add to the last
	newNode := newNode(&dataHash{
		val: val,
		key: key,
	})

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

}

func (l *linkedList) PeekChain(key string) int {
	current := l.head
	for current != nil {
		if current.val.key == key {
			return current.val.val
		}
		current = current.next
	}

	return -1
}

func (l *linkedList) DeleteChain(key string) {
	// check start from the head
	if l.head.val.key == key {
		if l.head.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
		}
		return
	}

	prev := l.head
	current := l.head.next
	for current != nil {
		if current.val.key == key {
			prev.next = current.next
		}
		prev = current
		current = current.next
	}
}

func (l *linkedList) XDeleteChain(key string) {
	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	l.tail = current
	l.tail.next = nil
}

func (l *linkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf(" --> Key: %v, Value: %v\n", current.val.key, current.val.val)
		current = current.next
	}
}
