package twosum

import (
	"fmt"
	"testing"
)

func TestSolutionOne(t *testing.T) {
	nums := []int{2, 7, -3, 11, 15}
	target := -1
	tsw := NewTwoSum(nums, target)
	fmt.Printf("tsw.SolutionOne(): %v\n", tsw.SolutionOne())
}

func TestSolutionTwo(t *testing.T) {
	nums := []int{2, 7, -3, 11, 15}
	target := -1
	tsw := NewTwoSum(nums, target)
	fmt.Printf("tsw.SolutionTwo(): %v\n", tsw.SolutionTwo())
}

func TestSolutionThree(t *testing.T) {
	nums := []int{2, 7, -3, 11, 15}
	target := -1
	tsw := NewTwoSum(nums, target)
	fmt.Printf("tsw.SolutionThree(): %v\n", tsw.SolutionThree())
}
