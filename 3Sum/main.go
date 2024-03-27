package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for left := 0; left < len(nums)-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		res = append(res, twoSum(nums, left+1, -nums[left])...)
	}

	return res
}

func twoSum(nums []int, leftIndex int, target int) [][]int {
	var res [][]int
	rightIndex := len(nums) - 1

	for leftIndex < rightIndex {
		sum := nums[leftIndex] + nums[rightIndex]

		if sum == target {
			res = append(res, []int{-target, nums[leftIndex], nums[rightIndex]})

			leftIndex++
			rightIndex--

			for leftIndex < rightIndex && nums[rightIndex] == nums[rightIndex+1] {
				rightIndex--
			}

			for leftIndex < rightIndex && nums[leftIndex] == nums[leftIndex-1] {
				leftIndex++
			}

			continue
		}

		if sum < target {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return res
}

func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
