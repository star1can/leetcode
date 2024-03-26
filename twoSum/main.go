package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if pairIndex, ok := m[target-num]; ok {
			return []int{i, pairIndex}
		}

		m[num] = i
	}

	return []int{}
}

func main() {
	res := twoSum([]int{3, 2, 4}, 6)
	fmt.Print(res)
}
