package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := math.MinInt64
	sort.Ints(nums)

	for i, num := range nums {
		sum := num + twoSumClosest(nums, i, target-num)
		if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
			res = sum
		}
	}

	return res
}

func twoSumClosest(nums []int, left, target int) int {
	left, right := left+1, len(nums)-1
	res := math.MinInt64 - target

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return target
		}

		if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
			res = sum
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return res
}

func main() {
	fmt.Print(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
