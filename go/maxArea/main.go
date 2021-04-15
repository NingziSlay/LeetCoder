package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

func maxArea(height []int) int {
	maxA := 0

	left, right := 0, len(height)-1
	for left != right {
		x, y := height[left], height[right]

		maxA = max(maxA, min(x, y)*(right-left))

		fmt.Printf("left: %d, right: %d, maxA: %d\n", left, right, maxA)
		if x > y {
			right -= 1
		} else {
			left += 1
		}
	}
	return maxA
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
