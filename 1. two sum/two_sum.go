// https://leetcode.com/problems/two-sum/description/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 3)) // [0 1]
}
