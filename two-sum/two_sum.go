package main

func main() {}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		i := target - num
		if _, ok := m[i]; ok {
			return []int{m[target-num], index}
		} else {
			m[num] = index
		}
	}

	return nil
}
