package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := mergeTwoSortedArrays(nums1, nums2)
	fmt.Println(merge)
	result := findMedianOfArray(merge)
	return float64(result)
}

func findMedianOfArray(nums []int) float64 {
	var (
		low  int
		high int = len(nums) - 1
	)
	if len(nums)%2 != 0 {
		return float64(nums[int((low+high)/2)])
	} else {
		mid := (float64(low) + float64(high)) / 2
		return (float64(nums[int(mid)]) + float64(nums[int(mid)+1])) / 2
	}
}

func mergeTwoSortedArrays(nums1, nums2 []int) []int {
	var (
		result = make([]int, 0, len(nums1)+len(nums2))
		i, j   int
	)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}
	return result
}
