package main

func calculateAmountOfWater(leftHeight, rightHeight, leftBorder, rightBorder int) int {
	return min(leftHeight, rightHeight) * (rightBorder - leftBorder)
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0

	for left != right {
		leftHeight := height[left]
		rightHeight := height[right]

		res = max(res, calculateAmountOfWater(leftHeight, rightHeight, left, right))
		if leftHeight < rightHeight {
			left++
			continue
		}

		right--
	}

	return res
}

func main() {

}
