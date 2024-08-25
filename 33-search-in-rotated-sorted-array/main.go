package main

import "fmt"

func main() {
	fmt.Println("result = ", search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
}

func search(nums []int, target int) int {
	var (
		low  int
		high int = len(nums) - 1
	)

	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}
	return -1
}
