package linkedlist

import "fmt"

type LinkedList interface {
	Print()
	AddLast(val int)
	AddFirst(val int)
	IndexOf(val int) int
	DeleteFirst()
	DeleteLast()
	Contains(val int) bool
}

type node struct {
	val  int
	next *node
}

func NewNode(val int) *node {
	return &node{val: val, next: nil}
}

type linkedList struct {
	head *node
	tail *node
}

func NewLinkedList() LinkedList {
	return &linkedList{
		head: nil,
		tail: nil,
	}
}

func (l *linkedList) AddLast(val int) {
	newNode := NewNode(val)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

}

func (l *linkedList) AddFirst(val int) {
	newNode := NewNode(val)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		prevHead := l.head
		l.head = newNode
		l.head.next = prevHead
	}

}

func (l *linkedList) IndexOf(val int) int {
	idx := 0
	current := l.head
	for current != nil {
		if current.val == val {
			return idx
		}
		current = current.next
		idx++
	}
	return -1
}

func (l *linkedList) Contains(val int) bool {
	return l.IndexOf(val) != -1
}

func (l *linkedList) DeleteFirst() {

	if l.head == nil && l.tail == nil {
		fmt.Println("Empty Linked List")
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	if l.head.next != nil {
		l.head = l.head.next
	}

}

func (l *linkedList) DeleteLast() {
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
		fmt.Printf("current.val: %v\n", current.val)
		current = current.next
	}
}
