package array

import (
	"errors"
	"fmt"
)

type Array struct {
	Data     []int
	capacity int
	count    int
}

func NewArray(capacity int) *Array {
	return &Array{
		Data:     make([]int, capacity),
		count:    0,
		capacity: capacity,
	}
}

func (arr *Array) Push(data int) {
	// when the array capacity is full, array will copy
	// existing to new array with bigger capacity
	// for example, double the array capacity
	if arr.count == arr.capacity {
		newData := make([]int, arr.capacity*2)

		// reset the count & capacity
		arr.count = 0
		arr.capacity *= 2

		// time complexity is O(n)
		for range arr.Data {
			newData[arr.count] = arr.Data[arr.count]
			arr.count++
		}

		arr.Data = newData
	}

	// time complexity : O(1) -> constant, because use index
	arr.Data[arr.count] = data
	arr.count++
}

func (arr *Array) Pop(index int) error {
	if err := arr.validate(index); err != nil {
		return err
	}

	// ex : [10, 20, 30, 40, 50]
	// if index 1 is removed / pop, then array will become like this
	// [10, 30, 40, 50]
	// it means 30, 40, 50 will be shifted to the left
	for i := index; i < arr.count; i++ {
		if (i + 1) >= arr.count {
			arr.Data[i] = 0
			continue
		}
		arr.Data[i] = arr.Data[i+1]
	}

	arr.count--
	return nil
}

func (arr *Array) IndexOf(val int) int {
	// return index number of the val
	for i := 0; i < arr.count; i++ {
		if arr.Data[i] == val {
			return i
		}
	}

	return -1
}

func (arr *Array) Peek(index int) (int, error) {
	if err := arr.validate(index); err != nil {
		return -1, err
	}

	// time complexity : O(1) -> constant, because use index
	return arr.Data[index], nil
}

func (arr *Array) Print() {
	fmt.Printf("Array Data: %v\n", arr.Data)
}

func (arr *Array) validate(index int) error {
	if (index < 0) || (index > arr.count) {
		return errors.New("Array index is out of capacity")
	}

	return nil
}
