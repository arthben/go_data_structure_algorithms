package queque

import "fmt"

/*
- add selalu dari tail
- remove selalu dari head
- operasi keluar masuk pada Queue adalah
FIFO ( First In First Out). Data selalu keluar dari head
*/

type Queue interface {
	isEmpty() bool
	Add(val int)
	Remove()
	Print()
	Peek()
}

type node struct {
	val  int
	next *node
}

func NewNode(val int) *node {
	return &node{val: val, next: nil}
}

type queData struct {
	head *node
	tail *node
}

func NewQueue() Queue {
	return &queData{
		head: nil,
		tail: nil,
	}
}

func (q *queData) isEmpty() bool {
	return q.head == nil
}

func (q *queData) Add(val int) {
	newData := NewNode(val)
	if q.head == nil {
		q.head = newData
		q.tail = newData
		return
	}

	q.tail.next = newData
	q.tail = newData
}

func (q *queData) Remove() {
	if q.head == nil {
		fmt.Println("Queue is Empty")
		return
	}

	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return
	}

	if q.head.next != nil {
		q.head = q.head.next
	}
}

func (q *queData) Print() {
	current := q.head
	for current != nil {
		fmt.Printf("current.val: %v\n", current.val)
		current = current.next
	}
}

func (q *queData) Peek() {
	fmt.Println(q.tail.val)
}
