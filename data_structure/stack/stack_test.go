package stack

import (
	"fmt"
	"testing"
)

func TestAja(t *testing.T) {
	st := NewStack()
	// fmt.Printf("st.IsEmpty(): %v\n", st.IsEmpty())

	// angka yang dicari adalah 4, 1, 2
	// jawaban 4 adalah -1
	// jawaban 1 adalah 3
	// jawaban 2 adalah -1
	cari := 2
	angka := []int{1, 3, 4, 2}
	for i := len(angka) - 1; i >= 0; i-- {
		for !st.IsEmpty() && st.Peek() <= angka[i] {
			st.Pop()
		}

		if angka[i] == cari {
			fmt.Printf("st.Peek(): %v\n", st.Peek())
			return
		}
		st.Push(angka[i])

	}

	st.Print()
}

func TestStack(t *testing.T) {
	st := NewStack()
	fmt.Printf("Stack is empty : %v\n", st.IsEmpty())

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
