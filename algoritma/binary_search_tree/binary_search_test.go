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

func TestSearch(t *testing.T) {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 3
	// nums := []int{1, 3}
	// target := 1
	// nums := []int{1}
	// target := 0
	nums := []int{1, 3, 5}
	target := 1
	fmt.Printf("nums: %v, target: %v\n", nums, target)
	fmt.Printf("hasil search(): %v\n", search(nums, target))
}

func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	mid := (left + right) / 2
	fmt.Printf("mid: %v\n", mid)
	if n <= 1 && mid == 0 && nums[mid] != target {
		return -1
	} else if mid == 0 && nums[mid] == target {
		return 0
	}

	fmt.Printf("nums[mid]: %v\n", nums[mid+1])
	if nums[mid+1] >= target {
		fmt.Println("Masuk sini")
		left = mid + 1
	} else {
		right = mid - 1
	}

	for left <= right {
		fmt.Printf("nums[left]: %v\n", nums[left])
		if nums[left] == target {
			return left
		}

		left++
	}

	return -1
}
