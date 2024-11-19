package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestSortedToBST(t *testing.T) {

	array := []int{-5, 1, 5, 8, 12, 16, 19, 22, 66}
	bst := NewBinarySearchTree()
	res := bst.SortedArrayToBST(array)
	fmt.Printf("Sorted Array To BST: %v\n", bst.PrintBST(res))

	rev := bst.BSTToSortedArray(res)
	fmt.Printf("Binary Search Tree to Sorted Array: %v\n", rev)
}

/*
Search In Rotated Search Array (Leetcode #33)

There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1
*/
func TestSearch(t *testing.T) {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 0
	// jawaban := 4

	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 3
	// jawaban := -1

	// nums := []int{1, 3}
	// target := 1
	// jawaban := 0

	// nums := []int{1}
	// target := 0
	// jawaban := -1

	// nums := []int{1, 3, 5}
	// target := 1
	// jawaban := 0

	// nums := []int{1}
	// target := 1
	// jawaban := 0

	nums := []int{3, 5, 1}
	target := 3
	jawaban := 0

	fmt.Printf("nums: %v, target: %v\n", nums, target)
	fmt.Printf("hasil search(): %v\n", search(nums, target) == jawaban)
}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

/*
Find Minimum in Sorted Array ( Leetcode #153 )

Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
For example, the array nums = [0,1,2,4,5,6,7] might become:
-- [4,5,6,7,0,1,2] if it was rotated 4 times.
-- [0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Example 1:
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
*/
func TestMin(t *testing.T) {
	scenarios := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "A", nums: []int{3, 4, 5, 1, 2}, expected: 1},
		{name: "B", nums: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},
		{name: "C", nums: []int{11, 13, 15, 17}, expected: 11},
		{name: "D", nums: []int{2, 1}, expected: 1},
		{name: "E", nums: []int{1}, expected: 1},
		{name: "F", nums: []int{2, 3, 4, 5, 1}, expected: 1},
		{name: "G", nums: []int{5, 1, 2, 3, 4}, expected: 1},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {

			ans := findMin(scenario.nums)
			fmt.Printf("Input: %v, Expected: %v\n", scenario.nums, scenario.expected)

			if ans != scenario.expected {
				t.Errorf("Skenario: %v hasilnya: %v\n", scenario.name, ans == scenario.expected)
				return
			} else {
				fmt.Printf("Skenario: %v hasilnya: %v\n", scenario.name, ans == scenario.expected)

			}

		})
	}
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

/*
Search a 2D matrix II ( Leetcode #240 )

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Example 1:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Example 2:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false
*/
func TestMatrix(t *testing.T) {
	scenarios := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name:     "A",
			matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target:   5,
			expected: true,
		},
		// {
		// 	name:     "B",
		// 	matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
		// 	target:   20,
		// 	expected: false,
		// },
		// {
		// 	name:     "C",
		// 	matrix:   [][]int{{1, 1}},
		// 	target:   1,
		// 	expected: true,
		// },
		// {
		// 	name:     "D",
		// 	matrix:   [][]int{{-5}},
		// 	target:   -5,
		// 	expected: true,
		// },
		// {
		// 	name:     "E",
		// 	matrix:   [][]int{{-1, 3}},
		// 	target:   3,
		// 	expected: true,
		// },
		// {
		// 	name:     "F",
		// 	matrix:   [][]int{{5}, {6}},
		// 	target:   6,
		// 	expected: true,
		// },
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ans := searchMatrix(scenario.matrix, scenario.target)
			if ans == scenario.expected {
				fmt.Printf("Scenario: %v, result: %v\n", scenario.name, scenario.expected)
			} else {
				t.Errorf("Scenario: %v GAGAL!!", scenario.name)
				return
			}
		})
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for row, col := 0, n-1; row < m && col >= 0; {
		fmt.Printf("row: %v\n", row)
		fmt.Printf("col: %v\n", col)

		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col-- // the entire last col is greater than target
		} else {
			row++ // the entire row is smaller than target
		}
	}

	return false
}

func XsearchMatrix(matrix [][]int, target int) bool {

	for _, detail := range matrix {
		left, right := 0, len(detail)-1
		for left < right {
			mid := (left + right) / 2

			if detail[mid] == target {
				return true
			}

			if detail[mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if detail[left] == target {
			return true
		}
	}
	return false

}

func TestSederhana(t *testing.T) {
	// nums := []int{0, 1, 2, 3, 4, 5, 6}
	nums := []int{0, 1, 2}
	target := 0

	n := len(nums)
	left, right := 0, n-1
	foundIndex := 0
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			foundIndex = mid
			break
		}

		// isi nums adalah integer berurutan dari angka kecil ke angka besar
		// disebelah kiri, selalu angka terkecil dan kanan adalah angka terbesar
		// jika nums ditengan lebih dari target, berarti yang dicari adalah sebelah kiri ( angka lebih kecil )
		if nums[mid] > target {
			right = mid // left tetap berisi 0
		} else {
			left = mid + 1 // right tetap berisi n-1
		}
	}

	if nums[foundIndex] == target {
		fmt.Println("Ketemu")
	} else {
		fmt.Println("TIDAK KETEMU!!")
	}
}

func TestRuwet(t *testing.T) {
	nums := []int{4, 5, 6, 0, 1, 2}
	target := 10

	n := len(nums)
	left, right := 0, n-1
	foundIndex := 0
	for left < right {
		mid := (left + right) / 2
		fmt.Printf("left: %v, right: %v, mid: %v, nums[mid]: %v\n", left, right, mid, nums[mid])

		if nums[mid] == target {
			foundIndex = mid
			break
		}

		if nums[left] > nums[mid] {
			if nums[mid] > target && target < nums[right] {
				left = mid
			} else {
				right = mid + 1
			}

		} else {

			if nums[mid] > target && target < nums[right] {
				left = mid
			} else {
				right = mid + 1
			}
		}

	}

	if nums[foundIndex] == target {
		fmt.Println("Ketemu")
	} else {
		fmt.Println("TIDAK KETEMU!!")
	}
}
