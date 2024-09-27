package prefixsum

import (
	"fmt"
	"testing"
)

func TestAja(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("SEBELUM data: %v\n", data)
	for i := 1; i < len(data); i++ {
		data[i] += data[i-1]
	}
	fmt.Printf("SETELAH data: %v\n", data)
}

/*
Range Sum Query - Immutable (LeetCode #303)
Given an integer array nums, handle multiple queries of the following type:
1. Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

contoh :
data = [-2, 0, 3, -5, 2, -1]
left = 0, right = 2

maka = data[0] + data[1] + data[2]

	=   -2  +    0    +    3
	= 1
*/
func TestSumQuerySolutionOne(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	left, right := 1, 3

	count := 0
	for x := 0; x < len(data); x++ {
		if x >= left && x <= right {
			count += data[x]
		}
	}
	fmt.Printf("Final Answer: %v\n", count)
}

func TestSumQuerySolutionTwo(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	left, right := 1, 3

	pf := NewPrefixSum(data)
	sums := pf.Proses()

	final := sums[right] - sums[left-1]
	fmt.Printf("Final Answer: %v\n", final)
}

/*
Sub array Sum Equals K (LeetCode #560)
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.

contoh:
nums = [1,1,1], k = 2, output = 2
nums = [1,2,3], k = 3, output = 2
*/
func TestSubArraySumSolutionOne(t *testing.T) {
	k := 3
	nums := []int{1, 2, 3}
	count := 0
	steps := 0
	for i := range nums {
		steps++
		count += nums[i]
		if count == k {
			break
		}
	}

	fmt.Printf("Final Answer: %v\n", steps)
}

func TestSubArraySumSolutionTwo(t *testing.T) {
	k := 3
	nums := []int{1, 2, 3}
	pf := NewPrefixSum(nums)
	sums := pf.Proses()
	count := 0

	for i := range sums {
		fmt.Printf("i: %v\n", i)
		if sums[i] == k {
			count = i + 1
			break
		}
	}
	fmt.Printf("Final Answer: %v\n", count)
}
