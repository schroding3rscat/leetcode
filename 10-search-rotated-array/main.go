package main

import "fmt"

func main() {
	sl := []int{1, 3}
	idx := search(sl, 3)
	fmt.Println(idx)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l := len(nums)
	left := 0
	right := l - 1

	for left < right {
		middle := (right + left) / 2
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if left == 0 { // slice already sorted, pivot == 0
		right = l - 1
	} else if nums[l-1] < target { // we'll search in the first part
		right = left
		left = 0
	} else {
		right = l - 1
	}

	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if left == right && nums[left] != target {
			return -1
		}
	}

	return -1
}
