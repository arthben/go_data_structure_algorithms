package latihan

/*
Maximum Widht Ramp (Leetcode #962)

A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i] <= nums[j]. The width of such a ramp is j - i.
Given an integer array nums, return the maximum width of a ramp in nums. If there is no ramp in nums, return 0.

Example 1:
Input: nums = [6,0,8,2,1,5]
Output: 4
Explanation: The maximum width ramp is achieved at
(i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.

Example 2:
Input: nums = [9,8,1,0,1,9,4,0,4,1]
Output: 7
Explanation: The maximum width ramp is achieved at
(i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.

Monotonic Stack: We traverse the array and maintain a stack of indices where the values of nums are in decreasing order.
This ensures that when we later iterate over the array from right to left, we can quickly find an index i such that nums[i] <= nums[j]
*/
func maxWidthRamp(nums []int) int {
	stack := []int{}

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	width := 0
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if j-i > width {
				width = j - i
			}
		}
	}
	return width
}
