package main

import (
	"fmt"
	"math"
)

type set map[rune]struct{}

func lengthOfLongestSubstring(s string) int {
	charSet := make(set)

	left := 0
	right := 0
	res := 0

	for _, ch := range s {
		for _, ok := charSet[ch]; ok; _, ok = charSet[ch] {
			delete(charSet, rune(s[left]))
			left++
		}

		charSet[ch] = struct{}{}
		res = int(math.Max(float64(res), float64(right-left+1)))
		right++
	}

	return res
}

func main() {
	fmt.Print(lengthOfLongestSubstring("abcabcbb"))
}
