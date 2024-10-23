package monotonicstack

import (
	"fmt"
	"strconv"
)

// Increasing monotonic stack
// value paling bawah dari stack adalah nilai terkecil
// value paling atas dari stack adalah nilai paling besar
func IncreasingStack() {
	arr := []int{1, 3, 10, 7, 5, 24, 4}
	stack := make([]int, 0)

	// monotonic increasing
	for i := range arr {
		// jika nilai plng atas stack lebih besar daripada nilai urutan arr,
		// maka keluarkan nilai tersebut dari stack
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	var final []string
	for i := range stack {
		final = append(final, strconv.Itoa(arr[i]))
	}

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("final: %v\n", final)
}

// Decreasing monotonic stack
// value paling bawah dari stack adalah nilai terbesar
// value paling atas dari stack adalah nilai terkecil
func DecreaseStack() {
	arr := []int{1, 3, 10, 7, 5, 24, 4}
	stack := make([]int, 0)

	// monotonic decreasing
	for i := range arr {
		// jika nilai plng atas stack lebih kecil daripada nilai urutan arr,
		// maka keluarkan nilai tersebut dari stack
		for len(stack) > 0 && stack[len(stack)-1] < arr[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, arr[i])
	}

	fmt.Printf("stack: %v\n", stack)
}
