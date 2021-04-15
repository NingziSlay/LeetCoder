package main

import "fmt"

// 最大子序和
// https://leetcode-cn.com/problems/maximum-subarray/

func main() {
	x := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray1(x))
	fmt.Println(x)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var result = sum(nums)
	for left, right := 0, len(nums)-1; left <= right; {
		if result-nums[left] >= result-nums[right] {
			left += 1
		} else {
			right -= 1
		}
		result = max(result, sum(nums[left:right+1]))
		fmt.Printf("result: %d, left: %d, right: %d\n", result, left, right)
	}

	return result
}

func sum(nums []int) int {
	var i int
	for _, x := range nums {
		i += x
	}
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray1(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
