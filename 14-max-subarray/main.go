package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	s := maxSubArray(nums)
	fmt.Println(s)

	nums = []int{-2, 1}
	s = maxSubArray(nums)
	fmt.Println(s)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	var max *int

	for {
		for left <= right {
			s := sum(nums[left : right+1])
			if max == nil || *max < s {
				max = &s
			}
			right--
		}
		if left == len(nums)-1 {
			break
		}
		right = len(nums) - 1
		left++
	}

	return *max
}

func sum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
