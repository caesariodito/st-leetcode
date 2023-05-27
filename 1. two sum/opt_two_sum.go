package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := numToIndex[target-nums[i]]; ok {
			return []int{j, i}
		}
		numToIndex[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3)) // [0 1]
}
