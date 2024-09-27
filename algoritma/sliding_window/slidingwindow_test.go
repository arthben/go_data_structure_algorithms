package slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	// s := "abcabcbb"
	// s := "bbbbb"
	s := "pwwkew"
	res := slideString(s)
	fmt.Printf("res: %v\n", res)
}

func slideString(s string) int {
	for left, v := range s {
		charValue := string(v)
		fmt.Printf("left: %v - %v\n", left, charValue)
	}
	return 0
}

func XslideString(s string) int {
	count := make(map[rune]int)
	inc := 0
	for _, v := range s {
		if _, ok := count[v]; !ok {
			count[v]++
			inc++
		} else {
			inc--
		}
	}

	fmt.Printf("count: %v\n", count)
	fmt.Printf("count: %v\n", inc)
	return 0
}

func TestSlide(t *testing.T) {
	// k := 3
	// nums := []int{2, 1, 5, 1, 3, 2}
	// k := 93
	// nums := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	// k := 7
	// nums := []int{493, 593, 1446, -6013, 2163, 8450, 3008, -1328, 6195, 4102, -6242, -9971, -8327, 4480, -4972, -8305, -1644, 2320, 331, 2465, 2955, -8386, -7580, 1759, -9576, -8088, -9446, -2916, -3600, 923, 1412, -5390, 4492, 9458, -9336, -4754, -3455, 9597, -324, -5628, 4242, 4076, 8159, -2423, -3449, 6744, 9029, -9231, 2283, 6118, -200, 3610, -7566, -6976, 2519, 9532, 2221, -5167, -2370, 1861, -1558, -9815, -1186, 2021, 7050, -4504, 4926, 8362, 7805, -8274, -351, 5826, 7623, -5520, -2395, -8466, -8461, 3261, -3197}

	// res := sliding(nums, k)
	// res := slide(nums, k)
	// fmt.Printf("res: %v\n", res)

	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}
	res := slideWindow(nums, k)
	fmt.Printf("res: %v\n", res)
}

/*
Maximum Average Sub Array 1 (LeetCode #643)
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/
func slideWindow(nums []int, k int) float64 {
	// contoh nums = [2, 1, 5, 1, 3, 2], k = 3
	left := 0
	current_sum := 0
	maxAverage := -9999999

	// membuat window sepanjang k
	// isi window berisi w := 2 + 1 + 5
	for left < k {
		current_sum += nums[left]
		left++
	}

	// geser window per 1 elemen dan lakukan kalkulasi.
	// geser elemen ke kanan, left akan ditambah dan right bertambah juga
	// idx   : 0  1  2  3  4  5
	// nums  : 2, 1, 5, 1, 3, 2
	// window: 2, 1, 5     == > geser ke kanan
	//     >>:    1, 5, 1  == > saat geser, hasil kalkukasi kurangi dengan angka 2 dan tambah dengan angka 1

	// karena window sudah terbentuk sampai dengan left / k,
	// maka untuk geser dimulai dari posisi terakhir window
	for right := left; right < len(nums)-1; right++ {
		current_sum += nums[right]
		current_sum -= nums[right-k]

		maxAverage = max(maxAverage, current_sum)
	}
	return float64(maxAverage) / float64(k)
}

func sliding(nums []int, k int) float64 {
	maxNumber := float64(math.MinInt)

	for i := 0; i < len(nums)-k+1; i++ {
		count := float64(0)
		for x := i; x < k+i; x++ {
			count += float64(nums[x])
		}

		maxNumber = max(count, maxNumber)
	}

	return maxNumber / float64(k)
}
