package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 2}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	fmt.Println(nums)
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return i < j })

	r := make([][]int, 0)
	for i := 0; i < len(nums); i++ {

	}
	return r
}
