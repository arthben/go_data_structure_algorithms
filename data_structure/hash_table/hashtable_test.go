package hashtable

import (
	"fmt"
	"strconv"
	"testing"
)

func TestChain(t *testing.T) {
	ch := NewChain()
	ch.AddChain("baa", 1)
	ch.AddChain("aab", 2)
	ch.AddChain("aba", 3)
	ch.Print()

	fmt.Println("\nDo Remove key aab")
	ch.DeleteChain("aab")
	fmt.Println("\nTambah data lagi")
	ch.AddChain("aab", 4)
	ch.Print()

	v := ch.PeekChain("bac")
	fmt.Printf("v: %v\n", v)
}

func TestHashTable(t *testing.T) {
	ht := NewHashTable()
	ht.Push("ba", 1)
	ht.Push("ab", 2)
	ht.Push("b", 3)

	// hapus
	ht.Pop("b")

	ht.Show()

	// v := ht.Get("ab")
	// fmt.Printf("Value dari Key %v: %v\n", "ab", v)
}

func TestHashFunc(t *testing.T) {
	key := "catur"
	countIdx := 0
	for _, v := range key {
		strIdx := fmt.Sprintf("%v", v)
		intIdx, _ := strconv.Atoi(strIdx)
		countIdx += intIdx
	}

	countIdx = countIdx % 10
	fmt.Printf("intIdx: %v\n", countIdx)
}
