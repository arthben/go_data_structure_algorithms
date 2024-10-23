package overlap

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

/*
Non Overlapping Intervals (Leetcode #435)

Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number
of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.

Example 1:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

Example 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
*/
func TestErase(t *testing.T) {
	// intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	// intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	fmt.Printf("Input: %v\n", intervals)
	fmt.Printf("eraseOverlapIntervals(): %v\n", eraseOverlapIntervals(intervals))
	// fmt.Println("Output: 1 {1, 3} is overlap")
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	top := intervals[0]
	for _, interval := range intervals[1:] {
		if slices.Equal(interval, top) {
			count++
			continue
		}

		if top[1] <= interval[0] {
			top = interval
		} else {
			if top[1] > interval[1] {
				top = interval
			}
			count++
		}
	}

	return count
}

/*
Insert Interval (Leetcode #57)

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi]
represent the start and the end of the ith interval and intervals is sorted in ascending order by starti.
You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals
still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.
Note : that you don't need to modify intervals in-place. You can make a new array and return it.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/
func TestInsert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Printf("Input Intervals: %v , input interval: %v\n", intervals, newInterval)
	fmt.Printf("insert(): %v\n", insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var temp [][]int
	temp = append(temp, intervals[0])
	for _, interval := range intervals[1:] {
		prevEnd := temp[len(temp)-1][1]
		currStart := interval[0]
		currEnd := interval[1]

		if currStart > prevEnd {
			temp = append(temp, interval)
		} else if currEnd > prevEnd {
			temp[len(temp)-1][1] = currEnd
		}
	}

	return temp
}

/*
Merge Intervals (Leetcode #56)

Given an array of intervals where intervals[i] = [starti, endi],
merge all overlapping intervals, and return an array of the non-overlapping
intervals that cover all the intervals in the input.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/
func TestMerge(t *testing.T) {
	// intervals := [][]int{{1, 4}, {0, 2}, {3, 5}}
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("Input: %v\n", intervals)
	fmt.Printf("merge(): %v\n", merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var temp [][]int
	temp = append(temp, intervals[0])
	for _, interval := range intervals[1:] {
		prevEnd := temp[len(temp)-1][1]
		currStart := interval[0]
		currEnd := interval[1]

		if currStart > prevEnd {
			temp = append(temp, interval)
		} else if currEnd > prevEnd {
			temp[len(temp)-1][1] = currEnd
		}
	}
	return temp
}
