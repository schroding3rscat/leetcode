package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	fmt.Printf("%+v\n\n", insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	i := 1
	j := 0
	for i <= len(intervals)-1 {
		if intervals[i][0] >= intervals[j][0] && intervals[i][0] <= intervals[j][1] {
			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		} else if intervals[j][0] >= intervals[i][0] && intervals[j][0] <= intervals[i][1] {
			intervals[j][0] = intervals[i][0]
			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		} else {
			j++
			intervals[j] = intervals[i]
		}
		i++
	}

	return intervals[:j+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
