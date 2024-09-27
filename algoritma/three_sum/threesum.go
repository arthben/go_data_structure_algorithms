package threesum

import (
	"fmt"
	"sort"
	"strings"
)

type ThreeSum interface {
	SolutionOne() [][]int
	SolutionTwo() [][]int
}

type threeData struct {
	data   []int
	target int
}

func NewThreeSum(data []int, target int) ThreeSum {
	return &threeData{data: data, target: target}
}

func (this *threeData) SolutionTwo() [][]int {
	nums := this.data
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		// if nums[i] > 0 {
		// 	break
		// }

		// avoid duplicate input in result
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < 0 {
				left++

			} else {
				right--
			}
		}
	}

	return result
}

func (this *threeData) SolutionOne() [][]int {
	// pendekatanya menggunakn loop 3x
	// time complexity : O(n^3)

	final := make([][]int, 0)
	temp := make(map[string]int, 0)

	nums := this.data
	sort.Ints(nums)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == this.target {
					new := []string{
						fmt.Sprintf("%v", nums[i]),
						fmt.Sprintf("%v", nums[j]),
						fmt.Sprintf("%v", nums[k])}
					keyMap := strings.Join(new, "")
					if _, ok := temp[keyMap]; !ok {
						final = append(final, []int{nums[i], nums[j], nums[k]})
					}
					temp[keyMap] = k
				}
			}
		}
	}

	return final
}
