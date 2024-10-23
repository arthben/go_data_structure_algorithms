package heaps

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type Data struct {
	sum    int
	n1, n2 int
}
type MyHeaps []Data

func (h MyHeaps) Len() int            { return len(h) }
func (h MyHeaps) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h MyHeaps) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MyHeaps) Push(x interface{}) { *h = append(*h, x.(Data)) }
func (h *MyHeaps) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*
Find K Pairs with smallest Sums (leetcode #373)

You are given two integer arrays nums1 and nums2 sorted in non-decreasing order and an integer k.
Define a pair (u, v) which consists of one element from the first array and one element from the second array.
Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

Example 1:
Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

Example 2:
Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [[1,1],[1,1]]
Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
*/
func TestSmallestPairs(t *testing.T) {
	k := 3
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	fmt.Printf("nums1: %v, nums2: %v, k: %v\n", nums1, nums2, k)
	fmt.Printf("kSmallestPairs: %v\n", kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	sumMinHeap := &MyHeaps{}
	for j := 0; j < len(nums2) && j < k; j++ {
		*sumMinHeap = append(*sumMinHeap, Data{nums1[0] + nums2[j], 0, j})
	}
	fmt.Printf("sumMinHeap: %v\n", sumMinHeap)

	heap.Init(sumMinHeap)
	res := [][]int{}
	for len(res) < k && sumMinHeap.Len() > 0 {
		smallestSum := heap.Pop(sumMinHeap).(Data)
		res = append(res, []int{nums1[smallestSum.n1], nums2[smallestSum.n2]})
		if smallestSum.n1 < len(nums1)-1 {
			heap.Push(sumMinHeap, Data{nums1[smallestSum.n1+1] + nums2[smallestSum.n2],
				smallestSum.n1 + 1, smallestSum.n2})
		}
	}
	return res
}

// fungsi ini berjalan, tetapi tidak optimal saat data mencapai ribuan
func xkSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var temp [][]int

	for i := 0; i < len(nums1); i++ {

		j := len(nums2) - 1
		for j >= 0 {
			temp = append(temp, []int{nums1[i], nums2[j]})
			j--
		}
	}

	sort.SliceStable(temp, func(i, j int) bool {
		n1 := temp[i][0] + temp[i][1]
		n2 := temp[j][0] + temp[j][1]
		return n1 < n2
	})

	return temp[:k]
}

/*
Top K Frequent element (Leetcode #347)

Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]
*/
func TestTopKFrequent(t *testing.T) {
	k := 2
	nums := []int{1, 1, 1, 2, 2, 3}
	// k := 1
	// nums := []int{3, 0, 1, 0}
	// nums := []int{-1, -1}
	fmt.Printf("Input: %v, k: %v\n", nums, k)
	fmt.Printf("topKFrequent(): %v\n", topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	temp := make(map[int]int, 0)
	for _, num := range nums {
		temp[num]++
	}

	var keys []int
	for key := range temp {
		keys = append(keys, key)
	}

	if len(keys) == 1 {
		return keys
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return temp[keys[i]] > temp[keys[j]]
	})

	var final []int
	for _, x := range keys {
		if k == 0 {
			break
		}
		final = append(final, x)
		k--
	}

	return final
}

/*
Kth Largest Element in array (Leetcode #245)

Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Can you solve it without sorting?

Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
*/
func TestKthLargest(t *testing.T) {
	k := 2
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Printf("Input: %v, k: %v\n", nums, k)
	fmt.Printf("findKthLargest: %v\n", findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	maxHeap := &MaxHeap{}
	*maxHeap = nums
	heap.Init(maxHeap)

	var res int
	for k > 0 {
		res = heap.Pop(maxHeap).(int)
		k--
	}

	return res
}

func TestMinHeapByInit(t *testing.T) {
	// Time complexity : O(n)
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	fmt.Printf("Array Input : %v\n", array)

	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	*minHeap = array

	heap.Init(minHeap) // <-- bagian ini menggunakan method bawaan dari package container/heap
	fmt.Println("Heap By Init: ", *minHeap)
	processMinHeap(minHeap)
}

func TestMinHeapByPush(t *testing.T) {
	// Time complexity : O(logn)
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	fmt.Printf("Array Input : %v\n", array)

	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	for _, num := range array {
		heap.Push(minHeap, num)
	}
	fmt.Println("Heap By Push: ", *minHeap)
	processMinHeap(minHeap)
}

func processMinHeap(minHeap *MinHeap) {
	// apply Pop method: remove and return the minimum element
	elementToRemove := heap.Pop(minHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)
	fmt.Println("Heap after Pop: ", *minHeap)

	// apply Fix method
	// change number from 15 t0 50 at the index 2
	(*minHeap)[2] = 50
	fmt.Println("Modify min heap value at index 2: ", *minHeap)
	heap.Fix(minHeap, 2)
	fmt.Println("Fix min heap: ", *minHeap)

	// apply Remove method
	// remove number at index 4
	heap.Remove(minHeap, 4)
	fmt.Println("Heap after remove element in index 4: ", *minHeap)

}

func TestMaxHeapByInit(t *testing.T) {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	fmt.Printf("Array Input : %v\n", array)

	// initialize the MaxHeap that has implement the heap.Interface
	maxHeap := &MaxHeap{}
	*maxHeap = array
	heap.Init(maxHeap)
	fmt.Println("MaxHeap By Init: ", *maxHeap)
	processMaxHeap(maxHeap)
}

func TestMaxHeapByPush(t *testing.T) {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	fmt.Printf("Array Input : %v\n", array)

	// initialize the MaxHeap that has implement the heap.Interface
	maxHeap := &MaxHeap{}
	for _, num := range array {
		heap.Push(maxHeap, num)
	}
	fmt.Println("MaxHeap By Push: ", *maxHeap)
	processMaxHeap(maxHeap)
}

func processMaxHeap(maxHeap *MaxHeap) {
	// apply Pop method: remove and return the minimum element
	elementToRemove := heap.Pop(maxHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)
	fmt.Println("MaxHeap after pop element: ", *maxHeap)

	// apply Fix method
	// change number from 15 t0 50 at the index 2
	(*maxHeap)[2] = 50
	fmt.Println("Modify max heap value at index 2: ", *maxHeap)
	heap.Fix(maxHeap, 2)
	fmt.Println("Fix max heap: ", *maxHeap)

	// apply Remove method
	fmt.Println("Current max heap: ", *maxHeap)
	// remove number at index 4
	heap.Remove(maxHeap, 4)
	fmt.Println("Current max heap after remove element in index 4: ", *maxHeap)
}
