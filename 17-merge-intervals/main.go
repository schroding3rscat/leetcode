package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("%+v\n\n", merge(intervals))

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Printf("%+v\n\n", merge(intervals))

	intervals = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Printf("%+v\n\n", merge(intervals))

	intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("%+v\n\n", merge(intervals))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

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
