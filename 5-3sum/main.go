package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	trip := threeSum(nums)
	fmt.Println(trip)

	nums = []int{0, 0, 0}
	trip = threeSum(nums)
	fmt.Println(trip)

	nums = []int{0, 0, 0, 0}
	trip = threeSum(nums)
	fmt.Println(trip)

	nums = []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	trip = threeSum(nums)
	fmt.Println(trip)

	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	trip = threeSum(nums)
	fmt.Println(trip)
}

func threeSum(nums []int) [][]int {
	l := len(nums) - 1
	triplets := [][]int{}
	if l+1 < 3 {
		return triplets
	}

	sort.Ints(nums)
	hash := map[[3]int]bool{}

	for i := 0; i <= l-2; i++ {
		if nums[i] > 0 {
			return triplets
		}

		j := i + 1
		k := l
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			} else if sum > 0 {
				k--
				continue
			} else if nums[i]+nums[j]+nums[k] == 0 {
				sl := [3]int{nums[i], nums[j], nums[k]}
				if _, ok := hash[sl]; ok {
					j++
					continue
				}
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				hash[sl] = true
			}
			j++
		}
	}

	return triplets
}
