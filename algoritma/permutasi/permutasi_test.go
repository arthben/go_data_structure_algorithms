package permutasi

import (
	"fmt"
	"sort"
	"testing"
)

func TestDua(t *testing.T) {
	// input
	nums := []int{1, 0, -1, 0, -2, 2}
	sort.Ints(nums)

	// final := make([][]int, 0)
	for i := range nums {
		left1 := i
		left2 := left1 + 1
		right1 := len(nums) - i
		right2 := right1 - 1
		fmt.Printf("l1: %v, l2: %v, r1: %v, r2: %v\n", left1, left2, right1, right2)

		// 	if nums[left1]+nums[left2]+nums[right1]+nums[right2] == 0 {
		// 		t := []int{nums[left1], nums[left2], nums[right1], nums[right2]}
		// 		final = append(final, t)
		// 	}
	}

	// fmt.Printf("final: %v\n", final)
}

func TestSatu(t *testing.T) {
	// input
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	sums := 4

	x := 0
	m := len(nums)
	temp := make([]int, sums)
	pair := make([][]int, m)

	for range faktorialNonRecursive(m) {
		// fmt.Printf("nums: %v\n", nums)
		if x >= m-1 {
			x = 0
		}

		count := 0
		for i := 0; i < sums; i++ {
			temp[i] = nums[i]
			count += nums[i]
		}

		fmt.Printf("temp: %v\n", temp)

		if count == target {
			pair = append(pair, temp)
		}

		nums[x], nums[x+1] = nums[x+1], nums[x]
		x++
	}

	// fmt.Printf("pair: %v\n", pair)
}

func TestPermutasi(t *testing.T) {
	kata := "catur"
	p := NewPermutasi(kata)
	result := p.Proses()
	fmt.Printf("result: %v\n", result)
}

func TestFaktorial(t *testing.T) {
	n := 10
	// faktorial menggunakan metode recursive
	x := faktorialRecursive(n, 1)
	fmt.Printf("Recursive - Faktorial dari %v (%v!) adalah %v\n\n", n, n, x)

	y := faktorialNonRecursive(n)
	fmt.Printf("NON Recursive - Faktorial dari %v (%v!) adalah %v\n\n", n, n, y)
}

func faktorialRecursive(n int, total int) int {
	if n >= 1 {
		total *= n
		return faktorialRecursive(n-1, total)
	}
	return total
}

func faktorialNonRecursive(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
