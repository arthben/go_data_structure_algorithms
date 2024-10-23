package monotonicstack

import (
	"fmt"
	"slices"
	"testing"

	"github.com/arthben/go_data_structure_algorithms/data_structure/stack"
)

func TestIncrease(t *testing.T) {
	IncreasingStack()
}

func TestDecrease(t *testing.T) {
	DecreaseStack()
}

/*
Next greater element I (Leetcode #496)

The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.
You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.
For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

Example 1:
Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.

Example 2:
Input: nums1 = [2,4], nums2 = [1,2,3,4]
Output: [3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
- 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
*/
func TestNextGreater(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	// nums1 := []int{2, 4}
	// nums2 := []int{1, 2, 3, 4}
	// jawaban : [3, -1]
	res := nextGreaterElement(nums1, nums2)
	fmt.Printf("nextGreaterElement: %v\n", res)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	output := []int{}
	for _, num := range nums1 {
		// r := monotonicDecrease(nums2, num)
		r := nextGenerator(nums2, num)
		output = append(output, r)
		fmt.Printf("num: %v, nextGreater: %v\n", num, r)
	}

	return output
}

// implement with stack
func monotonicDecrease(nums []int, nextGreater int) int {
	st := stack.NewStack()
	for i := len(nums) - 1; i >= 0; i-- {
		for !st.IsEmpty() && st.Peek() <= nums[i] {
			st.Pop()
		}

		if nums[i] == nextGreater {
			return st.Peek()
		}

		st.Push(nums[i])
	}

	return st.Peek()
}

// implement using array as stack
func nextGenerator(nums []int, nextGreater int) int {
	stack := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		topStack := len(stack)
		if topStack > 0 {
			for j := topStack - 1; j >= 0; j-- {
				if stack[j] <= nums[i] {
					stack = slices.Delete(stack, j, j+1)
				}
			}
		}

		if nums[i] == nextGreater {
			if len(stack) == 0 {
				return -1
			} else {
				return stack[len(stack)-1]
			}
		}

		stack = append(stack, nums[i])
	}

	return -1
}

/*
Daily Temperatures ( Leetcode #739)

Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the ith day
to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]
*/
func TestDaily(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("temperatures: %v\n", temperatures)
	fmt.Printf("dailyTemperatures(): %v\n", dailyTemperatures(temperatures))
	fmt.Println("correct answer : [1,1,4,2,1,1,0,0]")
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	final := make([]int, len(temperatures), len(temperatures))
	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			final[j] = i - j
		}

		stack = append(stack, i)
	}

	return final
}

/*
Largest Rectangle in Histogram (Leetcode #84)
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
return the area of the largest rectangle in the histogram.

Example:
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

Example
Input: heights = [2,4]
Output: 4
*/
func TestLargest(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	// heights := []int{4, 2}
	// heights := []int{0, 9}
	fmt.Printf("Input: %v\n", heights)
	fmt.Printf("largestRectangleArea(): %v\n", largestRectangleArea(heights))
	// fmt.Println("correct answer : 10")
}

func largestRectangleArea(heights []int) int {
	type data struct {
		index  int
		height int
	}
	stack := make([]data, 0)
	maxArea := 0

	for i, val := range heights {
		prevIndex := i
		for len(stack) > 0 && stack[len(stack)-1].height > val {
			// calculate area
			prev := stack[len(stack)-1]
			widht := i - prev.index
			maxArea = max(maxArea, prev.height*widht)
			prevIndex = prev.index

			// if previous height at top of stack is less than current, pop it from stack
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, data{index: prevIndex, height: val})
	}

	for _, value := range stack {
		area := value.height * (len(heights) - value.index)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
