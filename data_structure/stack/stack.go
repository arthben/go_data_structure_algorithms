package stack

import "fmt"

/*
operasi keluar masuk data pada stack adalah
LIFO ( Last In First Out). Data selalu masuk dan keluar dari top
*/

type Stack interface {
	IsEmpty() bool
	Peek() int
	Push(data int)
	Pop()
	Print()
	Count()
}

type node struct {
	val  int
	next *node
}

func NewNode(val int) *node {
	return &node{val: val, next: nil}
}

type stackData struct {
	top   *node
	count int
}

func NewStack() Stack {
	return &stackData{top: nil, count: 0}
}

func (s *stackData) IsEmpty() bool {
	return s.top == nil
}

func (s *stackData) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.val
}

func (s *stackData) Push(data int) {
	newData := NewNode(data)

	if s.top == nil {
		s.top = newData
		s.count++
		return
	}

	newData.next = s.top
	s.top = newData
	s.count++
}

func (s *stackData) Pop() {
	if s.top != nil {
		s.top = s.top.next
		s.count--
	}
}

func (s *stackData) Print() {
	current := s.top
	for current != nil {
		fmt.Printf("current.val: %v\n", current.val)
		current = current.next
	}
}

func (s *stackData) Count() {
	fmt.Printf("Data on stack is : %v\n", s.count)
}
