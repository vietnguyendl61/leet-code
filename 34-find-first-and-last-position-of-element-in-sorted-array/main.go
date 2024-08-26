package main

import "fmt"

func main() {
	fmt.Println(upper([]int{2, 2}, 2))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := upper(nums, target-1)
	right := upper(nums, target)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{left, right - 1}
}

func upper(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) / 2

		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
