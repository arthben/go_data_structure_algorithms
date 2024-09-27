package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStack()
	fmt.Printf("Stack is empty : %v\n", st.isEmpty())

	st.Push(1)
	st.Push(2)
	fmt.Println("Before Pop")
	st.Print()

	st.Pop()
	fmt.Println("After Pop")
	st.Print()

	st.Count()
	fmt.Printf("st.Peek(): %v\n", st.Peek())
}
