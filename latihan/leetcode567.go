package latihan

import (
	"fmt"
	"testing"
)

/*
Permutation in string (leetcode #567)

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.

Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false

*/

func TestInc(t *testing.T) {
	// s1 := "ab"
	// s2 := "eidbaooo"
	// s2 := "eidboaoo"

	// s1 := "hello"
	// s2 := "ooolleoooleh"

	s1 := "a"
	s2 := "ab"

	// s1 := "adc"
	// s2 := "dcda"

	fmt.Printf("s1: %v - s2: %v\n", s1, s2)
	res := MycheckInclusion(s1, s2)
	fmt.Printf("res: %v\n", res)
}

func MycheckInclusion(s1 string, s2 string) bool {
	pattern := make(map[byte]int, len(s1))
	for i := range s1 {
		key := s1[i]
		pattern[key]++
	}

	k := len(s1)
	j := 0
	i := 0
	for i < len(s2)+1 {
		if i >= k {
			window := s2[j:i]
			fmt.Printf("window: %v, length: %v\n", window, i-j)
			j++

			temp := make(map[byte]int, len(s1))
			for x := range window {
				key := window[x]
				if _, found := pattern[key]; found {
					temp[key]++
				}
			}

			fmt.Printf("temp: %v\n", temp)
			count := 0
			for y := range pattern {
				if pattern[y] == temp[y] {
					count++
				}
			}

			if count == len(pattern) {
				return true
			}
		}
		i++
	}

	return false
	// for key := range pattern {
	// 	if have[key] != pattern[key] {
	// 		return false
	// 	}
	// }
	// return true
}

func FasterCheckInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	c1, c2 := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		c1[s1[i]-'a']++
		c2[s2[i]-'a']++
	}
	if c1 == c2 {
		return true
	}
	for i := n; i < m; i++ {
		c2[s2[i]-'a']++
		c2[s2[i-n]-'a']--
		if c1 == c2 {
			return true
		}
	}

	return false
}
