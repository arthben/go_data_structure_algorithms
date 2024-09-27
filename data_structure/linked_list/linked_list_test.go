package linkedlist

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {
	lnk := NewLinkedList()
	lnk.AddLast(3)
	lnk.AddLast(4)
	lnk.AddLast(5)
	lnk.AddFirst(1)
	lnk.Print()

	lnk.DeleteFirst() // hapus 1
	fmt.Println("\nAfter Delete First")
	lnk.Print()

	lnk.DeleteLast() // hapus 5
	fmt.Println("\nAfter Delete Last")
	lnk.Print()

	lnk.AddLast(10)
	fmt.Println("\nAfter Add Last")
	lnk.Print()

	lnk.AddFirst(2)
	fmt.Println("\nAfter Add First")
	lnk.Print()

	idx := lnk.Contains(40)
	fmt.Printf("idx: %v\n", idx)
}
