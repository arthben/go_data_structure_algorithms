package hashtable

import (
	"fmt"
	"strconv"
)

/*
untuk menghindari collision, menggunakan open chaining
open chaining diimplementasikan dengan linked list
*/

type HashTable interface {
	Push(key string, data int)
	Get(key string) int
	Pop(key string)
	Show()
}

type hTable struct {
	table []Chain
}

func NewHashTable() HashTable {
	var tables []Chain
	hashTableCapacity := 10

	for i := 0; i <= hashTableCapacity-1; i++ {
		tables = append(tables, NewChain())
	}
	return &hTable{
		table: tables,
	}
}

func hashing(key string) int {
	countIdx := 0
	for _, v := range key {
		strIdx := fmt.Sprintf("%v", v)
		intIdx, _ := strconv.Atoi(strIdx)
		countIdx += intIdx
	}

	return countIdx % 10
}

func (h *hTable) Push(key string, data int) {
	tblIndex := hashing(key)
	h.table[tblIndex].AddChain(key, data)
}

func (h *hTable) Get(key string) int {
	tblIndex := hashing(key)
	return h.table[tblIndex].PeekChain(key)
}

func (h *hTable) Pop(key string) {
	tblIndex := hashing(key)
	h.table[tblIndex].DeleteChain(key)
}

func (h *hTable) Show() {
	for i := 0; i <= len(h.table)-1; i++ {
		fmt.Printf("Child on Index : %v\n", i)
		h.table[i].Print()
	}
}
