package fastslowpointers

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

/*
Happy Number - Leetcode #202

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
Example 2:

Input: n = 2
Output: false
*/

func TestHappy(t *testing.T) {
	n := 4
	fmt.Printf("isHappy: %v\n", isHappy(n))
	fmt.Printf("isOtherHappy: %v\n", isOtherHappy(n))
}

func slicing(n int) int {
	sum := 0
	for n > 0 {
		t := n % 10
		sum += t * t
		n /= 10
	}
	return sum
}

func isOtherHappy(n int) bool {
	calc := make(map[int]bool, 0)
	for {
		n = slicing(n)
		if n == 1 || calc[n] {
			break
		}

		calc[n] = true
	}

	return n == 1
}

func isHappy(n int) bool {
	slow, fast := n, n
	slow = slicing(slow)
	fast = slicing(slicing(fast))

	for slow != fast {
		slow = slicing(slow)
		fast = slicing(slicing(fast))
	}
	return slow == 1
}
