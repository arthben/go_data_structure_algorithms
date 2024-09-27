package threesum

import (
	"fmt"
	"testing"
)

func TestSolutionOne(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}
	nums := []int{-2, 0, 1, 1, 2}
	target := 0

	ths := NewThreeSum(nums, target)
	fmt.Printf("ths.SolutionOne(): %v\n", ths.SolutionOne())
}

func TestSolutionTwo(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0, 0}
	nums := []int{-2, 0, 1, 1, 2}
	target := 0

	ths := NewThreeSum(nums, target)
	fmt.Printf("ths.SolutionTwo(): %v\n", ths.SolutionTwo())
}
