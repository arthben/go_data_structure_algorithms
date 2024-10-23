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
	LinkededListHead() *Node
}

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val, Next: nil}
}

type linkedList struct {
	head *Node
	tail *Node
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
		l.tail.Next = newNode
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
		l.head.Next = prevHead
	}

}

func (l *linkedList) IndexOf(val int) int {
	idx := 0
	current := l.head
	for current != nil {
		if current.Val == val {
			return idx
		}
		current = current.Next
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

	if l.head.Next != nil {
		l.head = l.head.Next
	}

}

func (l *linkedList) DeleteLast() {
	current := l.head
	for current.Next.Next != nil {
		current = current.Next
	}
	l.tail = current
	l.tail.Next = nil
}

func (l *linkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("current.val: %v\n", current.Val)
		current = current.Next
	}
}

func (l *linkedList) LinkededListHead() *Node {
	return l.head
}
