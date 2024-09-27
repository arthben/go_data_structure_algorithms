package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	capacity := 2
	arr := NewArray(capacity)

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	// arr.Push(4)
	// arr.Push(5)

	arr.Pop(1)

	// find the index number of val 1
	idxOf := arr.IndexOf(1)
	fmt.Printf("idxOf: %v\n", idxOf)

	// // print all the array value
	arr.Print()
}
