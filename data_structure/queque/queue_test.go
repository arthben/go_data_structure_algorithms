package queque

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Add(1)
	q.Add(2)
	q.Remove()
	fmt.Printf("Queue is Empty : %v\n", q.isEmpty())
	q.Print()

	q.Peek()
}
